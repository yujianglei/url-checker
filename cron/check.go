package cron

import (
	"crypto/tls"
	"log"
	"strconv"
	"time"

	"url-checker/common"
	"url-checker/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/toolkits/container/list"
)

var CheckResultQueue *list.SafeLinkedList
var WorkerChan chan int

func Init() {
	WorkerChan = make(chan int, 2)
	CheckResultQueue = list.NewSafeLinkedList()
}

type CheckResult struct {
	ID           int64  `json:"id"`
	InstanceName string `json:"instance_name"`
	Item         string `json:"item"`
	RespCode     string `json:"resp_code"`
	RespTime     int64  `json:"resp_time"`
	PushTime     int64  `json:"push_time"`
	Status       int64  `json:"status"`
}

func StartCheck() {
	duca, err := beego.AppConfig.Int64("check.frequency")
	if err != nil {
		log.Fatal(err.Error())
		duca = 10
	}

	t1 := time.NewTicker(time.Duration(duca) * time.Second)
	var GetItems models.UrlItem
	for {
		items, err := GetItems.FindAllItem()
		common.WriteLogErr(err)

		for _, item := range items {
			WorkerChan <- 1
			go Check(item)
		}
		<-t1.C
	}
}

func Check(item *models.UrlItem) {
	defer func() {
		<-WorkerChan
	}()

	checkdata := doCheck(item)
	// fmt.Println(checkdata.Address, checkdata.RespCode, checkdata.RespTime)
	CheckResultQueue.PushFront(checkdata)

}

func doCheck(item *models.UrlItem) (checkResult *CheckResult) {
	checkResult = &CheckResult{
		InstanceName: item.InstanceName,
		Item:         item.Item,
		RespCode:     "0",
	}

	reqStartTime := time.Now()
	req := httplib.Get(item.Item)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetTimeout(3*time.Second, 10*time.Second)
	req.Header("Content-Type", "application/x-www-form-urlencoded; param=value")

	resp, err := req.Response()
	checkResult.PushTime = time.Now().Unix()
	if err != nil {
		log.Println("[ERROR]:", item.Item, err.Error())
		checkResult.Status = common.Failure
		return
	}
	defer resp.Body.Close()

	respCode := strconv.Itoa(resp.StatusCode)
	checkResult.RespCode = respCode

	respDucatime := int64(time.Now().Sub(reqStartTime).Nanoseconds() / 1000000)
	checkResult.RespTime = respDucatime

	if respDucatime > (item.Timeout * 1000) {
		checkResult.Status = common.Failure
		return
	}
	checkResult.Status = common.Success
	return
}

package cron

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/UrlMonitorTool/common"
	"github.com/UrlMonitorTool/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type CheckResult struct {
	ID           int64  `json:"id"`
	InstanceName string `json:"instance_name"`
	Address      string `json:"address"`
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
			log.Fatal(item)
			go checkresultStatus(item)
		}
		<-t1.C
	}
}

func checkresultStatus(item *models.UrlItem) (checkResult *CheckResult) {
	log.Fatal(item)
	checkResult = &CheckResult{
		InstanceName: item.InstanceName,
		Address:      item.Address,
		RespCode:     "0",
	}
	log.Fatal(checkResult)
	reqStartTime := time.Now()
	req := httplib.Get(item.Address)
	// req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetTimeout(3*time.Second, 10*time.Second)
	req.Header("Content-Type", "application/x-www-form-urlencoded; param=value")

	resp, err := req.Response()
	checkResult.PushTime = time.Now().Unix()
	if err != nil {
		log.Fatal("[ERROR]:", item.Address, err.Error())
		checkResult.Status = common.Failure
		return
	}
	defer resp.Body.Close()

	respCode := strconv.Itoa(resp.StatusCode)
	checkResult.RespCode = respCode

	respDucatime := int64(time.Now().Sub(reqStartTime).Nanoseconds() / 1000000)
	checkResult.RespTime = respDucatime

	if respDucatime > item.Timeout {
		checkResult.Status = common.Failure
		return
	}
	checkResult.Status = common.Success
	fmt.Println(checkResult)
	return
}

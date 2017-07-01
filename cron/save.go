package cron

import (
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/influxdata/influxdb/client/v2"
)

func Startsave() {
	for {
		resultSlice := make([]*CheckResult, 0)
		itemResults := CheckResultQueue.PopBack(500)
		if len(itemResults) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		for _, v := range itemResults {
			one := v.(*CheckResult)
			resultSlice = append(resultSlice, one)
		}

		doSave(resultSlice)

	}
}

func doSave(item []*CheckResult) {

	exitWhere := beego.AppConfig.String("exit.where")
	// var test CheckResult
	if exitWhere == "influxdb" {
		for _, v := range item {
			// fmt.Println(v)
			go saveInfluxdb(v)
		}
	}

}

func saveInfluxdb(this *CheckResult) error {
	host := beego.AppConfig.String("indb.host")
	user := beego.AppConfig.String("indb.user")
	pass := beego.AppConfig.String("indb.pass")
	// dbname := beego.AppConfig.String("indb.db")

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     host,
		Username: user,
		Password: pass,
	})
	if err != nil {
		log.Print("conn influxdb failed.")
		log.Fatal(err)
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "test",
		Precision: "s",
	})
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	tags := map[string]string{}

	fileds := map[string]interface{}{
		"item":         this.Item,
		"respcode":     this.RespCode,
		"resptime":     this.RespTime,
		"status":       this.Status,
		"pushtime":     this.PushTime,
		"instancename": this.InstanceName,
	}

	pt, err := client.NewPoint("url", tags, fileds, time.Now())
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	bp.AddPoint(pt)

	if err := c.Write(bp); err != nil {
		log.Println(err)
		return err
	} else {
		c.Close()
		return nil
	}

}

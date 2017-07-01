package main

import (
	"log"
	"runtime"

	"github.com/UrlMonitorTool/cron"
	"github.com/UrlMonitorTool/models"
	_ "github.com/UrlMonitorTool/routers"
	"github.com/astaxie/beego"
)

func prepare() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func init() {
	cron.Init()
	models.InitSQL()

}

func main() {
	go cron.Startsave()
	go cron.StartCheck()
	beego.Run()
}

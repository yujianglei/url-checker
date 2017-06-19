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
	models.InitSQL()
	go cron.StartCheck()
}

func main() {
	beego.Run()
}

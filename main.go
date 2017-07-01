package main

import (
	"log"
	"runtime"

	"url-checker/cron"
	"url-checker/models"
	_ "url-checker/routers"

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

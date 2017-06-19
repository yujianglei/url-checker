package routers

import (
	"github.com/UrlMonitorTool/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/api/item/list", &controllers.ItemController{}, "*:ListItem")
	beego.Router("/api/item/list/:id([0-9]+)", &controllers.ItemController{}, "get:ListItemById")
	beego.Router("/api/item/add", &controllers.ItemController{}, "post:AddItem")
}

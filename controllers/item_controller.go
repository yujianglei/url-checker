package controllers

import (
	"encoding/json"
	"log"

	"github.com/UrlMonitorTool/common"
	"github.com/UrlMonitorTool/models"
	"github.com/astaxie/beego"
)

var urlitem models.UrlItem
var result common.JsonResult

type ItemController struct {
	beego.Controller
}

func (this *ItemController) ListItem() {

	items, err := urlitem.FindAllItem()
	if err != nil {
		log.Fatal(err)
		result.Code = 1
		result.Message = err.Error()
	} else {
		result.Code = 0
		result.Data = items
	}
	this.Data["json"] = result
	this.ServeJSON()

}

func (this *ItemController) ListItemById() {

	id, err := this.GetInt64(":id")
	if err != nil {
		log.Fatal(err)
	}

	item, err := urlitem.FindOneItem(id)
	if err != nil {
		log.Fatal(err.Error())
		result.Code = 1
		result.Message = err.Error()
	} else {
		result.Code = 0
		result.Data = item
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ItemController) AddItem() {
	var urlitem models.UrlItem
	json.Unmarshal(this.Ctx.Input.RequestBody, &urlitem)

	if err := urlitem.AddItem(&urlitem); err != nil {
		log.Fatal(err.Error())
		result.Code = 1
		result.Message = err.Error()
	} else {
		result.Code = 0
	}
	this.Data["json"] = result
	this.ServeJSON()

}

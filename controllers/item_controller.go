package controllers

import (
	"encoding/json"
	"log"

	"url-checker/common"
	"url-checker/models"

	"github.com/astaxie/beego"
)

var urlitem models.UrlItem
var result common.JsonResult

type ItemController struct {
	beego.Controller
}

func (this *ItemController) ListItem() {
	var null []interface{}

	items, err := urlitem.FindAllItem()
	if err != nil {
		log.Println(err.Error())
		result.Code = 1
		result.Message = err.Error()
		result.Data = null
	} else {
		result.Code = 0
		result.Message = ""
		result.Data = items
	}
	this.Data["json"] = result
	this.ServeJSON()

}

func (this *ItemController) ListItemById() {
	var null interface{}
	id, err := this.GetInt64(":id")
	if err != nil {
		log.Fatalln(err.Error())
	}

	item, err := urlitem.FindOneItem(id)
	if err != nil {
		log.Println(err.Error())
		result.Code = 1
		result.Message = err.Error()
		result.Data = null
	} else {
		result.Code = 0
		result.Message = ""
		result.Data = item
	}
	this.Data["json"] = result
	this.ServeJSON()
}

func (this *ItemController) AddItem() {
	var urlitem models.UrlItem
	json.Unmarshal(this.Ctx.Input.RequestBody, &urlitem)

	if err := urlitem.AddItem(&urlitem); err != nil {
		log.Println(err.Error())
		result.Code = 1
		result.Message = err.Error()
	} else {
		result.Code = 0
	}
	this.Data["json"] = result
	this.ServeJSON()

}

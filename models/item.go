package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

type UrlItem struct {
	Id           int64
	InstanceName string
	Item         string
	UrlType      string
	Timeout      int64
	Keyword      string
	Maintainer   string
}

func (u *UrlItem) TableName() string {
	return "instance_url_items"
}

func (this *UrlItem) FindAllItem() ([]*UrlItem, error) {
	var items []*UrlItem
	o := orm.NewOrm()
	table := this.TableName()
	_, err := o.QueryTable(table).OrderBy("-id").All(&items)
	return items, err
}

func (this *UrlItem) FindOneItem(id int64) (UrlItem, error) {
	urlitem := UrlItem{Id: id}
	o := orm.NewOrm()
	if err := o.Read(&urlitem); err != nil {
		return UrlItem{}, err
	}
	return urlitem, nil
}

func (this *UrlItem) AddItem(item *UrlItem) error {
	o := orm.NewOrm()
	if _, err := o.Insert(item); err != nil {
		return errors.New("add failed.")
	}
	return nil
}

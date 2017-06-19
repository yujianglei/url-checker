package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

type UrlItem struct {
	Id           int64
	InstanceName string
	Address      string
	Type         string
	Protocol     string
	Timeout      int64
	Keyword      string
	Frequency    int64
	Option       string
}

func (this *UrlItem) FindAllItem() ([]*UrlItem, error) {
	var items []*UrlItem
	o := orm.NewOrm()
	_, err := o.QueryTable("url_item").OrderBy("-id").All(&items)
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

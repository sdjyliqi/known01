package model

import (
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"known01/utils"
	"time"
)

type Reference struct {
	Id           int              `json:"id" xorm:"not null pk INT(11)"`
	Name         string           `json:"name" xorm:"not null VARCHAR(128)"`
	CategoryId   utils.EngineType `json:"category_id" xorm:"not null comment('分类 1银行 2快递 3中奖') INT(4)"`
	AliasNames   string           `json:"alias_names" xorm:"comment('别名') VARCHAR(1024)"`
	Phone        string           `json:"phone" xorm:"VARCHAR(1024)"`
	ManualPhone  string           `json:"manual_phone" xorm:"not null VARCHAR(32)"`
	Website      string           `json:"website" xorm:"not null comment('官网地址') VARCHAR(255)"`
	MessageId    string           `json:"message_id" xorm:"VARCHAR(1024)"`
	Domain       string           `json:"domain" xorm:"not null comment('多个域名有英文，分割') VARCHAR(4096)"`
	LastModified time.Time        `json:"last_modified" xorm:"DATETIME"`
}

func (t Reference) TableName() string {
	return "reference"
}

func (t Reference) GetItems(engine *xorm.Engine) ([]*Reference, error) {
	var items []*Reference
	err := engine.Find(&items)
	if err != nil {
		glog.Errorf("Get items form table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}
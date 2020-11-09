package models

import (
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"known01/utils"
	"time"
)

type News struct {
	Id      int       `json:"id" xorm:"not null pk INT(11)"`
	Title   string    `json:"title" xorm:"not null VARCHAR(128)"`
	Status  int       `json:"status" xorm:"comment('0 初始值   1 发布  2 下架') TINYINT(4)"`
	Url     string    `json:"url" xorm:"not null VARCHAR(256)"`
	IsReal  int       `json:"is_real" xorm:"not null TINYINT(4)"`
	Publish time.Time `json:"publish" xorm:"not null DATETIME"`
	Author  string    `json:"author" xorm:"not null VARCHAR(32)"`
}

func (t News) TableName() string {
	return "news"
}

func (t News) GetItems(engine *xorm.Engine, pageID int) ([]*News, error) {
	var items []*News
	err := engine.Limit(utils.PageEntry, pageID*utils.PageEntry).Find(&items)
	if err != nil {
		glog.Errorf("Get items form table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

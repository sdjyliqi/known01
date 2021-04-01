package model

import (
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"github.com/sdjyliqi/known01/utils"
	"time"
)

type Templates struct {
	Id           int              `json:"id" xorm:"not null pk autoincr INT(11)"`
	CategoryId   utils.EngineType `json:"category_id" xorm:"not null comment('分类 1银行 2快递 3中奖') INT(4)"`
	SimHash      int64            `json:"simhash" xorm:"not null BIGINT(20)"`
	Detail       string           `json:"detail" xorm:"not null VARCHAR(1024)"`
	Enable       int              `json:"enable" xorm:"comment('是否启动 0，未启动，1 启用') TINYINT(3)"`
	LastModified time.Time        `json:"last_modified" xorm:"not null comment('模板最后更新时间') DATETIME"`
}

func (t Templates) TableName() string {
	return "templates"
}

func (t Templates) GetItems(engine *xorm.Engine) ([]*Templates, error) {
	var items []*Templates
	err := engine.Find(&items)
	if err != nil {
		glog.Fatalf("Get items form table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

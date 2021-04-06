package model

import (
	"github.com/go-xorm/xorm"
	"github.com/prometheus/common/log"
)

type Assist struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name     string `json:"name" xorm:"not null VARCHAR(64)"`
	Enable   int    `json:"enable" xorm:"not null TINYINT(4)"`
	Category string `json:"category" xorm:"not null VARCHAR(32)"`
}

func (t Assist) TableName() string {
	return "assist"
}

func (t Assist) GetItems(engine *xorm.Engine) ([]*Assist, error) {
	var items []*Assist
	err := engine.Find(&items)
	if err != nil {
		log.Errorf("Get items form table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

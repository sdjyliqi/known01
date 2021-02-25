package model

import (
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"time"
)

type Score struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Score        int       `json:"score" xorm:"not null INT(8)"`
	Dimension    string    `json:"dimension" xorm:"not null VARCHAR(64)"`
	LastModified time.Time `json:"last_modified" xorm:"not null DATETIME"`
}

func (t Score) TableName() string {
	return "score"
}

func (t Score) GetItems(engine *xorm.Engine) ([]*Score, error) {
	var items []*Score
	err := engine.Find(&items)
	if err != nil {
		glog.Errorf("Get items form table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

func (t Score) GetItemByIdx(idx string, engine *xorm.Engine) (*Score, error) {
	var item Score
	ok, err := engine.Where("dimension=?", idx).Get(&item)
	if err != nil {
		glog.Errorf("Get items form table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	if !ok {
		glog.Errorf("Do not find the item by dimension %+v from %+v.", idx, t.TableName())
		return nil, errors.New("not-existed")
	}
	return &item, nil
}

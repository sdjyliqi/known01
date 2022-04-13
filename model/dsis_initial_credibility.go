package model

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
	"known01/utils"
	"time"
)

var InitialCredibilityModel DsisInitialCredibility

type DsisInitialCredibility struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Dimension    string    `json:"dimension" xorm:"not null pk unique VARCHAR(64)"`
	Score        int       `json:"score" xorm:"not null INT(8)"`
	LastModified time.Time `json:"last_modified" xorm:"not null DATETIME"`
}

func (t DsisInitialCredibility) TableName() string {
	return "dsis_initial_credibility"
}

func (t DsisInitialCredibility) GetItems() ([]*DsisInitialCredibility, error) {
	var items []*DsisInitialCredibility
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		fmt.Errorf("Get items form table %s failed,err:%+v", t.TableName(), err)
		glog.Fatal("Get items form table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

func (t DsisInitialCredibility) GetItemByIdx(idx string) (*DsisInitialCredibility, error) {
	var item DsisInitialCredibility
	ok, err := utils.GetMysqlClient().Where("dimension=?", idx).Get(&item)
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

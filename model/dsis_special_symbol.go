package model

import (
	"fmt"
	"github.com/golang/glog"
	"known01/utils"
	"time"
)

var SpecialSymbolModel DsisSpecialSymbol

type DsisSpecialSymbol struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Character    string    `json:"special_character" xorm:"unique VARCHAR(16)"`
	WriteUser    string    `json:"submitter" xorm:"VARCHAR(255)"`
	LastModified time.Time `json:"last_modified" xorm:"DATE"`
}

func (t DsisSpecialSymbol) TableName() string {
	return "dsis_special_symbol"
}

func (t DsisSpecialSymbol) GetItems() ([]*DsisSpecialSymbol, error) {
	var items []*DsisSpecialSymbol
	err := utils.GetMysqlClient().Find(&items)
	if err != nil {
		fmt.Errorf("Get items form table %s failed,err:%+v", t.TableName(), err)
		glog.Errorf("Get items form table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

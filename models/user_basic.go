package models

import (
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
)

type UserBasic struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	UserName    string `json:"user_name" xorm:"VARCHAR(64)"`
	Passport    string `json:"passport" xorm:"VARCHAR(64)"`
	Chn         string `json:"chn" xorm:"VARCHAR(128)"`
	MobilePhone string `json:"mphone" xorm:"VARCHAR(64)"`
}

func (t UserBasic) TableName() string {
	return "user_basic"
}

func (t UserBasic) GetUserBasic(client *xorm.Engine, userName string) (*UserBasic, error) {
	item := &UserBasic{}
	_, err := client.Where("user_name = ?", userName).Get(item)
	if err != nil {
		glog.Errorf("[mysql]Get the item for user_name %+v from table %s failed,err:%+v", userName, t.TableName(), err)
		return nil, err
	}
	return item, nil
}

func (t UserBasic) ChkPassportValid(client *xorm.Engine, userName, passport string) (bool, error) {
	item := &UserBasic{}
	_, err := client.Where("user_name = ?", userName).Get(item)
	if err != nil {
		glog.Errorf("[mysql]Get the item for user_name %+v from table %s failed,err:%+v", userName, t.TableName(), err)
		return false, err
	}
	return item.Passport == passport, nil
}

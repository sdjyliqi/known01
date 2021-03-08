package model

import (
	"encoding/json"
	"errors"
	"github.com/golang/glog"
	"known01/utils"
	"time"
)

type User struct {
	Id         int       `json:"id" xorm:"not null pk INT(11)"`
	Name       string    `Name:"Name" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
	Password   string    `json:"password" xorm:"not null comment('登录密码') VARCHAR(64)"`
	Manager    string    `json:"manager" xorm:"comment('负责人') VARCHAR(255)"`
	Phone      string    `json:"phone" xorm:"default '' comment('负责人电话') VARCHAR(32)"`
	Enable     int       `json:"enable" xorm:"comment('是否禁用') TINYINT(4)"`
	Department string    `json:"department" xorm:"default '' comment('部门名称') VARCHAR(128)"`
	LastLogin  time.Time `json:"last_login" xorm:"comment('最后一次登录日期') DATETIME"`
}

func (t User) TableName() string {
	return "user"
}

func (t User) ChkPassword(name, password string) (bool, error) {
	var item User
	ok, err := utils.GetMysqlClient().Where("name = ?", name).Get(&item)
	if err != nil {
		glog.Errorf("Get item from table %s failed,err:%+v", t.TableName(), err)
		return false, err
	}
	if ok {
		return password == item.Password, nil
	}
	return false, errors.New("not-find")
}

//GetItems ...按页获取数据库中的数据，page从0开始
//返items类型为[]User ，原来的[]*User报错
func (t User) GetItems(page, entry int) (string, error) {
	var items []User
	err := utils.GetMysqlClient().Limit(entry, (page-1)*entry).Find(&items)
	if err != nil {
		glog.Errorf("Get items from table %s failed,err:%+v", t.TableName(), err)
		return "nil", err
	}
	result, err := json.Marshal(&items)
	if err != nil {
		return "nil", err
	}
	return string(result), errors.New("not-find")
}

//ModifyEnable ...修改用户状态，账号是否可以使用
func (t User) ModifyEnable(name string) (bool, error) {
	var item User
	ok, err := utils.GetMysqlClient().Where("name = ?", name).Get(&item)
	if err != nil {
		glog.Errorf("Get items from table %s failed, err: %+v", t.TableName(), err)
		return false, err
	}
	if ok || item.Enable == 1 {
		sql := "update user set enable = ? where name = ?"
		_, err := utils.GetMysqlClient().Exec(sql, 1, name)
		if err != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err)
			return false, err
		}
		return true, nil
	}
	if ok || item.Enable == 0 {
		sql := "update user set enable = ? where name = ?"
		_, err := utils.GetMysqlClient().Exec(sql, 0, name)
		if err != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err)
			return false, err
		}
		return true, nil
	}
	return false, errors.New("not-find")
}

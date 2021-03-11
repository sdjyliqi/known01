package model

import (
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
type AddUser struct {
	Name       string `Name:"Name" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
	Manager    string `json:"manager" xorm:"comment('负责人') VARCHAR(255)"`
	Phone      string `json:"phone" xorm:"default '' comment('负责人电话') VARCHAR(32)"`
	Department string `json:"department" xorm:"default '' comment('部门名称') VARCHAR(128)"`
}
type UserInf struct {
	Name       string    `Name:"Name" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
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
func (t User) GetItems(page, entry int) ([]User, error) {
	var items []User
	err := utils.GetMysqlClient().Limit(entry, (page-1)*entry).Find(&items)
	if err != nil {
		glog.Errorf("Get items from table %s failed,err:%+v", t.TableName(), err)
		return items, err
	}
	return items, nil
}

//ShowInf   ...查看用户详细信息
func (t User) ShowInf(name string) (UserInf, error) {
	var inf UserInf
	sql := "Select name, manager, phone, enable, department, last_login from user where name = ?"
	ok, err := utils.GetMysqlClient().SQL(sql, name).Get(&inf)
	if err != nil {
		glog.Errorf("Get item from table %s failed,err:%+v", t.TableName(), err)
		return inf, err
	}
	if ok {
		return inf, nil
	}
	return inf, errors.New("not-find")
}

//ModifyEnable ...修改用户状态，账号是否可以使用
func (t User) ModifyEnable(name string) (bool, error) {
	var item User
	ok, err := utils.GetMysqlClient().Where("name = ?", name).Get(&item)
	if err != nil {
		glog.Errorf("Get items from table %s failed, err: %+v", t.TableName(), err)
		return false, err
	}
	if ok && item.Enable == 1 {
		sql := "update user set enable = ? where name = ?"
		_, err := utils.GetMysqlClient().Exec(sql, 0, name)
		if err != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err)
			return false, err
		}
		return true, nil
	}
	if ok && item.Enable == 0 {
		sql := "update user set enable = ? where name = ?"
		_, err := utils.GetMysqlClient().Exec(sql, 1, name)
		if err != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err)
			return false, err
		}
		return true, nil
	}
	return false, errors.New("not-find")
}

//AddData  ... 增加用户
func (t User) AddData(AddUser AddUser) (bool, error) {
	enable := 1
	sql := "Insert into user(name, password, manager, phone, enable, department, last_login) " +
		"values (?, ?, ?, ?, ?, ?, ?)"
	_, err := utils.GetMysqlClient().Exec(sql, AddUser.Name, "Ceb2732@", AddUser.Manager, AddUser.Phone, enable, AddUser.Department, time.Now().Local())
	if err != nil {
		glog.Errorf("%s table insert data is failed, err: %+v", t.TableName(), err)
		return false, err
	}
	return true, nil
}

//ResetPas ... 重置密码
func (t User) ResetPas(name string) (bool, error) {
	var item User
	ok, err := utils.GetMysqlClient().Where("name = ?", name).Get(&item)
	if ok {
		sql := "update user set password = ? where name = ?"
		_, err := utils.GetMysqlClient().Exec(sql, "Ceb2732@", name)
		if err != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err)
			return false, err
		}
		return true, nil
	}
	return false, err
}

//ChangePas   ... 用户自己修改密码
func (t User) ChangePas(name, newpas string) (bool, error) {
	var item User
	ok, err := utils.GetMysqlClient().Where("name = ?", name).Get(&item)
	if ok {
		sql := "update user set password = ? where name = ?"
		_, err := utils.GetMysqlClient().Exec(sql, newpas, name)
		if err != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err)
			return false, err
		}
		return true, nil
	}
	return false, err
}

//
func (t User) UpdateInf(name, phone, department string) (bool, error) {
	var item User
	ok, _ := utils.GetMysqlClient().Where("name = ?", name).Get(&item)
	if ok {
		sql := "update user set phone = ?, department = ? where name = ?"
		_, err := utils.GetMysqlClient().Exec(sql, phone, department, name)
		if err != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err)
			return false, err
		}
		return true, nil
	}
	return false, errors.New("not found")
}

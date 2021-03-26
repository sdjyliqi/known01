package model

import (
	"errors"
	"github.com/golang/glog"
	"known01/utils"
	"time"
)

//User  ...数据库表结构
type User struct {
	Id           int       `json:"id" xorm:"not null pk INT(11)"`
	Keyid        string    `json:"keyid" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
	Password     string    `json:"password" xorm:"not null comment('登录密码') VARCHAR(64)"`
	Manager      string    `json:"manager" xorm:"default '' comment('负责人') VARCHAR(255)"`
	Roles        string    `json:"roles" xorm:"not null comment('用户权限') VARCHAR(32)"`
	Mobilephone  string    `json:"mobilephone" xorm:"default '' comment('负责人手机号') VARCHAR(32)"`
	Telephone    string    `json:"telephone" xorm:"default '' comment('负责人座机号') VARCHAR(32)"`
	Email        string    `json:"email" xorm:"default '' comment('负责人邮箱') VARCHAR(64)"`
	Enable       int       `json:"enable" xorm:"comment('是否禁用') TINYINT(4)"`
	Organization string    `json:"organization" xorm:"default '' comment('机构名称') VARCHAR(64)"`
	Department   string    `json:"department" xorm:"default '' comment('部门名称') VARCHAR(64)"`
	Office       string    `json:"office" xorm:"default '' comment('处室名称') VARCHAR(64)"`
	LastLogin    time.Time `json:"last_login" xorm:"comment('最后一次登录日期') DATETIME"`
}

//ListUser   ...列表展示用户详情页
type ListUser struct {
	Keyid        string `Json:"keyid" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
	Manager      string `json:"manager" xorm:"comment('负责人') VARCHAR(255)"`
	Mobilephone  string `json:"mobilephone" xorm:"default '' comment('负责人手机号') VARCHAR(32)"`
	Email        string `json:"email" xorm:"default '' comment('负责人邮箱') VARCHAR(64)"`
	Organization string `json:"organization" xorm:"default '' comment('机构名称') VARCHAR(64)"`
	Department   string `json:"department" xorm:"default '' comment('部门名称') VARCHAR(64)"`
}

//UserInf  ... 查看用户详细信息
type UserInf struct {
	Keyid        string    `json:"keyid" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
	Manager      string    `json:"manager" xorm:"default '' comment('负责人') VARCHAR(255)"`
	Roles        string    `json:"roles" xorm:"not null comment('用户权限') VARCHAR(32)"`
	Mobilephone  string    `json:"mobilephone" xorm:"default '' comment('负责人手机号') VARCHAR(32)"`
	Telephone    string    `json:"telephone" xorm:"default '' comment('负责人座机号') VARCHAR(32)"`
	Email        string    `json:"email" xorm:"default '' comment('负责人邮箱') VARCHAR(64)"`
	Enable       int       `json:"enable" xorm:"comment('是否禁用') TINYINT(4)"`
	Organization string    `json:"organization" xorm:"default '' comment('机构名称') VARCHAR(64)"`
	Department   string    `json:"department" xorm:"default '' comment('部门名称') VARCHAR(64)"`
	Office       string    `json:"office" xorm:"default '' comment('处室名称') VARCHAR(64)"`
	LastLogin    time.Time `json:"last_login" xorm:"comment('最后一次登录日期') DATETIME"`
}

func (t User) TableName() string {
	return "user"
}

//
func (t User) GetItemById(keyid string) (User, error) {
	var item User
	ok, err := utils.GetMysqlClient().Where("keyid = ?", keyid).Get(&item)
	if err != nil {
		glog.Errorf("Get item from table %s failed,err:%+v", t.TableName(), err)
		return item, err
	}
	if ok {
		return item, nil
	}
	return item, errors.New("not-find")
}

//GetItems ...按页获取数据库中的数据，page从0开始
//返items类型为[]User ，原来的[]*User报错
func (t User) GetItems(page, entry int) ([]ListUser, error) {
	var items []ListUser
	cols := []string{"keyid", "manager", "mobilephone", "email", "organization", "department"}
	err := utils.GetMysqlClient().Table("user").Cols(cols...).Limit(entry, (page-1)*entry).Find(&items)
	if err != nil {
		glog.Errorf("Get items from table %s failed,err:%+v", t.TableName(), err)
		return items, err
	}
	return items, nil
}

//ShowInf   ...查看用户详细信息
func (t User) ShowInf(keyid string) (UserInf, error) {
	var inf UserInf
	cols := []string{"keyid", "manager", "roles", "mobilephone", "telephone", "email", "enable", "organization",
		"department", "office", "last_login"}
	ok, err := utils.GetMysqlClient().Cols(cols...).Table("user").Where("keyid", keyid).Get(&inf)
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
func (t User) ModifyEnable(keyid string) (bool, error) {
	var item User
	ok, err := utils.GetMysqlClient().Where("keyid = ?", keyid).Get(&item)
	if err != nil {
		glog.Errorf("Get items from table %s failed, err: %+v", t.TableName(), err)
		return false, err
	}
	//如果原来的状态为1则改为0.如果为0改为1
	if ok && item.Enable == 1 {
		item.Enable = 0
		_, err := utils.GetMysqlClient().Cols("enable").Where("keyid", item.Keyid).Update(&item)
		if err != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err)
			return false, err
		}
		return true, nil
	}
	if ok && item.Enable == 0 {
		item.Enable = 1
		_, err := utils.GetMysqlClient().Cols("enable").Where("keyid", item.Keyid).Update(&item)
		if err != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err)
			return false, err
		}
		return true, nil
	}
	return false, errors.New("not-find")
}

//AddData  ... 增加用户
func (t User) InsertItem(AddUser User) (bool, error) {
	AddUser.Enable = 1
	AddUser.Password = "Ceb2732@"
	_, err := utils.GetMysqlClient().Insert(&AddUser)
	if err != nil {
		glog.Errorf("%s table insert data is failed, err: %+v", t.TableName(), err)
		return false, err
	}
	return true, nil
}

//ResetPas ... 重置密码
func (t User) ResetPas(keyid string) (bool, error) {
	var item User
	ok, err1 := utils.GetMysqlClient().Cols("password").Where("keyid = ?", keyid).Get(&item)
	if ok {
		item.Password = "Ceb2732@" // 将初始密码赋给查询到的Item中，然后进行更新
		_, err2 := utils.GetMysqlClient().Cols("password").Where("keyid", item.Keyid).Update(&item)
		if err2 != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err2)
			return false, err2
		}
		return true, nil
	}
	return false, err1
}

//ChangePas   ... 用户自己修改密码
func (t User) ChangePas(keyid, newpas string) (bool, error) {
	var item User
	ok, err1 := utils.GetMysqlClient().Where("keyid = ?", keyid).Get(&item)
	if ok {
		item.Password = newpas // 将新密码赋给查询到的Item中，然后进行更新
		_, err2 := utils.GetMysqlClient().Cols("password").Where("keyid", item.Keyid).Update(&item)
		if err2 != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err2)
			return false, err2
		}
		return true, nil
	}
	return false, err1
}

//UpdateInf   ...用户更新个人信息
func (t User) UpdateItemById(UserUpdate User) (bool, error) {
	var item User
	ok, _ := utils.GetMysqlClient().Where("keyid = ?", UserUpdate.Keyid).Get(&item)
	if ok {
		cols := []string{"mobilephone", "telephone", "email", "office"}
		_, err := utils.GetMysqlClient().Cols(cols...).Where("keyid", UserUpdate.Keyid).Update(&UserUpdate)
		if err != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err)
			return false, err
		}
		return true, nil
	}
	return false, errors.New("not found")
}

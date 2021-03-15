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
	Name         string    `json:"name" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
	Password     string    `json:"password" xorm:"not null comment('登录密码') VARCHAR(64)"`
	Manager      string    `json:"manager" xorm:"default '' comment('负责人') VARCHAR(255)"`
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
	Name         string `Json:"Name" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
	Manager      string `json:"manager" xorm:"comment('负责人') VARCHAR(255)"`
	Mobilephone  string `json:"mobilephone" xorm:"default '' comment('负责人手机号') VARCHAR(32)"`
	Email        string `json:"email" xorm:"default '' comment('负责人邮箱') VARCHAR(64)"`
	Organization string `json:"organization" xorm:"default '' comment('机构名称') VARCHAR(64)"`
	Department   string `json:"department" xorm:"default '' comment('部门名称') VARCHAR(64)"`
}

//AddUser   ... 管理员添加用户前台传入数据
type AddUser struct {
	Name         string `Json:"Name" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
	Manager      string `json:"manager" xorm:"comment('负责人') VARCHAR(255)"`
	Mobilephone  string `json:"mobilephone" xorm:"default '' comment('负责人手机号') VARCHAR(32)"`
	Telephone    string `json:"telephone" xorm:"default '' comment('负责人座机号') VARCHAR(32)"`
	Email        string `json:"email" xorm:"default '' comment('负责人邮箱') VARCHAR(64)"`
	Organization string `json:"organization" xorm:"default '' comment('机构名称') VARCHAR(64)"`
	Department   string `json:"department" xorm:"default '' comment('部门名称') VARCHAR(64)"`
	Office       string `json:"office" xorm:"default '' comment('处室名称') VARCHAR(64)"`
}

//UserInf  ... 查看用户详细信息
type UserInf struct {
	Name         string    `json:"name" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
	Manager      string    `json:"manager" xorm:"default '' comment('负责人') VARCHAR(255)"`
	Mobilephone  string    `json:"mobilephone" xorm:"default '' comment('负责人手机号') VARCHAR(32)"`
	Telephone    string    `json:"telephone" xorm:"default '' comment('负责人座机号') VARCHAR(32)"`
	Email        string    `json:"email" xorm:"default '' comment('负责人邮箱') VARCHAR(64)"`
	Enable       int       `json:"enable" xorm:"comment('是否禁用') TINYINT(4)"`
	Organization string    `json:"organization" xorm:"default '' comment('机构名称') VARCHAR(64)"`
	Department   string    `json:"department" xorm:"default '' comment('部门名称') VARCHAR(64)"`
	Office       string    `json:"office" xorm:"default '' comment('处室名称') VARCHAR(64)"`
	LastLogin    time.Time `json:"last_login" xorm:"comment('最后一次登录日期') DATETIME"`
}

//UserUpdate  ... 用户更新个人信息
type UserUpdate struct {
	Name        string `json:"name" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
	Mobilephone string `json:"mobilephone" xorm:"default '' comment('负责人手机号') VARCHAR(32)"`
	Telephone   string `json:"telephone" xorm:"default '' comment('负责人座机号') VARCHAR(32)"`
	Email       string `json:"email" xorm:"default '' comment('负责人邮箱') VARCHAR(64)"`
	Office      string `json:"office" xorm:"default '' comment('处室名称') VARCHAR(64)"`
}

func (t User) TableName() string {
	return "user"
}

//ChkPassword   ...核对用户密码
func (t User) ChkPassword(name, password string) (bool, error) {
	var item User
	ok, err := utils.GetMysqlClient().Where("name = ?", name).Get(&item)
	if err != nil {
		glog.Errorf("Get item from table %s failed,err:%+v", t.TableName(), err)
		return false, err
	}
	if ok {
		//更新last_login
		sql := "update user set last_login = ? where name = ?"
		_, err := utils.GetMysqlClient().Exec(sql, time.Now().Local(), name)
		if err != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err)
			return false, err
		}
		return password == item.Password, nil

	}
	return false, errors.New("not-find")
}

//GetItems ...按页获取数据库中的数据，page从0开始
//返items类型为[]User ，原来的[]*User报错
func (t User) GetItems(page, entry int) ([]ListUser, error) {
	var items []ListUser
	sql := "Select * from user Limit ? OFFSET ?"
	err := utils.GetMysqlClient().SQL(sql, entry, (page-1)*entry).Find(&items)
	if err != nil {
		glog.Errorf("Get items from table %s failed,err:%+v", t.TableName(), err)
		return items, err
	}
	return items, nil
}

//ShowInf   ...查看用户详细信息
func (t User) ShowInf(name string) (UserInf, error) {
	var inf UserInf
	sql := "Select name, manager, mobilephone, telephone, email, enable, organization, " +
		"department, office, last_login from user where name = ?"
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
	sql := "Insert into user(name, password, manager, mobilephone, telephone, email, enable, organization, " +
		"department, office, last_login) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := utils.GetMysqlClient().Exec(sql, AddUser.Name, "Ceb2732@", AddUser.Manager, AddUser.Mobilephone,
		AddUser.Telephone, AddUser.Email, enable, AddUser.Organization, AddUser.Department, AddUser.Office, time.Now().Local())
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

//UpdateInf   ...用户修改个人信息
func (t User) UpdateInf(UserUpdate UserUpdate) (bool, error) {
	var item User
	ok, _ := utils.GetMysqlClient().Where("name = ?", UserUpdate.Name).Get(&item)
	if ok {
		sql := "update user set mobilephone = ?, telephone = ?, email = ?, office = ? where name = ?"
		_, err := utils.GetMysqlClient().Exec(sql, UserUpdate.Mobilephone, UserUpdate.Telephone,
			UserUpdate.Email, UserUpdate.Office, UserUpdate.Name)
		if err != nil {
			glog.Errorf("%s table update data is failed, err: %+v", t.TableName(), err)
			return false, err
		}
		return true, nil
	}
	return false, errors.New("not found")
}

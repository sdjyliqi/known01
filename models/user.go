package models

import (
	"time"
)

type User struct {
	Id         int       `json:"id" xorm:"not null pk INT(11)"`
	Key        string    `json:"key" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
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

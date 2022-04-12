package models

import (
	"time"
)

type DsisUserCenter struct {
	Id            int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Keyid         string    `json:"keyid" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
	Manager       string    `json:"manager" xorm:"default '' comment('负责人') VARCHAR(255)"`
	Roles         string    `json:"roles" xorm:"not null comment('用户权限') VARCHAR(32)"`
	Mobilephone   string    `json:"mobilephone" xorm:"default '' comment('负责人手机号') VARCHAR(32)"`
	Telephone     string    `json:"telephone" xorm:"default '' comment('负责人座机号') VARCHAR(32)"`
	Email         string    `json:"email" xorm:"default '' comment('负责人邮箱') VARCHAR(64)"`
	Enable        int       `json:"enable" xorm:"comment('是否禁用') TINYINT(4)"`
	Organization  string    `json:"organization" xorm:"default '' comment('机构名称') VARCHAR(64)"`
	Department    string    `json:"department" xorm:"default '' comment('部门名称') VARCHAR(64)"`
	SubDepartment string    `json:"sub_department" xorm:"default '' comment('处室名称') VARCHAR(64)"`
	LastLogin     time.Time `json:"last_login" xorm:"comment('最后一次登录日期') DATETIME"`
}

func (t DsisUserCenter) TableName() string {
	return "dsis_user_center"
}

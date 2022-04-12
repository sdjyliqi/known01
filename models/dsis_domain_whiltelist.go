package models

import (
	"time"
)

type DsisDomainWhiltelist struct {
	Id           int       `json:"id" xorm:"not null pk autoincr comment('域名白名单自增id') INT(11)"`
	Company      string    `json:"company" xorm:"not null comment('发件人邮箱简称') VARCHAR(64)"`
	Domain       string    `json:"domain" xorm:"not null comment('发件人邮箱后缀') VARCHAR(32)"`
	WriteUser    string    `json:"write_user" xorm:"VARCHAR(16)"`
	Desc         string    `json:"desc" xorm:"VARCHAR(128)"`
	LastModified time.Time `json:"last_modified" xorm:"DATETIME"`
}

func (t DsisDomainWhiltelist) TableName() string {
	return "dsis_domain_whiltelist"
}

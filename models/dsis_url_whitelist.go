package models

import (
	"time"
)

type DsisUrlWhitelist struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	ObjectName   string    `json:"object_name" xorm:"comment('URL所属对象名称') VARCHAR(128)"`
	Url          string    `json:"url" xorm:"not null comment('URL地址') VARCHAR(255)"`
	Status       int       `json:"status" xorm:"not null default 0 comment('是否启用，0为禁用，1为启用') TINYINT(4)"`
	Submitter    string    `json:"submitter" xorm:"default 'admin' comment('最后更新者') VARCHAR(36)"`
	LastModified time.Time `json:"last_modified" xorm:"comment('最后更新日期') DATETIME"`
}

func (t DsisUrlWhitelist) TableName() string {
	return "dsis_url_whitelist"
}

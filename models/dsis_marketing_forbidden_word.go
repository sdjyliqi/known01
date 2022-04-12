package models

import (
	"time"
)

type DsisMarketingForbiddenWord struct {
	Id            int       `json:"id" xorm:"not null pk autoincr unique INT(11)"`
	ForbiddenWord string    `json:"forbidden_word" xorm:"not null comment('营销文案禁用词') VARCHAR(128)"`
	Status        int       `json:"status" xorm:"not null default 0 comment('是否启用，0为禁用，1为启用') TINYINT(4)"`
	Submitter     string    `json:"submitter" xorm:"default 'admin' comment('最后提交者') VARCHAR(32)"`
	LastModified  time.Time `json:"last_modified" xorm:"comment('最后更新时间') DATETIME"`
}

func (t DsisMarketingForbiddenWord) TableName() string {
	return "dsis_marketing_forbidden_word"
}

package models

import (
	"time"
)

type Reference struct {
	Id           int       `json:"id" xorm:"not null pk INT(11)"`
	CategoryId   int       `json:"category_id" xorm:"not null INT(11)"`
	Name         string    `json:"name" xorm:"not null VARCHAR(128)"`
	AliasNames   string    `json:"alias_names" xorm:"comment('别名') VARCHAR(1024)"`
	Phone        string    `json:"phone" xorm:"VARCHAR(1024)"`
	SenderId     string    `json:"sender_id" xorm:"VARCHAR(32)"`
	ManualPhone  string    `json:"manual_phone" xorm:"not null VARCHAR(32)"`
	Website      string    `json:"website" xorm:"not null comment('官网地址') VARCHAR(255)"`
	MessageId    string    `json:"message_id" xorm:"VARCHAR(1024)"`
	Domain       string    `json:"domain" xorm:"not null comment('多个域名有英文，分割') VARCHAR(4096)"`
	LastModified time.Time `json:"last_modified" xorm:"DATETIME"`
}

func (t Reference) TableName() string {
	return "reference"
}

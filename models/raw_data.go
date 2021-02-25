package models

import (
	"time"
)

type RawData struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	CategoryId   int       `json:"category_id" xorm:"not null INT(4)"`
	Detail       string    `json:"detail" xorm:"not null VARCHAR(1024)"`
	Enable       int       `json:"enable" xorm:"not null TINYINT(3)"`
	LastModified time.Time `json:"last_modified" xorm:"not null DATETIME"`
}

func (t RawData) TableName() string {
	return "raw_data"
}

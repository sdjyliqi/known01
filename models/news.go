package models

import (
	"time"
)

type News struct {
	Id      int       `json:"id" xorm:"not null pk INT(11)"`
	Thumb   string    `json:"thumb" xorm:"not null VARCHAR(256)"`
	Title   string    `json:"title" xorm:"not null VARCHAR(128)"`
	IsReal  int       `json:"is_real" xorm:"TINYINT(4)"`
	Status  int       `json:"status" xorm:"comment('0 初始值   1 发布  2 下架') TINYINT(4)"`
	Url     string    `json:"url" xorm:"not null VARCHAR(256)"`
	Publish time.Time `json:"publish" xorm:"not null DATETIME"`
	Author  string    `json:"author" xorm:"not null VARCHAR(32)"`
	Comment string    `json:"comment" xorm:"VARCHAR(256)"`
}

func (t News) TableName() string {
	return "news"
}

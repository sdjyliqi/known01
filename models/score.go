package models

import (
	"time"
)

type Score struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Dimension    string    `json:"dimension" xorm:"not null pk unique VARCHAR(64)"`
	Score        int       `json:"score" xorm:"not null INT(8)"`
	LastModified time.Time `json:"last_modified" xorm:"not null DATETIME"`
}

func (t Score) TableName() string {
	return "score"
}

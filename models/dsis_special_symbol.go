package models

import (
	"time"
)

type DsisSpecialSymbol struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Character    string    `json:"character" xorm:"unique VARCHAR(16)"`
	WriteUser    string    `json:"write_user" xorm:"VARCHAR(255)"`
	LastModified time.Time `json:"last_modified" xorm:"DATE"`
}

func (t DsisSpecialSymbol) TableName() string {
	return "dsis_special_symbol"
}

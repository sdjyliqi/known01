package models

import (
	"time"
)

type DsisCorrectionCharacter struct {
	Id           int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Raw          string    `json:"raw" xorm:"not null unique VARCHAR(1)"`
	Replace      string    `json:"replace" xorm:"VARCHAR(1)"`
	WriteUser    string    `json:"write_user" xorm:"VARCHAR(16)"`
	LastModified time.Time `json:"last_modified" xorm:"DATE"`
}

func (t DsisCorrectionCharacter) TableName() string {
	return "dsis_correction_character"
}

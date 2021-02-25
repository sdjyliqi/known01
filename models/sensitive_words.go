package models

type SensitiveWords struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Word     string `json:"word" xorm:"not null unique CHAR(128)"`
	Category string `json:"category" xorm:"default '' VARCHAR(128)"`
}

func (t SensitiveWords) TableName() string {
	return "sensitive_words"
}

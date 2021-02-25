package models

type Assist struct {
	Id       int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	Name     string `json:"name" xorm:"not null VARCHAR(64)"`
	Enable   int    `json:"enable" xorm:"not null TINYINT(4)"`
	Category string `json:"category" xorm:"not null VARCHAR(32)"`
}

func (t Assist) TableName() string {
	return "assist"
}

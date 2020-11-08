package models

type Category struct {
	Id   int    `json:"id" xorm:"not null pk INT(11)"`
	Name string `json:"name" xorm:"not null VARCHAR(255)"`
}

func (t Category) TableName() string {
	return "category"
}

package model

type Weight struct {
	Id           int     `json:"id" xorm:"not null pk INT(11)"`
	NameWeight   float64 `json:"name_weight" xorm:"DOUBLE"`
	DomainWeight float64 `json:"domain_weight" xorm:"DOUBLE"`
	PhoneWeight  float64 `json:"phone_weight" xorm:"DOUBLE"`
}

func (t Weight) TableName() string {
	return "weight"
}

package models

type TrainingData struct {
	Id          int    `json:"id" xorm:"not null pk autoincr INT(11)"`
	CategoryId  int    `json:"category_id" xorm:"not null INT(4)"`
	Detail      string `json:"detail" xorm:"not null VARCHAR(1024)"`
	Dimension   string `json:"dimension" xorm:"not null VARCHAR(64)"`
	Domain      int    `json:"domain" xorm:"not null TINYINT(4)"`
	MessageId   int    `json:"message_id" xorm:"not null TINYINT(4)"`
	Phone       int    `json:"phone" xorm:"not null TINYINT(4)"`
	TrueOrFalse int    `json:"true_or_false" xorm:"not null TINYINT(3)"`
	Enable      int    `json:"enable" xorm:"not null TINYINT(3)"`
}

func (t TrainingData) TableName() string {
	return "training_data"
}

package model

import (
	"github.com/golang/glog"
	"known01/utils"
	"time"
)

var ResultTool Results

type Results struct {
	Id           int              `json:"id" xorm:"not null pk INT(11)"`
	CategoryId   utils.EngineType `json:"category_id" xorm:"INT(11)"`
	Detail       string           `json:"detail" xorm:"TEXT"`
	Extract      string           `json:"extract" xorm:"comment('从原始数据中提取的有效数据组成的json') VARCHAR(4096)"`
	Compare      string           `json:"compare" xorm:"TEXT"`
	Flag         int              `json:"flag" xorm:"TINYINT(4)"`
	Suggest      string           `json:"suggest" xorm:"comment('返回的结果') VARCHAR(4096)"`
	LastModified time.Time        `json:"last_modified" xorm:"DATETIME"`
}

func (t Results) TableName() string {
	return "results"
}

//Store ... 存储
func (t Results) Store(item *Results) error {
	_, err := utils.GetMysqlClient().Insert(item)
	if err != nil {
		glog.Errorf("Insert item %+v to table %s failed,err:%+v", item, t.TableName(), err)
	}
	return err
}

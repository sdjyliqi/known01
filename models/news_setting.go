package models

import (
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"time"
)

type NewsSetting struct {
	Id         int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	Status     int       `json:"status" xorm:"comment('是否上线该rule，1 上线，2 下线 0 初始值') TINYINT(4)"`
	RuleOrder  int       `json:"rule_order" xorm:"comment('策略顺序，值越小，优先级越高') SMALLINT(6)"`
	TimeStart  int       `json:"time_start" xorm:"INT(11)"`
	TimeStop   int       `json:"time_stop" xorm:"INT(11)"`
	ChnList    string    `json:"chn_list" xorm:" varchar(4096)"`
	ChnSel     int       `json:"chn_sel" xorm:"comment('渠道是否正选或者反选') TINYINT(4)"`
	CitySel    int       `json:"city_sel" xorm:"comment('正选是0，反选是1') TINYINT(4)"`
	CityList   string    `json:"city_list" xorm:"comment('城市列表json marshal后的结果') varchar(4096)"`
	HidSel     int       `json:"hid_sel" xorm:"TINYINT(4)"`
	HidList    string    `json:"hid_list" xorm:"TEXT"`
	VersionSel int       `json:"version_sel" xorm:"TINYINT(4)"`
	VersionMin string    `json:"version_min" xorm:"VARCHAR(64)"`
	Response   string    `json:"response" xorm:"comment('下发的内容(json格式数据）') VARCHAR(2048)"`
	VersionMax string    `json:"version_max" xorm:"VARCHAR(64)"`
	LastUpdate time.Time `json:"last_update" xorm:"DATETIME"`
}

func (t NewsSetting) TableName() string {
	return "news_setting"
}

//FindItems ...获取全部列表
func (t NewsSetting) FindItems(client *xorm.Engine) ([]*NewsSetting, error) {
	var items []*NewsSetting
	err := client.Where("status = 1").OrderBy("rule_order").Find(&items)
	if err != nil {
		glog.Errorf("[NewsSetting] FindItems failed,err :%+v", err)
		return nil, err
	}
	return items, nil
}

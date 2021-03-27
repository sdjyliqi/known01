package model

import (
	"errors"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"known01/utils"
	"time"
)

type Reference struct {
	Id           int              `json:"id" xorm:"not null pk INT(11)"`
	Name         string           `json:"name" xorm:"not null VARCHAR(128)"`
	CategoryId   utils.EngineType `json:"category_id" xorm:"not null comment('分类 1银行 2快递 3中奖') INT(4)"`
	AliasNames   string           `json:"alias_names" xorm:"comment('别名') VARCHAR(1024)"`
	Phone        string           `json:"phone" xorm:"VARCHAR(1024)"`
	ManualPhone  string           `json:"manual_phone" xorm:"not null VARCHAR(32)"`
	Website      string           `json:"website" xorm:"not null comment('官网地址') VARCHAR(255)"`
	MessageId    string           `json:"message_id" xorm:"VARCHAR(1024)"`
	Domain       string           `json:"domain" xorm:"not null comment('多个域名有英文，分割') VARCHAR(4096)"`
	LastModified time.Time        `json:"last_modified" xorm:"DATETIME"`
}

//鉴别短信真假时将读取基准数据
func (t Reference) TableName() string {
	return "reference"
}
func (t Reference) GetItems(engine *xorm.Engine) ([]*Reference, error) {
	var items []*Reference
	err := engine.Find(&items)
	if err != nil {
		glog.Errorf("Get items form table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

//GetPages   ...分页查询数据
func (t Reference) GetPages(page, entry int) ([]*Reference, error) {
	var items []*Reference
	err := utils.GetMysqlClient().Limit(entry, page*entry).Find(&items)
	if err != nil {
		glog.Errorf("Get items form table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return items, nil
}

func (t Reference) GetItemByID(ID int) (*Reference, error) {
	var item Reference
	_, err := utils.GetMysqlClient().ID(ID).Get(&item)
	if err != nil {
		glog.Errorf("Get items from table %s failed,err:%+v", t.TableName(), err)
		return nil, err
	}
	return &item, nil
}

//UpdateItemByID ...根据ID 更新数据
func (t Reference) UpdateItemByID(ID int, item *Reference) error {
	item.LastModified = time.Now()
	cols := []string{"name", "category_id", "alias_names", "phone", "manual_phone", "website", "message_id", "domain", "last_modified"}
	_, err := utils.GetMysqlClient().ID(ID).Cols(cols...).Update(item)
	if err != nil {
		glog.Errorf("Update item %+v from table %s failed,err:%+v", item, t.TableName(), err)
		return err
	}
	return nil
}

//InsertItemByID ... 向数据库中插入一条数据，如果name字段已经存在，，无需重复插入
func (t Reference) InsertItemByID(item Reference) error {
	var r Reference
	item.LastModified = time.Now()
	ok, err := utils.GetMysqlClient().Where("name=?", item.Name).Get(&r)
	if err != nil {
		glog.Errorf("Get item by name %+v from table %s failed,err:%+v", item.Name, t.TableName(), err)
		return err
	}
	if ok {
		glog.Infof("The item %+v already existed in table %+s", item, t.TableName())
		return errors.New("already-existed")
	}
	_, err = utils.GetMysqlClient().Insert(&item)
	if err != nil {
		glog.Errorf("Insert item  %+v to table %s failed,err:%+v", item, t.TableName(), err)
		return err
	}
	return nil
}

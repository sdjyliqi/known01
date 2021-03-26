package utils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/golang/glog"
	"sync"
)

var mysqlOnce sync.Once
var msqlEngine *xorm.Engine

//InitMySQL ...初始化mysql，记得要先解密
func InitMySQL(addr string, showSQL bool) (*xorm.Engine, error) {
	var err error
	addrDecode, _ := Decrypt(addr)
	mysqlOnce.Do(func() {
		msqlEngine, err = xorm.NewEngine("mysql", addrDecode)
		msqlEngine.ShowSQL(showSQL) //将查询语句展示在控制台
		if err != nil {
			glog.Errorf("[init] Initialize mysql client for addr %s failed,err:%+v,please check the config.", addr, err)

		}
	})
	return msqlEngine, err
}

//GetMysqlClient ...获取mysql客户端连接
func GetMysqlClient() *xorm.Engine {
	return msqlEngine
}

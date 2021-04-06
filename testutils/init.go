package testutils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/prometheus/common/log"
	"github.com/sdjyliqi/known01/conf"
	"github.com/sdjyliqi/known01/utils"
)

var DBEngineTest *xorm.Engine

func init() {
	DBEngineTest, _ = utils.InitMySQL(conf.DefaultConfig.DBMysql, true)
	utils.InitSegDic()
}

func Init() {
	log.Info("Init mysql client for test")
}

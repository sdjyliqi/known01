package testutils

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sdjyliqi/known01/conf"
	"github.com/sdjyliqi/known01/utils"
)

var DBEngineTest *xorm.Engine

func init() {
	DBEngineTest, _ = utils.InitMySQL(conf.DefaultConfig.DBMysql, true)
}

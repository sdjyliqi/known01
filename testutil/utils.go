package testutil

//UT 相关的变量初始化
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/sdjyliqi/feirars/conf"
	"github.com/sdjyliqi/feirars/utils"
	"sync"
	"testing"
)

var testOnce sync.Once
var TestCfg conf.FeirarConfig
var TestMysql *xorm.Engine

func init() {
	testing.Init()
	testOnce.Do(func() {
		TestCfg = conf.DefaultConfig

		//初始化一下mysql
		TestMysql, _ = utils.InitMySQL(TestCfg.DBMysql, true)
		TestMysql = utils.GetMysqlClient()
		TestMysql.ShowSQL(true)
		//初始化一下redis
		//utils.InitRedisClients(TestCfg.DBRedis)
		//TestRedis = utils.GetRedisClient()

	})
}

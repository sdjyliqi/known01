package main

import (
	"flag"
	"fmt"
	"github.com/cleey/glogrotate"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"known01/conf"
	"known01/handle"
	"known01/router"
	"known01/utils"
	"math/rand"
	"time"
)

//logCut...完成日志切割,默认值为1小时切割一次，只需设置保留的时间即可。
func logCut() {
	glogrotate.Start(glogrotate.RotateOption{
		Remain: time.Hour * 24 * 7,
	})
}

func init() {
	var ymlPath string
	flag.StringVar(&ymlPath, "c", "", "configuration file")
	flag.Parse()
	if ymlPath == "" {
		glog.Fatal("You must input path of the yml ....")
	}
	//初始化配置，覆盖原来的默认配置参数苏
	conf.InitConfig(ymlPath, &conf.DefaultConfig)
	//检查配置项的合法性，如果任何一项为空，立即fatal掉
	if conf.DefaultConfig.DBMysql == "" || conf.DefaultConfig.Port == 0 || conf.DefaultConfig.WordDic == "" {
		fmt.Println("The content of yml is invalid.")
		glog.Errorln("The content of yml is invalid.")
		glog.Fatal("The content of yml is invalid.")
	}
	rand.Seed(time.Now().UnixNano())
	utils.InitMySQL(conf.DefaultConfig.DBMysql, true) //建立MySQL连接
	utils.InitSegDic()                                //初始化分词词表
	handle.InitBrain()                                //初始化
	logCut()
}

func main() {
	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	// register the `/metrics` route.
	router.InitRouter(r)
	r.Run(fmt.Sprintf("0.0.0.0:%d", conf.DefaultConfig.Port))
}

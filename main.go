package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"known01/conf"
	"known01/router"
	"known01/utils"
	"math/rand"
	"time"
)

func init() {
	var ymlPath string
	flag.StringVar(&ymlPath, "c", "", "configuration file")
	flag.Parse()
	if ymlPath == "" {
		glog.Fatal("You must input path of the yml ....")
	}
	conf.InitConfig(ymlPath, &conf.DefaultConfig)
	//check the content items from yml
	if conf.DefaultConfig.DBMysql == "" || conf.DefaultConfig.Port == 0 || conf.DefaultConfig.WordDic == "" {
		glog.Fatal("The content of yml is invalid.")
	}
	rand.Seed(time.Now().UnixNano())
	utils.InitMySQL(conf.DefaultConfig.DBMysql, true) //建立MySQL连接
	utils.InitSegDic()                                //初始化分词词表
}

func main() {
	defer glog.Flush()
	r := gin.Default()
	// register the `/metrics` route.
	router.InitRouter(r)
	r.Run(fmt.Sprintf("0.0.0.0:%d", conf.DefaultConfig.Port))
}

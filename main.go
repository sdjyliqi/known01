package main

import (
	"fmt"
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"known01/conf"
	"known01/router"
)

//func init() {
//	var confFile string
//	flag.StringVar(&confFile, "c", "", "configuration file")
//	flag.Parse()
//	if confFile == "" {
//		glog.Fatal("You must input the conf....")
//	}
//	conf.Init(confFile, &conf.DefaultConfig)
//	rand.Seed(time.Now().UnixNano())
//	//utils.InitMySQL(conf.DefaultConfig.DBMysql, true) //建立MySQL连接
//	//utils.InitSegDic()                                //初始化分词词表
//}

func main() {
	defer glog.Flush()
	r := gin.Default()
	// register the `/metrics` route.
	router.InitRouter(r)
	ginpprof.Wrapper(r)
	r.Run(fmt.Sprintf("0.0.0.0:%d", conf.DefaultConfig.Port))
}

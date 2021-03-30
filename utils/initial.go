package utils

import (
	"flag"
	"github.com/golang/glog"
	"known01/conf"
	"math/rand"
	"time"
)

func init() {
	var confFile string
	flag.StringVar(&confFile, "c", "", "configuration file")
	flag.Parse()
	if confFile == "" {
		glog.Fatal("You must input path of the yml ....")
	}
	conf.Init(confFile, &conf.DefaultConfig)
	rand.Seed(time.Now().UnixNano())
	InitMySQL(conf.DefaultConfig.DBMysql, true) //建立MySQL连接
	InitSegDic()                                //初始化分词词表
}

package utils

import (
	"known01/conf"
)

func init() {
	InitMySQL(conf.DefaultConfig.DBMysql, true) //建立MySQL连接
	//InitSegDic()                                //初始化分词词表
}

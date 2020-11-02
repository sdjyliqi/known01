package utils

import "known01/conf"

func init() {
	InitMySQL(conf.DefaultConfig.DBMysql, true)
}

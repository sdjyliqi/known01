package conf

// Known01Config ...
type Known01Config struct {
	DBMysql string `yaml:"db_mysql"`
	DBRedis string `yaml:"db_redis"`
	WordDic string `yaml:"word_dic"`
}

// DefaultConfig .
var DefaultConfig = Known01Config{
	DBMysql: "root:Bit0123456789!@tcp(114.55.139.105:3306)/brain?charset=utf8mb4",
	DBRedis: "redis://1:@127.0.0.1:6379",
	WordDic: "D:\\gowork\\src\\known01\\data\\dictionary.txt",
}

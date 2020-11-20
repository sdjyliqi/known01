package conf

// Known01Config ...
type Known01Config struct {
	DBMysql string `yaml:"db_mysql"`
	DBRedis string `yaml:"db_redis"`
	WordDic string `yaml:"word_dic"`
}

// DefaultConfig .
var DefaultConfig = Known01Config{
	DBMysql: "wdzHhOX/SSdWWziV4TDy0AYqXfr0dwPoVWNGPbgg26gLOoV0731EyR/b49lfJSSf6dnK0C9s5Il4QyRmaFsNTc6XOtu1ApToSaYGns+OVasYdbGpKsbRqyYRroZ0sirBC8VEyx8FbcWlXQ==",
	WordDic: "D:\\gowork\\src\\known01\\data\\dictionary.txt",
}

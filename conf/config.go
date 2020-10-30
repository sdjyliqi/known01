package conf

// Known01Config ...
type Known01Config struct {
	DBMysql string `yaml:"db_mysql"`
	DBRedis string `yaml:"db_redis"`
}

// DefaultConfig .
var DefaultConfig = Known01Config{
	DBMysql: "pingback-0001a:Pinback-123987!@tcp(127.0.0.1:3306)/pingback?charset=utf8mb4",
	DBRedis: "redis://1:@127.0.0.1:6379",
}

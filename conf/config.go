package conf

// Known01Config ...
type Known01Config struct {
	DBMysql string `yaml:"db_mysql"`
	DBRedis string `yaml:"db_redis"`
	WordDic string `yaml:"word_dic"`
}

// DefaultConfig .
var DefaultConfig = Known01Config{
	DBMysql: "ba0A1eeS3EWffh4vqPk1ni97oJfTXxWajESBsjSKOUgKu+OpPp4oFSUkcoGoxGfLH/YFCy5g6PgaN9iwiBNQ/2ADQG2q2H1a8N0AfeMWS7axkbTA8eJy7kBrQMCaqNvwfD17xs9KmiAWPvSEBHPBoIROUYc=",
	WordDic: "D:\\learn\\gowork\\known01\\data\\dictionary.txt",
}

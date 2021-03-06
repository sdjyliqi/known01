package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// ConfigFile ..
var ConfigFile string

// Known01Config ...
type Known01Config struct {
	DBMysql string `yaml:"db_mysql"`
	Port    int    `yaml:"port"`
	WordDic string `yaml:"word_dic"`
}

// YAMLLoad 加载文件并解析，包含加密项的自动解密
func YAMLLoad(fn string, v *Known01Config) error {
	dat, err := ioutil.ReadFile(fn)
	if err != nil {
		return fmt.Errorf("read config file %v error. err = %v", fn, err)
	}

	err = yaml.Unmarshal(dat, v)
	if err != nil {
		return fmt.Errorf("parse config file %v error. err = %v", fn, err)
	}
	log.Printf("config initialize success. config = %v\n", v)
	return nil
}

// Init 传入带有默认值的 config, 加载配置到 config 中
func InitConfig(f string, v *Known01Config) {
	log.Printf("Init the yaml:%s", f)
	err := YAMLLoad(f, v)
	if err != nil {
		log.Fatalf("Call YAMLLoad failed,err:%+v", err)
	}
}

//DefaultConfig .
var DefaultConfig = Known01Config{
	DBMysql: "wdzHhOX/SSdWWziV4TDy0AYqXfr0dwPoVWNGPbgg26gLOoV0731EyR/b49lfJSSf6dnK0C9s5Il4QyRmaFsNTc6XOtu1ApToSaYGns+OVasYdbGpKsbRqyYRroZ0sirBC8VEyx8FbcWlXQ==",
	WordDic: "D:\\gowork\\src\\known01\\data\\dictionary.txt",
	Port:    8899,
}

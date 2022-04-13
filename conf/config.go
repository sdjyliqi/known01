package conf

import (
	"fmt"
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

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
	glog.Infof("config initialize success. config = %v\n", v)
	return nil
}

// Init 传入带有默认值的 config, 加载配置到 config 中
func InitConfig(f string, v *Known01Config) {
	err := YAMLLoad(f, v)
	if err != nil {
		fmt.Println("Init yaml failed:", err)
		glog.Fatalf("Call YAMLLoad failed,err:%+v", err)
	}
	fmt.Println("yaml conf:", v)
}

//DefaultConfig .
var DefaultConfig = Known01Config{
	DBMysql: "root:Bit0123456789!@tcp(127.0.0.1:3306)/data_guarder?charset=utf8mb4",
	WordDic: "D:\\gowork\\src\\known01\\data\\dictionary.txt",
	Port:    25001,
}

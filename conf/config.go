package conf

import (
	"fmt"
	"github.com/golang/glog"
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
func Init(f string, v *Known01Config) {
	glog.Infof("Init the yaml:%", f)
	err := YAMLLoad(f, v)
	if err != nil {
		glog.Fatal("Call YAMLLoad failed,err:%+v", err)
	}
}

//DefaultConfig .
var DefaultConfig = Known01Config{
	DBMysql: "1wesK74jdqpFdSxpVSWyXTemUDaumleec+lIAr9+viJTSy0YwWnfg0o6fWiM5Rq6Fo0fZ9xSWQ6GfIVZZ95684tsvF4TcLIl8/oVdMYSI8Vx2TwRuFVAza12+WhKLetyDzvpMNlKymA35PRY9rU=",
	WordDic: "D:\\learn\\gowork\\master\\known01\\data\\dictionary.txt",
	Port:    8899,
}

package brain

import "github.com/gansidui/ahocorasick"

var bankNames = map[string][]string{
	"中国人民银行": []string{"人民银行", "中央银行", "央行"},
	"广大银行":   []string{"中国光大银行", "光大"},
	"工商银行":   []string{"中国工商银行", "工商行", "工行"},
	"招商银行":   []string{"中国招商银行", "招商行", "招行"},
}

type bankBrain struct {
	aliasNames map[string]string //存储别名映射
	allNames   []string
	acMatch    *ahocorasick.Matcher
}

//CreateBankBrain ... 创建银行鉴别引擎
func CreateBankBrain() *bankBrain {

	return &bankBrain{}
}

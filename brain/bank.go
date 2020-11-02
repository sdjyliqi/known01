package brain

import "github.com/gansidui/ahocorasick"

var message = "[招商银行]尊敬的客户，一张闪电贷专属礼券为你呈上！用券条款可享受专属利率优惠，优惠日截止2020年10月31日。" +
	"点击http://a.cmbchina.com/personal/cmhkas13快速申请，详情请咨询95599,400-66666888,15210510285"

var bankNames = map[string][]string{
	"中国人民银行": []string{"人民银行", "中央银行", "央行"},
	"广大银行":   []string{"中国光大银行", "光大"},
	"工商银行":   []string{"中国工商银行", "工商行", "工行"},
	"招商银行":   []string{"中国招商银行", "招商行", "招行"},
}

var bankKeywords = []string{"人民银行", "中央银行", "招商银行", "中国招商银行", "招商行", "招行"}
var customerPhones = []string{"400-66666888", "95599", "12345", "你", "光大"}

type bankBrain struct {
	aliasNames    map[string]string //存储别名映射
	allNames      []string
	customerPhone []string
	acMatch       *ahocorasick.Matcher
	phoneACMatch  *ahocorasick.Matcher
}

//CreateBankBrain ... 创建银行鉴别引擎
func CreateBankBrain() *bankBrain {
	brain := bankBrain{
		aliasNames:   map[string]string{},
		allNames:     []string{},
		acMatch:      nil,
		phoneACMatch: nil,
	}
	ac := ahocorasick.NewMatcher()
	brain.allNames = bankKeywords
	ac.Build(bankKeywords)
	brain.acMatch = ac

	acCustomerPhone := ahocorasick.NewMatcher()
	brain.customerPhone = customerPhones
	ac.Build(customerPhones)
	brain.phoneACMatch = acCustomerPhone

	//通过别名，找到标准名称
	for k, v := range bankNames {
		brain.aliasNames[k] = k
		for _, vv := range v {
			brain.aliasNames[vv] = k
		}
	}
	return &brain
}

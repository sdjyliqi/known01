package brain

import (
	"github.com/gansidui/ahocorasick"
	"github.com/sdjyliqi/known01/model"
	"github.com/sdjyliqi/known01/utils"
)

//propertiesVec... 定义文本识别提取数据维度
type propertiesVec struct {
	govName       string //特征名称，如中国工商银行
	webDomain     string //域名，目前只匹配一级域
	mobilePhone   string //手机号码
	customerPhone string //客服电话，如400-***-******
	senderID      string //发信者号码，如95599，可能和客服电话不一致
	fixedPhone    string //固定电话，如010-88888888
}

//indexWordDid ... 还需要融合mysql中结构名称的相关信息。
var indexWordDic = map[utils.EngineType][]string{
	utils.EngineBank:   []string{"汇款", "转账", "打钱", "存款", "银行", "储蓄", "取款", "ATM", "信贷", "信用卡", "储蓄卡", "利息", "贷款", "利率", "负债"},
	utils.EngineReward: []string{"中奖"},
}

type Center struct {
	messageTemplates      map[string]utils.EngineType   //短信模块内容列表
	messageTemplatesItems []*model.DsisMessageTemplates //短信模块内容列表
	cutWords              []string                      //副助词列表

	acCustomerPhoneMatch *ahocorasick.Matcher                    //提取官方客服电话的ac自动机
	customerPhoneDic     map[string][]*model.DsisEnterpriseBasic //银行类短信模块内容列表
	customerPhones       []string                                //客服电话列表，ac自动机匹配查询使用

	//构建分类的关键词ac自动机
	indexWords      []string             //客服电话列表，ac自动机匹配查询使用
	acIndexWords    *ahocorasick.Matcher //提取官方客服电话的ac自动机
	indexWordsDic   map[string]utils.EngineType
	bank            *bankBrain
	referencesItems []*model.DsisEnterpriseBasic
	acCutWords      *ahocorasick.Matcher //副助词匹配自动机，作用，进行模板匹配前，需要将辅助词剔除。

}

//CreateCenter ...创建控制中心
func CreateCenter() Center {
	c := Center{
		messageTemplates: map[string]utils.EngineType{},
		indexWordsDic:    map[string]utils.EngineType{},
		customerPhoneDic: map[string][]*model.DsisEnterpriseBasic{},
	}
	c.init()
	c.bank = CreateBankBrain(c.referencesItems)
	return c
}

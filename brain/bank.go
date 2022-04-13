package brain

import (
	"github.com/gansidui/ahocorasick"
	"known01/model"
)

type bankBrain struct {
	aliasNames      map[string]string //存储别名和标准名称的映射
	allNames        []string
	acMatch         *ahocorasick.Matcher
	phoneNumDic     map[string][]string                   //存储客服电话到标准名称的映射。
	bankDic         map[string]*model.DsisEnterpriseBasic //名称到详情的映射表
	referencesItems []*model.DsisEnterpriseBasic          //基准数据
	scoreDict       map[string]*model.DsisInitialCredibility
}

//CreateBankBrain ... 创建银行鉴别引擎
func CreateBankBrain(items []*model.DsisEnterpriseBasic) *bankBrain {
	brain := bankBrain{
		aliasNames:  map[string]string{},
		allNames:    []string{},
		phoneNumDic: map[string][]string{},
		bankDic:     map[string]*model.DsisEnterpriseBasic{},
		acMatch:     nil,
		scoreDict:   map[string]*model.DsisInitialCredibility{},
	}
	brain.Init(items)
	return &brain
}

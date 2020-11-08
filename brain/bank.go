package brain

import (
	"github.com/gansidui/ahocorasick"
	"known01/models"
)

var messageTest = "[招商银行]尊敬的客户，一张闪电贷专属礼券为你呈上！用券条款可享受专属利率优惠，优惠日截止2020年10月31日。" +
	"点击http://a.cmbchina.com/personal/cmhkas13快速申请，详情请咨询95599,400-66666888,15210510285"

type bankBrain struct {
	aliasNames map[string]string //存储别名映射
	allNames   []string
	acMatch    *ahocorasick.Matcher

	phoneNumDic     map[string]string            //存储客服电话到标准名称的映射。
	bankDic         map[string]*models.Reference //名称到详情的映射表
	referencesItems []*models.Reference          //基准数据

}

//CreateBankBrain ... 创建银行鉴别引擎
func CreateBankBrain(items []*models.Reference) *bankBrain {
	brain := bankBrain{
		aliasNames:  map[string]string{},
		allNames:    []string{},
		phoneNumDic: map[string]string{},
		bankDic:     map[string]*models.Reference{},
		acMatch:     nil,
	}
	brain.Init(items)
	return &brain
}

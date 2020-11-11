package brain

import (
	"fmt"
	"github.com/gansidui/ahocorasick"
	"github.com/golang/glog"
	"known01/models"
	"known01/utils"
	"strings"
)

var messageTest = "[招商银行]尊敬的客户，一张闪电贷专属礼券为你呈上！用券条款可享受专属利率优惠，优惠日截止2020年10月31日。" +
	"点击http://a.cmbchina.com/personal/cmhkas13快速申请，详情请咨询95599,400-66666888,15210510285"

//getBankNameByPhoneID ...通过客服电话查找银行名称
func (bb *bankBrain) Init(items []*models.Reference) error {
	//初始化 PhoneNumDic，aliasNames
	aliasNamesDic := map[string]string{}
	var bankAllNames []string
	for _, v := range items {
		//初始化银行名称的映射关系和关键字列表
		aliasNamesDic[v.Name] = v.Name
		bankAllNames = append(bankAllNames, v.Name)
		if len(v.AliasNames) > 0 {
			names := strings.Split(v.AliasNames, ",")
			bankAllNames = append(bankAllNames, names...)
			for _, vv := range names {
				aliasNamesDic[vv] = v.Name
			}
		}

		bb.bankDic[v.Name] = v
		//客服电话映射表
		phone := strings.ReplaceAll(v.ManualPhone, "-", "")
		bb.phoneNumDic[phone] = v.Name
		if len(v.Phone) > 0 {
			phoneIDs := strings.Split(v.Phone, ",")
			for _, vv := range phoneIDs {
				bb.phoneNumDic[vv] = v.Name
			}
		}
	}
	bb.aliasNames = aliasNamesDic
	//基于银行名称创建ac自动机
	ac := ahocorasick.NewMatcher()
	bb.allNames = bankAllNames
	ac.Build(bankAllNames)
	bb.acMatch = ac
	return nil
}

//getBankNameByPhoneID ...通过客服电话查找银行名称
func (bb *bankBrain) getBankNameByPhoneID(phone string) (string, bool) {
	v, ok := bb.phoneNumDic[phone]
	if !ok {
		glog.Errorf("Do not find the bank-name by customer phone %s", phone)
		return "", false
	}
	return v, ok
}

//PickupProperties ... 摘取核心内容
func (bb *bankBrain) PickupProperties(msg, phoneID string) (propertiesVec, bool) {
	pickVal := propertiesVec{}
	//优先通过客服电话id获取银行名称，如果找不到，只能通过ac自动机来寻找银行关键字。
	if len(phoneID) > 0 {
		govName, ok := bb.getBankNameByPhoneID(phoneID)
		if ok {
			pickVal.govName = govName
		}
	} else {
		govName, ok := bb.pickupName(msg)
		if ok {
			pickVal.govName = govName
		}
	}

	firstDomain, ok := bb.pickupWebDomain(msg)
	if ok {
		pickVal.webDomain = firstDomain
	}

	mobilePhone, ok := bb.pickupMobilePhone(msg)
	if ok {
		pickVal.mobilePhone = mobilePhone
	}
	return pickVal, true
}

//pickupName ... 寻找银行名称，返回值为标准名称
func (bb *bankBrain) pickupName(msg string) (string, bool) {
	matchIndex := bb.acMatch.Match(msg)
	if len(matchIndex) > 0 {
		idx := bb.allNames[matchIndex[0]]
		v, ok := bb.aliasNames[idx]
		if !ok {
			glog.Exitf("Do not find the key %s in dic.", idx)
			return "", false
		}
		return v, true
	}
	return "", false
}

//pickupWebDomain ...寻找一级域名，返回值中已经剔除.com，.cn等辅助信息
func (bb *bankBrain) pickupWebDomain(msg string) (string, bool) {
	return utils.ExtractWebFirstDomain(msg)
}

//pickupMobilePhone ...寻找手机号码
func (bb *bankBrain) pickupMobilePhone(msg string) (string, bool) {
	return utils.ExtractMobilePhone(msg)
}

func (bb *bankBrain) JudgeMessage(msg, phoneID string) (float64, string) {
	v, ok := bb.PickupProperties(msg, phoneID)
	if !ok {
		return 0, ""
	}
	return bb.MatchScoreV2(v)
}

//createMatchScoreIndex ...创建匹配字符串
func (bb *bankBrain) createMatchScoreIndex(pickup propertiesVec) (string, *models.Reference) {
	domainIdx, msgIDIdx, phoneIDIdx := "D0", "M0", "P0"
	if pickup.govName == "" {
		return "", nil
	}
	item, ok := bb.bankDic[pickup.govName]
	if !ok {
		return "", nil
	}
	//checkout website domain
	if item.Domain == pickup.webDomain {
		domainIdx = "D1"
	} else {
		domainIdx = "D2"
	}
	//checkout message sender id
	if strings.HasSuffix(pickup.senderID, item.MessageId) {
		msgIDIdx = "M1"
	} else {
		msgIDIdx = "M2"
	}

	//checkout customer phone id
	if len(pickup.customerPhone) > 0 {
		_, ok := bb.phoneNumDic[pickup.customerPhone]
		if ok {
			phoneIDIdx = "P1"
		} else {
			phoneIDIdx = "P2"
		}
	}
	return domainIdx + msgIDIdx + phoneIDIdx, item
}

func (bb *bankBrain) MatchScoreV2(pickup propertiesVec) (float64, string) {
	notFindMessage := "尊敬的用户，是真是假APP提示您，你接收的短信类型为【金融】，目前未识别出关键信息，请加强安全意识，切勿泄露个人信息，认准官方。"
	matchMessage := "尊敬的用户，是真是假APP提示您，你接收的短信类型为【金融】，目前判断短信内容可信度为%d%%，请致电官方客服%s或登录官方网站%s进行再次确认，避免上当，谢谢您使用时真是假APP。。"
	matchScore := 0.0

	idx, bankItem := bb.createMatchScoreIndex(pickup)
	if idx == "" {
		return matchScore, notFindMessage
	}
	scoreItem, err := models.Score{}.GetItemByIdx(idx, utils.GetMysqlClient())
	if err != nil {
		return matchScore, notFindMessage
	}
	suggest := fmt.Sprintf(matchMessage, int(scoreItem.Score), bankItem.ManualPhone, bankItem.Website)
	return matchScore, suggest

}

//MatchScore ...计算匹配分数，分支越高，可信度越高。
func (bb *bankBrain) MatchScore(pickup propertiesVec) (float64, string) {
	notFindMessage := "尊敬的用户，是真是假APP提示您，你接收的短信类型为【金融】，目前未识别出关键信息，请加强安全意识，切勿泄露个人信息，认准官方。"
	matchMessage := "尊敬的用户，是真是假APP提示您，你接收的短信类型为【金融】，目前判断短信内容可信度为%d%%，请致电官方客服%s或登录官方网站%s进行再次确认，避免上当，谢谢您使用时真是假APP。。"
	matchScore := 0.0
	if pickup.govName == "" {
		return 0, notFindMessage
	}
	referenceItem, ok := bb.bankDic[pickup.govName]
	if !ok {
		return 0, notFindMessage
	}
	if pickup.senderID != "" && pickup.webDomain != "" && pickup.customerPhone != "" {
		//如果全部正确
		matchScore := 0.0
		_, phoneExisted := bb.phoneNumDic[pickup.customerPhone]
		if pickup.webDomain == referenceItem.Domain && phoneExisted {
			if strings.HasSuffix(pickup.senderID, referenceItem.MessageId) {
				matchScore = 1.0
			} else {
				matchScore = 0.6
			}
		}
		if pickup.webDomain == referenceItem.Domain && !phoneExisted {
			if strings.HasSuffix(pickup.senderID, referenceItem.MessageId) {
				matchScore = 0.9
			} else {
				matchScore = 0.55
			}
		}
		suggest := fmt.Sprintf(matchMessage, int(matchScore*100), referenceItem.ManualPhone, referenceItem.Website)
		return matchScore, suggest
	}

	//如果提交的sender_id为空
	if pickup.webDomain != "" {
		if pickup.webDomain == referenceItem.Domain {
			matchScore = 0.6
		}
		if len(pickup.customerPhone) > 0 {
			_, ok := bb.phoneNumDic[pickup.customerPhone]
			if ok {
				matchScore += 0.2
			}
		}
	}
	suggest := fmt.Sprintf(matchMessage, int(matchScore*100), referenceItem.ManualPhone, referenceItem.Website)
	return matchScore, suggest
}

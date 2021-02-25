package brain

import (
	"fmt"
	"github.com/gansidui/ahocorasick"
	"github.com/golang/glog"
	"known01/model"
	"known01/utils"
	"strings"
)

//getBankNameByPhoneID ...通过客服电话查找银行名称
func (bb *bankBrain) Init(items []*model.Reference) error {
	//初始化 PhoneNumDic，aliasNames
	aliasNamesDic := map[string]string{}
	var bankAllNames []string
	for _, v := range items {
		if v.CategoryId != utils.EngineBank {
			continue
		}
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
	//初始化分数字典
	err := bb.InitScoreItems()
	if err != nil {
		glog.Errorf("call InitScoreItems failed,err:%+v", err)
		return err
	}
	return nil
}

func (bb *bankBrain) InitScoreItems() error {
	scoreDic := map[string]*model.Score{}
	items, err := model.Score{}.GetItems(utils.GetMysqlClient())
	if err != nil {
		return err
	}
	for _, v := range items {
		scoreDic[v.Dimension] = v
	}
	bb.scoreDict = scoreDic
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
func (bb *bankBrain) PickupProperties(msg, phoneID, sender string) (propertiesVec, bool) {
	pickVal := propertiesVec{senderID: sender}
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

func (bb *bankBrain) JudgeMessage(msg, phoneID, sender string) (int, string) {
	v, ok := bb.PickupProperties(msg, phoneID, sender)
	if !ok {
		return 0, ""
	}
	score, suggest := bb.MatchScoreV2(v)
	return score, suggest
}

//createMatchScoreIndex ...创建匹配字符串
func (bb *bankBrain) createMatchScoreIndex(pickup propertiesVec) (string, *model.Reference) {
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

func (bb *bankBrain) MatchScoreV2(pickup propertiesVec) (int, string) {
	findMobilePhoneScore := -10
	notFindMessage := "尊敬的用户，是真是假APP提示您，你接收的短信类型为【金融】，目前未识别出关键信息，请加强安全意识，切勿泄露个人信息，认准官方。"
	matchMessage := "尊敬的用户，是真是假APP提示您，你接收的短信类型为【金融】，目前判断短信内容可信度为%d%%，请致电官方客服%s或登录官方网站%s进行再次确认，避免上当，谢谢您使用时真是假APP。。"
	matchScore := 0
	idx, bankItem := bb.createMatchScoreIndex(pickup)
	if idx == "" {
		return 0, notFindMessage
	}
	scoreItem, ok := bb.scoreDict[idx]
	if !ok {
		return 0, notFindMessage
	}
	//如果出现手机号，减分项目
	matchScore = scoreItem.Score
	if len(pickup.mobilePhone) > 1 {
		matchScore = matchScore + findMobilePhoneScore
	}
	if matchScore < 0 {
		matchScore = 0
	}
	suggest := fmt.Sprintf(matchMessage, matchScore, bankItem.ManualPhone, bankItem.Website)
	return matchScore, suggest
}

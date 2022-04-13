package brain

import (
	"github.com/gansidui/ahocorasick"
	"github.com/golang/glog"
	"known01/model"
	"known01/utils"
	"strings"
)

//getBankNameByPhoneID ...通过客服电话查找银行名称
func (bb *bankBrain) Init(items []*model.DsisEnterpriseBasic) error {
	//初始化 PhoneNumDic，aliasNames
	aliasNamesDic := map[string]string{}
	var bankAllNames []string
	for _, v := range items {
		aliasNamesDic[v.Name] = v.Name
		bankAllNames = append(bankAllNames, v.Name)
		if len(v.AliasNames) > 0 {
			names := strings.Split(v.AliasNames, utils.SplitChar)
			bankAllNames = append(bankAllNames, names...)
			for _, vv := range names {
				aliasNamesDic[vv] = v.Name
			}
		}
		bb.bankDic[v.Name] = v
		//客服电话映射表
		phone := strings.ReplaceAll(v.ManualPhone, "-", "")
		_, ok := bb.phoneNumDic[phone]
		if ok {
			bb.phoneNumDic[phone] = utils.SliceUnique(append(bb.phoneNumDic[phone], v.Name))
		} else {
			bb.phoneNumDic[phone] = []string{v.Name}
		}

		if len(v.Phone) > 0 {
			phoneIDs := strings.Split(v.Phone, ",")
			for _, vv := range phoneIDs {
				_, ok := bb.phoneNumDic[vv]
				if ok {
					bb.phoneNumDic[vv] = utils.SliceUnique(append(bb.phoneNumDic[vv], v.Name))
				} else {
					bb.phoneNumDic[vv] = []string{v.Name}
				}
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
	scoreDic := map[string]*model.DsisInitialCredibility{}
	items, err := model.InitialCredibilityModel.GetItems()
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
func (bb *bankBrain) getBankNameByPhoneID(phone, hit, msg string) (string, bool) {
	//优先处理【】符合中的内容，如果名称为整理的基准数据，直接使用该值。
	if len(hit) > 0 {
		v, ok := bb.aliasNames[hit]
		if ok {
			return v, true
		}
	}
	items, ok := bb.phoneNumDic[phone]
	if !ok {
		glog.Errorf("Do not find the bank-name by customer phone %s", phone)
		return "", false
	}
	if len(items) == 1 {
		return items[0], true
	}
	//如果电话对应多个标准名称，使用标准名称查找到基准数据，然后利用基准数据中的name和别名去待鉴别的短信中去查找
	for _, v := range items {
		item, ok := bb.bankDic[v]
		if !ok {
			continue
		}
		if strings.Contains(msg, item.Name) {
			return item.Name, true
		}
		//如果待鉴别短信中包括昵称信息，直接返回对应基准数据的标准名称
		aliasNames := strings.Split(item.AliasNames, utils.SplitChar)
		for _, v := range aliasNames {
			if strings.Contains(msg, v) {
				return item.Name, true
			}
		}
	}
	return "", false
}

//PickupProperties ... 摘取核心内容,特别需要注意的是，一个客服电话对应多个单位的时候。
func (bb *bankBrain) PickupProperties(msg, phoneID, sender string) (propertiesVec, bool) {
	pickVal := propertiesVec{senderID: sender, fixedPhone: phoneID}
	//优先通过客服电话id获取银行名称，如果找不到，只能通过ac自动机来寻找银行关键字。
	hit := utils.PickupHits(msg)
	if len(phoneID) > 0 {
		govName, ok := bb.getBankNameByPhoneID(phoneID, hit, msg)
		if ok {
			pickVal.govName = govName
		}
	} else {
		govName, ok := bb.pickupName(hit, msg)
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

//pickupName ... 寻找银行名称，返回值为标准名称,如果【**】名称在AC自动机中匹配，优先使用。
func (bb *bankBrain) pickupName(hit, msg string) (string, bool) {
	matchIndex := bb.acMatch.Match(msg)
	if len(matchIndex) == 0 {
		return "", false
	}
	//优先处理【】符合中的内容，如果名称为整理的基准数据，直接使用该值。
	if len(hit) > 0 {
		for _, v := range matchIndex {
			name := bb.allNames[v.Index]
			if hit == name {
				v, ok := bb.aliasNames[name]
				if !ok {
					glog.Errorf("Do not find the key %s in dic.", name)
					return "", false
				}
				return v, true
			}
		}
	}
	if len(matchIndex) > 0 {
		//优先使用【值】的值
		idx := bb.allNames[matchIndex[0].Index]
		v, ok := bb.aliasNames[idx]
		if !ok {
			glog.Errorf("Do not find the key %s in dic.", idx)
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

func (bb *bankBrain) JudgeMessage(msg, phoneID, sender string) (int, *model.DsisEnterpriseBasic) {
	v, ok := bb.PickupProperties(msg, phoneID, sender)
	if !ok {
		return utils.OutsideKnown, nil
	}
	return bb.MatchScoreV2(v, sender)
}

//createMatchScoreIndex ...创建匹配字符串
func (bb *bankBrain) createMatchScoreIndex(pickup propertiesVec) (string, *model.DsisEnterpriseBasic) {
	domainIdx, msgIDIdx, phoneIDIdx := "D0", "M0", "P0"
	if pickup.govName == "" {
		return "", nil
	}
	item, ok := bb.bankDic[pickup.govName]
	if !ok {
		return "", nil
	}
	//checkout website domain
	if pickup.webDomain != "" {
		domains := strings.Split(item.Domain, ",")
		domainDic := map[string]bool{}
		for _, v := range domains {
			if len(v) > 1 {
				domainDic[strings.ToLower(v)] = true
			}
		}
		_, ok = domainDic[strings.ToLower(pickup.webDomain)]
		if ok {
			domainIdx = "D1"
		} else {
			domainIdx = "D2"
		}
	}
	//checkout message sender id
	if pickup.senderID != "" && item.SenderId != "" {
		if strings.HasSuffix(pickup.senderID, item.SenderId) {
			msgIDIdx = "M1"
		} else {
			msgIDIdx = "M2"
		}
	}
	//checkout customer phone id
	if len(pickup.fixedPhone) > 0 {
		_, ok := bb.phoneNumDic[pickup.fixedPhone]
		if ok {
			phoneIDIdx = "P1"
		} else {
			phoneIDIdx = "P2"
		}
	}
	return domainIdx + msgIDIdx + phoneIDIdx, item
}

func (bb *bankBrain) MatchScoreV2(pickup propertiesVec, sender string) (int, *model.DsisEnterpriseBasic) {
	findMobilePhoneScore, matchScore, senderScore := 0, 0, 0
	idx, bankItem := bb.createMatchScoreIndex(pickup)
	if idx == "" {
		return utils.OutsideKnown, bankItem
	}
	//如果各个维度均无法确认，直接返回未识别
	if idx == utils.OutsideIndex {
		return utils.OutsideKnown, bankItem
	}
	scoreItem, ok := bb.scoreDict[idx]
	if !ok {
		return utils.OutsideKnown, bankItem
	}

	if utils.ChkContentIsMobilePhone(sender) {
		senderScore = utils.ScoreSenderMobile
	}
	if pickup.mobilePhone != "" {
		findMobilePhoneScore = utils.ScoreFindMobile
	}
	//基础分值加两个维度的浮动分值
	matchScore = scoreItem.Score + senderScore + findMobilePhoneScore
	if len(pickup.mobilePhone) > 1 {
		matchScore = matchScore + findMobilePhoneScore
	}
	if matchScore < 0 {
		matchScore = 0
	}
	return matchScore, bankItem
}

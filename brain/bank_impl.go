package brain

import (
	"github.com/golang/glog"
	"known01/utils"
)

//getBankNameByPhoneID ...通过客服电话查找银行名称
func (bb *bankBrain) getBankNameByPhoneID(phone string) (string, bool) {
	v, ok := bb.customerPhoneDic[phone]
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
	if len(customerPhones) > 0 {
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
		pickVal.website = firstDomain
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

//MatchScore ...计算匹配分数，分支越高，可信度越高。
func (bb *bankBrain) MatchScore(pickup propertiesVec) float64 {
	return 0.0
}

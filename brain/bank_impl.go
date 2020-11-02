package brain

import (
	"fmt"
	"github.com/golang/glog"
	"known01/utils"
)

func (bb *bankBrain) PickupProperties(msg string) (propertiesVec, bool) {
	pickVal := propertiesVec{}
	govName, ok := bb.pickupName(msg)
	if ok {
		pickVal.govName = govName
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
	matchIndex := bb.acMatch.Match(message)
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

//pickupMobilePhone ...寻找官方客服电话号码
func (bb *bankBrain) pickupCustomerPhone(msg string) (string, bool) {
	matchIndex := bb.phoneACMatch.Match(message)
	fmt.Println("======", matchIndex)
	if len(matchIndex) > 0 {
		return bb.customerPhone[matchIndex[0]], true
	}
	return "", false
}

//pickupMobilePhone ...寻找手机号码
func (bb *bankBrain) pickupMobilePhone(msg string) (string, bool) {
	return utils.ExtractMobilePhone(msg)
}

//MatchScore ...计算匹配分数，分支越高，可信度越高。
func (bb *bankBrain) MatchScore(pickup propertiesVec) float64 {

	return 0.0
}

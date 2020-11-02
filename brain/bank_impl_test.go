package brain

import (
	"testing"
)

func Test_pickupProperties(t *testing.T) {

	util := CreateBankBrain()
	t.Log(util)

	matchIndex := util.acMatch.Match(message)
	t.Log(matchIndex)

	//识别银行名称
	name, ok := util.pickupName(message)
	t.Log(ok, name)

	//识别域名
	name, ok = util.pickupWebDomain(message)
	t.Log(ok, name)

	//识别手机电话
	name, ok = util.pickupMobilePhone(message)
	t.Log(ok, name)

	name, ok = util.pickupCustomerPhone("光大咨询95599")
	t.Log(name, ok)

}

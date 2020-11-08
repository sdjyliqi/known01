package brain

import (
	"testing"
)

func Test_pickupProperties(t *testing.T) {

	util := CreateBankBrain()
	t.Log(util)

	matchIndex := util.acMatch.Match(messageTest)
	t.Log(matchIndex)

	//识别银行名称
	name, ok := util.pickupName(messageTest)
	t.Log(ok, name)

	//识别域名
	name, ok = util.pickupWebDomain(messageTest)
	t.Log(ok, name)

	//识别手机电话
	name, ok = util.pickupMobilePhone(messageTest)
	t.Log(ok, name)

}

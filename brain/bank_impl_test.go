package brain

import (
	"github.com/stretchr/testify/assert"
	"known01/models"
	"known01/utils"
	"testing"
)

var messageTest = "[招商银行]尊敬的客户，一张闪电贷专属礼券为你呈上！用券条款可享受专属利率优惠，优惠日截止2020年10月31日。" +
	"点击http://a.cmbchina.com/personal/cmhkas13快速申请，详情请咨询95599,400-66666888,15210510285"

func Test_pickupProperties(t *testing.T) {
	var messageTest = "[工行]尊敬的客户，一张闪电贷专属礼券为你呈上！用券条款可享受专属利率优惠，优惠日截止2020年10月31日,点击http://a.cmbchina.com/personal/cmhkas13快速申请，详情请咨询95599"
	items, err := models.Reference{}.GetItems(utils.GetMysqlClient())
	assert.Nil(t, err)
	util := CreateBankBrain(items)
	t.Log(util)

	matchIndex := util.acMatch.Match(messageTest)
	t.Log(matchIndex)

	//识别银行名称
	name, ok := util.pickupName(messageTest)
	t.Log("pickup name :", ok, name)

	//识别域名
	name, ok = util.pickupWebDomain(messageTest)
	t.Log(ok, name)

	//识别手机电话
	name, ok = util.pickupMobilePhone(messageTest)
	t.Log(ok, name)

}

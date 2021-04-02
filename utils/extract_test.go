package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ExtractMPhone(t *testing.T) {
	//for valid phone num
	phone := "ABCD15210510285 "
	v, ok := ExtractMobilePhone(phone)
	t.Log(v, ok)
	assert.True(t, ok)
	assert.Equal(t, v, "15210510285")

	//for valid phone num
	phone = "我的手机+8615210510285"
	v, ok = ExtractMobilePhone(phone)
	t.Log(v, ok)
	assert.True(t, ok)
	assert.Equal(t, v, "15210510285")

	//for invalid phone num
	phone = "152105"
	v, ok = ExtractMobilePhone(phone)
	t.Log(v, ok)
	assert.False(t, ok)
}

func Test_ChkContentIsMobilePhone(t *testing.T) {
	txt := "15210510285"
	ok := ChkContentIsMobilePhone(txt)
	assert.True(t, ok)

	txt = "+8615210510285"
	ok = ChkContentIsMobilePhone(txt)
	assert.True(t, ok)

	txt = "1521+0510285"
	ok = ChkContentIsMobilePhone(txt)
	assert.False(t, ok)

	txt = "1521+0510285"
	ok = ChkContentIsMobilePhone(txt)
	assert.False(t, ok)

	txt = "ok15210510285"
	ok = ChkContentIsMobilePhone(txt)
	assert.False(t, ok)
}

func Test_ExtractWeb(t *testing.T) {
	txt := "https://www.cmbchina.com/"
	v, ok := ExtractWebDomain(txt)
	t.Log(v, ok)

	txt = "cmbt.cn/uuY"
	v, ok = ExtractWebDomain(txt)
	t.Log(v, ok)
}

func Test_ExtractWebFirstDomain(t *testing.T) {
	txt := "请登录https://www.cmbchina.com/"
	v, ok := ExtractWebFirstDomain(txt)
	assert.True(t, ok)
	assert.Equal(t, "cmbchina", v)

	txt = "请登录cmbt.cn/uuY"
	v, ok = ExtractWebFirstDomain(txt)
	assert.True(t, ok)
	assert.Equal(t, "cmbt", v)

	txt = "请登录http://mail.cmbt.cn/uuY"
	v, ok = ExtractWebFirstDomain(txt)
	assert.True(t, ok)
	assert.Equal(t, "cmbt", v)

	txt = "请登录abc.cdf.hjk"
	v, ok = ExtractWebFirstDomain(txt)
	assert.False(t, ok)

	txt = "立即抢购请戳www.cebwn.com。联系电话95595理财非存款"
	v, ok = ExtractWebFirstDomain(txt)
	t.Log(v)
	assert.True(t, ok)
}

func TestPickTelephone(t *testing.T) {
	var tests = []struct {
		input  string
		output bool
	}{
		{"尊敬的王先生,您近期违规用卡逾期未还,我行于今日冻结您名下信用卡进黑名单,如需恢复正常请致电:031100797943【招商银行】", false},
		{"截止今晚24点我行将自动从您银行卡上扣除年费1200元。如有疑问,咨询电话:08776613854-08776615821 【农行通知】", true},
		{"尊敬的储蓄卡用户:您于本月18日在广晟百货用卡购买电器9886元,此款将从您帐上扣,如有问题请联 系0762-3926317 工商银行。", true},
		{"温馨提示:我行将于20: 00之前扣除您信用卡年费1280元,如有疑问详情请致电农行客服中心4008530010 【中国农业银行】", false},
		{"尊敬的用户:您的电子密码器即将失效,请尽快登录手机银行http//wap.95588op.com/升级维护01081234567232387给您带来不便敬请谅解《工商银行》", true},
		{"尊敬的用户:您的电子密码器即将失效,请尽快登录手机银行http//wap.95588op.com/升级维护01091234567给您带来不便敬请谅解《工商银行》", false},
		{"尊敬的工行用户:您的账户积分58513即将逾期清空,请登录手机网wap.noicco.com兑换988.5元现金,逾期失效【工商银行】", false},
	}
	for _, test := range tests {
		output1 := PickTelephone(test.input)
		t.Log(output1)
		bo := output1 == ""
		if bo == test.output {
			t.Errorf("短信内容为\"%s\"，座机号为：%s\n", test.input, output1)
		}
	}
}

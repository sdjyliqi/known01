package brain

import (
	"github.com/go-playground/assert/v2"
	"known01/utils"
	"testing"
)

func Test_acFindPhoneNum(t *testing.T) {
	c := CreateCenter()
	v, ok := c.acFindPhoneID(messageTest)
	t.Log(v, ok)
}

func Test_amendMessage(t *testing.T) {
	c := CreateCenter()
	testMsg := "爸爸去哪里呀"
	v := c.amendMessage(testMsg)
	t.Log(v)
	assert.Equal(t, "去哪里呀", v)
}

func Test_matchEngineRate(t *testing.T) {
	c := CreateCenter()
	testMsg := "尊敬的用户：您的电子密码器于次日失效，请速登录手机维护网站wap.icbcsap.com进行更新。给你带来不变，敬请谅解！【工商银行】"
	rate, engine := c.matchEngineRate(testMsg)
	t.Log(rate, engine)
}

func Test_acFindIndexWord(t *testing.T) {
	c := CreateCenter()
	testMsg := "尊敬的用户：给我转点钱吧，敬请谅解！【工商银行】"
	v, ok := c.acFindIndexWord(testMsg)
	t.Log(v, ok)
	assert.Equal(t, true, ok)
}

func Test_GetEngineName(t *testing.T) {
	c := CreateCenter()
	//测试包括电话号码的情况
	testMsg := "尊敬的用户：给我转点钱吧，敬请谅解！95595"
	name, phoneID := c.GetEngineName(testMsg)
	t.Log(name, phoneID)
	assert.Equal(t, utils.EngineBank, name)
	assert.Equal(t, "95595", phoneID)
	//测试匹配模板的
	testMsg = "尊敬的客户，一张闪电贷专属礼券为你呈上！用券条款可享受专属利率优惠，优惠日截止年月日。"
	name, phoneID = c.GetEngineName(testMsg)
	t.Log(name, phoneID)
	assert.Equal(t, utils.EngineBank, name)
	assert.Equal(t, "", phoneID)

	//测试匹配关键字
	testMsg = "爸爸，给我转账1000元，着急用，账号：110000000000"
	name, phoneID = c.GetEngineName(testMsg)
	t.Log(name, phoneID)
	assert.Equal(t, utils.EngineBank, name)
	assert.Equal(t, "", phoneID)

}

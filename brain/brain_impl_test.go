package brain

import (
	"github.com/go-playground/assert/v2"
	"github.com/sdjyliqi/known01/testutils"
	"github.com/sdjyliqi/known01/utils"
	"testing"
)

func Test_acFindPhoneNum(t *testing.T) {
	testutils.Init()
	c := CreateCenter()
	v, ok := c.acFindPhoneID(messageTest)
	t.Log(v, ok)
}

func Test_amendMessage(t *testing.T) {
	testutils.Init()
	c := CreateCenter()
	testMsg := "爸爸我在外面北京遇到点事，不方便接听电话和短信，给我打10000元到我的工商银行，户主李奇，账号12121212121111"
	v := c.amendMessage(testMsg)
	t.Log(v)

	testMsg = "尊敬的工商银行客户，一张闪电贷专属礼券为你呈上！用券条款可享受专属利率优惠，优惠日截止2020年11月30日，请点击wap.cmbc188.com登录招商银行官网查看【工商银行】"
	v = c.amendMessage(testMsg)
	t.Log(v)
}

func Test_matchEngineRate(t *testing.T) {
	testutils.Init()
	c := CreateCenter()
	testMsg := "尊敬的用户：您的电子密码器于次日失效，请速登录手机维护网站wap.icbcsap.com进行更新。给你带来不变，敬请谅解,,具体请咨询95588！【工商银行】"
	rate, engine := c.matchEngineRate(testMsg)
	t.Log(rate, engine)
}

func Test_acFindIndexWord(t *testing.T) {
	testutils.Init()
	c := CreateCenter()
	testMsg := "尊敬的用户：给我转点钱吧，敬请谅解！【工商银行】"
	v, ok := c.acFindIndexWord(testMsg)
	t.Log(v, ok)
	assert.Equal(t, true, ok)
}

func Test_GetEngineName(t *testing.T) {
	testutils.Init()
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

func Test_JudgeMessage(t *testing.T) {
	testutils.Init()
	testMsg := "尊敬的用户：您的电子密码器于次日失效，请速登录手机维护网站wap.icbc.com进行更新。给你带来不变，敬请谅解,具体请咨询95588！【工商银行】"
	c := CreateCenter()
	score, reference := c.JudgeMessage(testMsg, "95588")
	t.Log(score, reference)

	testMsg = "今天阳光明媚，可以小酌一杯"
	score, reference = c.JudgeMessage(testMsg, "15210510285")
	t.Log(score, reference)

	testMsg = "恭喜！您的手机号码15210510288已经被栏目组随机抽取为场外的幸运用户。将获得由赞助商提供的奖金:价值元的笔记本电脑！详情请登录活动网站:验证码请牢记[中国好声音]"
	score, reference = c.JudgeMessage(testMsg, "800-12344444")
	t.Log(score, reference)
}

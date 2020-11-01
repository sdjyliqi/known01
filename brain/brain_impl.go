package brain

import (
	"fmt"
	"known01/utils"
)

func (c *Center) init() bool {
	//load templates about bank
	ok := c.loadBankTemplates([]string{})
	if !ok {
		return false
	}
	//load templates about bank
	ok = c.loadExpressTemplates([]string{})
	if !ok {
		return false
	}

	return true
}

func (c *Center) loadBankTemplates(templates []string) bool {
	templates = []string{"尊敬的工商银行用户：您的电子密码器于次日失效，请速登录手机维护网站wap.icbcsap.com进行更新。给你带来不变，敬请谅解！【工商银行】"}
	templates = append(templates, "尊敬的工商银行用户：您的积分将在近期失效，请速登录手机维护网站wap.icbcsap.com进行更新。给你带来不变，敬请谅解！【工商银行】")
	fmt.Println(templates)
	c.bankTemplates = templates
	return true
}

func (c *Center) loadExpressTemplates(templates []string) bool {
	return true
}

//GetEngineName ... 根据提交的信息，判断最符合那个鉴别引擎
func (c *Center) GetEngineName(msg string) EngineType {
	return EngineBank
}

//bankEngineRate ...计算银行类信息匹配度
func (c *Center) bankEngineRate(msg string) float64 {
	matchRate := 0.0
	for _, v := range c.bankTemplates {
		rate := utils.SimilarDegree(msg, v)
		if rate > matchRate {
			matchRate = rate
		}
	}
	return 0.0
}

func (c *Center) expressEngineRate(msg string) float64 {
	return 0.0
}

func (c *Center) rewardEngineRate(msg string) float64 {
	return 0.0
}

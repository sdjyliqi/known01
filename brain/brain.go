package brain

//定义银行判断大脑
const (
	EngineBank    = 0
	EngineExpress = 1
	EngineReward  = 2
)

type EngineType int



type Center struct {
	keywords         []string
	bankTemplates    []string    //银行类短信模块内容列表
	ExpressTemplates []string    //快递类短信模块内容列表
	RewardTemplates  []string    //中奖类短信模块内容列表
	bankKeywords     []string    //银行类关键词列表
	ExpressKeywords  []string    //快递类关键词列表
	RewardKeywords   []string    //中奖类关键词列表
}
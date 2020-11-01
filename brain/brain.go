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
	bankTemplates    []string //银行类短信模块内容列表
	ExpressTemplates []string //快递类短信模块内容列表
	RewardTemplates  []string //中奖类短信模块内容列表
	bankKeywords     []string //银行类关键词列表
	ExpressKeywords  []string //快递类关键词列表
	RewardKeywords   []string //中奖类关键词列表
}

//propertiesVec... 定义文本识别提取数据维度
type propertiesVec struct {
	govName       string //特征名称，如中国工商银行
	website       string //域名，目前只匹配一级域
	mobilePhone   string //手机号码
	customerPhone string //客服电话，如400-***-******
	senderID      string //发信者号码，如95599，可能和客服电话不一致
	fixedPhone    string //固定电话，如010-88888888
}

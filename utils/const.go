package utils

/*
   定义三个引擎的编号
   20201105 今天来暖气了，暖洋洋
*/
type EngineType int

const (
	EngineUnknown EngineType = 0
	EngineBank    EngineType = 1
	EngineReward  EngineType = 2

	PageEntry         = 10
	ScoreSenderMobile = -10
	ScoreFindMobile   = -5

	OutsideKnown = -99999 //如果不支持，返回该值
)

var SplitChar string = ","

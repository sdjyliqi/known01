package utils

/*
   定义三个引擎的编号
   20201105 今天来暖气了，暖洋洋
*/
type EngineType int

const (
	ENGINE_UNKNOWN EngineType = 0
	ENGINE_BANK    EngineType = 1
	ENGINE_REWARD  EngineType = 2

	PAGE_ENTRY          = 10
	SCORE_SENDER_MOBILE = -10
	SCORE_FIND_MOBILE   = -5
)

var SplitChar string = ","

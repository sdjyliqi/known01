package utils

import (
	"fmt"
	"github.com/huichen/sego"
	"testing"
)

func Test_CutWords(t *testing.T) {
	InitSegDic()
	msg := "向我光大银行转账100元"
	segments := segmenter.Segment([]byte(msg))

	t.Log(sego.SegmentsToString(segments, false))

	str := "爸爸，向我的光大银行转100元吧，现在有点着急不方便回信息。"
	str2 := "爸，现在有点着急，不方便回信息，给我光大银行转100元吧。"
	//str3:="nai nai ge xiong cao"
	//str hash 值
	hash := SimHashTool.Simhash(str)
	fmt.Println(str, hash)
	//str2 距离
	hash2 := SimHashTool.Simhash(str2)
	fmt.Println(str2, hash2)
	//计算相似度
	sm := SimHashTool.Similarity(hash, hash2)
	fmt.Println("similar score:", sm)
	//距离

}

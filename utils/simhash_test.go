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
	str2 := "现在有点事，不方便回信息，给我光大银行转100元吧，爸爸。"

	hash := SimHashTool.Simhash(str)
	fmt.Println(str, hash)
	//str2 距离
	hash2 := SimHashTool.Simhash(str2)
	fmt.Println(str2, hash2)
	//计算相似度
	sm := SimHashTool.Similarity(hash, hash2)
	fmt.Println("similar score:", sm)

	//距离
	str1 := "尊敬的李先生，光大银行送福利，线上刮大奖，瓜分百万红包，快戳www.cebbank.com刮奖"
	str2 = "光大银行送福利啦，线上刮大奖，瓜分百万红包，快戳www.cebbank.com刮奖"
	hash = SimHashTool.Simhash(str1)
	fmt.Println(str1, hash)
	//str2 距离
	hash2 = SimHashTool.Simhash(str2)
	fmt.Println(str2, hash2)
	//计算相似度
	sm = SimHashTool.Similarity(hash, hash2)
	fmt.Println("similar score:", sm)

}

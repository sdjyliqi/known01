package utils

import (
	"fmt"
	"math"
	"testing"
)



//min 获取小值
func min(a ...int) int {
	r := a[0]
	for _, v := range a {
		if r > v {
			r = v
		}
	}
	return r
}
//func SimilarText(first, second string, percent *float64) int {



func Test_simnet(t *testing.T) {
	fmt.Println(SimilarDegree("你好世界", "晚安世界"))
}



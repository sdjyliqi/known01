package utils

import (
	"fmt"
	"testing"
)

//func SimilarText(first, second string, percent *float64) int {

func Test_simnet(t *testing.T) {
	fmt.Println(SimilarDegree("你好世界", "晚安世界"))
}

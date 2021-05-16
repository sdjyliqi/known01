package handle

import (
	"known01/model"
	"testing"
	"time"
)

func TestDataValidation(t *testing.T) {
	var tests = []model.Reference{
		{10, "云闪付", 1, "银联", "95568,4008695568,4008695568", "95568", "95123", "www.dwdwd.cn", "dwdwd", time.Now()},
	}
	var outputs = []bool{
		true,
	}
	for i, test := range tests {
		bool1, res := DataValidation(test)
		if outputs[i] != bool1 {
			t.Errorf("校验结果为%t，错误信息为：%s\n", bool1, res)
		}
	}
}

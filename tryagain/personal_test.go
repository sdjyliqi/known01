package tryagain

import (
	"testing"
)

func Test_PhoneID(t *testing.T) {
	text := "我0.15210510987们在使14012346789用分列功能进1521051028578行出生日期提取之前，我手机号15210510285,0.115210510285，   18701516837开心"
	v := ExtractMobilePhoneDs(text)
	t.Log(v)
}

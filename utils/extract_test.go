package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)




func Test_ExtractMPhone(t *testing.T) {
	//for valid phone num
	phone :="ABCD15210510285 "
	v,ok := ExtractMPhone(phone,PhoneFormat)
	t.Log(v,ok)
	assert.True(t,ok)
	assert.Equal(t,v,"15210510285")

	//for valid phone num
	phone ="我的手机+8615210510285"
	v,ok = ExtractMPhone(phone,PhoneFormat)
	t.Log(v,ok)
	assert.True(t,ok)
	assert.Equal(t,v,"15210510285")

	//for invalid phone num
	phone ="152105"
	v,ok = ExtractMPhone(phone,PhoneFormat)
	t.Log(v,ok)
	assert.False(t,ok)

	//for 官方客服
	phone ="请咨询4008895555"
	v,ok = ExtractMPhone(phone,CustomerFormat)
	t.Log(v,ok)
	assert.True(t,ok)
	assert.Equal(t,v,"4008895555")

	phone ="请致电800-830-8855"
	v,ok = ExtractMPhone(phone,CustomerFormat)
	t.Log(v,ok)
	assert.True(t,ok)
	assert.Equal(t,v,"8008308855")
}


func Test_ExtractWeb(t *testing.T) {
	txt := "https://www.cmbchina.com/"
	v,ok := ExtractWebDomain(txt)
	t.Log(v,ok)

	txt  = "cmbt.cn/uuY"
	v,ok  = ExtractWebDomain(txt)
	t.Log(v,ok)
}




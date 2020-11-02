package utils

import (
	"regexp"
	"strings"
)

//手机号码格式
var PhoneFormat = "(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}"

//客服号码格式 ...400-699-5555
var CustomerFormat = "(400|800)\\d{7}"
var WebFormat = "(http|https)://[a-z0-9\\.]+"
var shortWebFormat = "[a-z0-9\\.]{2,12}.(cn|com)"

//ExtractMPhone ..提取手机号

func ExtractMobilePhone(txt string) (string, bool) {
	txt = strings.Replace(txt, "-", "", -1)
	txt = strings.Replace(txt, "+86", "", -1)
	phoneRegx := regexp.MustCompile(PhoneFormat)
	phoneNums := phoneRegx.FindStringSubmatch(txt)
	if len(phoneNums) > 1 {
		return phoneNums[0], true
	}
	return "", false
}

//ExtractWebDomain ..提取登录网址
func ExtractWebDomain(txt string) (string, bool) {
	txt = strings.ToLower(txt)
	formatRegx := regexp.MustCompile(WebFormat)
	values := formatRegx.FindStringSubmatch(txt)
	if len(values) >= 1 {
		return values[0], true
	}

	formatRegx = regexp.MustCompile(shortWebFormat)
	values = formatRegx.FindStringSubmatch(txt)
	if len(values) >= 1 {
		return values[0], true
	}
	return "", false
}

func ExtractWebFirstDomain(txt string) (string, bool) {
	website, ok := ExtractWebDomain(txt)
	if !ok {
		return "", false
	}
	startIdx, idxDCn := 0, 0
	idxDCn = strings.Index(website, ".com")
	if idxDCn < 0 {
		idxDCn = strings.Index(website, ".cn")
	}
	if idxDCn <= 0 {
		return "", false
	}
	startIdx = strings.LastIndexByte(website[0:idxDCn], '.')
	return website[startIdx+1 : idxDCn], true
}

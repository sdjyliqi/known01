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
var lenMobilePhone = 11

var AreaCode = map[string]bool{
	"010":  true,
	"020":  true,
	"021":  true,
	"022":  true,
	"023":  true,
	"024":  true,
	"025":  true,
	"027":  true,
	"028":  true,
	"029":  true,
	"0310": true,
	"0311": true,
	"0312": true,
	"0313": true,
	"0314": true,
	"0315": true,
	"0316": true,
	"0317": true,
	"0318": true,
	"0319": true,
	"0335": true,
	"0349": true,
	"0350": true,
	"0351": true,
	"0352": true,
	"0353": true,
	"0354": true,
	"0355": true,
	"0356": true,
	"0357": true,
	"0358": true,
	"0359": true,
	"0370": true,
	"0371": true,
	"0372": true,
	"0373": true,
	"0374": true,
	"0375": true,
	"0376": true,
	"0377": true,
	"0379": true,
	"0391": true,
	"0392": true,
	"0393": true,
	"0394": true,
	"0395": true,
	"0396": true,
	"0398": true,
	"0411": true,
	"0412": true,
	"0415": true,
	"0416": true,
	"0417": true,
	"0418": true,
	"0419": true,
	"0421": true,
	"0427": true,
	"0429": true,
	"0431": true,
	"0432": true,
	"0433": true,
	"0434": true,
	"0435": true,
	"0436": true,
	"0437": true,
	"0438": true,
	"0439": true,
	"0451": true,
	"0452": true,
	"0453": true,
	"0454": true,
	"0455": true,
	"0456": true,
	"0457": true,
	"0458": true,
	"0459": true,
	"0464": true,
	"0467": true,
	"0468": true,
	"0469": true,
	"0470": true,
	"0471": true,
	"0472": true,
	"0473": true,
	"0474": true,
	"0475": true,
	"0476": true,
	"0477": true,
	"0478": true,
	"0479": true,
	"0482": true,
	"0510": true,
	"0511": true,
	"0512": true,
	"0513": true,
	"0514": true,
	"0515": true,
	"0516": true,
	"0517": true,
	"0518": true,
	"0519": true,
	"0523": true,
	"0527": true,
	"0530": true,
	"0531": true,
	"0532": true,
	"0533": true,
	"0534": true,
	"0535": true,
	"0536": true,
	"0537": true,
	"0538": true,
	"0539": true,
	"0543": true,
	"0546": true,
	"0550": true,
	"0551": true,
	"0552": true,
	"0553": true,
	"0554": true,
	"0555": true,
	"0556": true,
	"0557": true,
	"0558": true,
	"0559": true,
	"0561": true,
	"0562": true,
	"0563": true,
	"0564": true,
	"0566": true,
	"0570": true,
	"0571": true,
	"0572": true,
	"0573": true,
	"0574": true,
	"0575": true,
	"0576": true,
	"0577": true,
	"0578": true,
	"0579": true,
	"0580": true,
	"0591": true,
	"0592": true,
	"0593": true,
	"0594": true,
	"0595": true,
	"0596": true,
	"0597": true,
	"0598": true,
	"0599": true,
	"0631": true,
	"0632": true,
	"0633": true,
	"0634": true,
	"0635": true,
	"0660": true,
	"0662": true,
	"0663": true,
	"0668": true,
	"0691": true,
	"0692": true,
	"0701": true,
	"0710": true,
	"0711": true,
	"0712": true,
	"0713": true,
	"0714": true,
	"0715": true,
	"0716": true,
	"0717": true,
	"0718": true,
	"0719": true,
	"0722": true,
	"0724": true,
	"0728": true,
	"0730": true,
	"0731": true,
	"0734": true,
	"0735": true,
	"0736": true,
	"0737": true,
	"0738": true,
	"0739": true,
	"0743": true,
	"0744": true,
	"0745": true,
	"0746": true,
	"0750": true,
	"0751": true,
	"0752": true,
	"0753": true,
	"0754": true,
	"0755": true,
	"0756": true,
	"0757": true,
	"0758": true,
	"0759": true,
	"0760": true,
	"0762": true,
	"0763": true,
	"0766": true,
	"0768": true,
	"0769": true,
	"0770": true,
	"0771": true,
	"0772": true,
	"0773": true,
	"0774": true,
	"0775": true,
	"0776": true,
	"0777": true,
	"0778": true,
	"0779": true,
	"0790": true,
	"0791": true,
	"0792": true,
	"0793": true,
	"0794": true,
	"0795": true,
	"0796": true,
	"0797": true,
	"0798": true,
	"0799": true,
	"0812": true,
	"0813": true,
	"0816": true,
	"0817": true,
	"0818": true,
	"0825": true,
	"0826": true,
	"0827": true,
	"0830": true,
	"0831": true,
	"0832": true,
	"0833": true,
	"0834": true,
	"0835": true,
	"0836": true,
	"0837": true,
	"0838": true,
	"0839": true,
	"0851": true,
	"0854": true,
	"0855": true,
	"0856": true,
	"0857": true,
	"0858": true,
	"0859": true,
	"0870": true,
	"0871": true,
	"0872": true,
	"0873": true,
	"0874": true,
	"0875": true,
	"0876": true,
	"0877": true,
	"0878": true,
	"0879": true,
	"0883": true,
	"0886": true,
	"0887": true,
	"0888": true,
	"0891": true,
	"0892": true,
	"0893": true,
	"0894": true,
	"0895": true,
	"0896": true,
	"0897": true,
	"0898": true,
	"0901": true,
	"0902": true,
	"0903": true,
	"0906": true,
	"0908": true,
	"0909": true,
	"0911": true,
	"0912": true,
	"0913": true,
	"0914": true,
	"0915": true,
	"0916": true,
	"0917": true,
	"0919": true,
	"0930": true,
	"0931": true,
	"0932": true,
	"0933": true,
	"0934": true,
	"0935": true,
	"0936": true,
	"0937": true,
	"0938": true,
	"0939": true,
	"0941": true,
	"0943": true,
	"0951": true,
	"0952": true,
	"0953": true,
	"0954": true,
	"0955": true,
	"0970": true,
	"0971": true,
	"0972": true,
	"0973": true,
	"0974": true,
	"0975": true,
	"0976": true,
	"0977": true,
	"0979": true,
	"0990": true,
	"0991": true,
	"0992": true,
	"0993": true,
	"0994": true,
	"0995": true,
	"0996": true,
	"0997": true,
	"0998": true,
	"0999": true,
}

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

func ChkContentIsMobilePhone(txt string) bool {
	txt = strings.Replace(txt, "+86", "", 1)
	if len(txt) != lenMobilePhone {
		return false
	}
	_, ok := ExtractMobilePhone(txt)
	return ok
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

//PickTelephone   ...提取一段话中的座机号，输入字符串str，返回值中bool为是否还有座机号，string为提取的内容
func PickTelephone(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "-", "", -1)
	regTelephone := `(?:\+?86 ?)?0[0-9]{2}[0-9\- ]{2}[0-9]{6,8}`
	reg := regexp.MustCompile(regTelephone)
	telephoneList := reg.FindAllString(str, -1)
	if telephoneList == nil {
		return ""
	}
	telephoneList[0] = strings.Replace(telephoneList[0], " ", "", -1)
	telephoneList[0] = strings.Replace(telephoneList[0], "-", "", -1)
	if AreaCode[telephoneList[0][0:3]] == true {
		if telephoneList[0][3:4] != "0" && telephoneList[0][3:4] != "1" && telephoneList[0][3:4] != "9" {
			return telephoneList[0]
		}
		return ""
	}
	if AreaCode[telephoneList[0][0:4]] == true {
		if telephoneList[0][3:4] != "0" && telephoneList[0][3:4] != "1" && telephoneList[0][3:4] != "9" {
			return telephoneList[0]
		}
	}
	return ""
}

//PickupHits...从待鉴别诈骗短信中提取核心信息，如【中国移动】话费余额已不足。。。。，需要提取中国移动
func PickupHits(msg string) string {
	regHit := regexp.MustCompile(`【(.{2,8})】`)
	hits := regHit.FindStringSubmatch(msg)
	if len(hits) > 1 {
		return hits[1]
	}
	return ""

}

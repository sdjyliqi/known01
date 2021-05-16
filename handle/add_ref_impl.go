package handle

import (
	"known01/model"
	"known01/utils"
	"regexp"
	"sort"
	"strings"
)

func DataValidation(data model.Reference) (bool, string) {
	//首先对基准数据本身进行校验
	//对name字段进行校验，判断是否全是中文字符、数字
	nameBool, _ := regexp.MatchString("[^\u4e00-\u9fa50-9]", data.Name)
	if nameBool {
		return !nameBool, "name字段只能由中文和数字组成"
	}
	//对alias_name字段进行校验，判断是否全是中文字符、数字和英文逗号','
	aliasnameBool, _ := regexp.MatchString("[^\u4e00-\u9fa50-9,]", data.AliasNames)
	if data.AliasNames != "" && aliasnameBool {
		return !aliasnameBool, "alias_name字段只能为空值、中文、数字和英文逗号','"
	}
	//判断name字段是否在alias_name字段中
	if strings.Contains(data.AliasNames, data.Name) {
		return false, "name字段不能在alias_name字段中重复出现"
	}

	//对phone字段进行校验，判断是否全是数字和英文逗号','
	phoneBool, _ := regexp.MatchString("[^0-9,]", data.Phone)
	if phoneBool {
		return !phoneBool, "phone字段只能由数字和英文逗号','组成"
	}
	//对phone字段中的每个元素都检验，看是否符合号码规则
	phoneList := strings.Split(data.Phone, ",")
	phonePrompt := ""
	for _, phoneElement := range phoneList {
		telResult := utils.PickTelephone(phoneElement)
		if telResult == "" && !ValidateShortNumber(phoneElement) && !ValidateShortTelephone(phoneElement) &&
			!Validate400Number(phoneElement) {
			phonePrompt = phoneElement + ","
		}
	}
	if phonePrompt != "" {
		return false, "phone字段中" + phonePrompt + "格式错误"
	}
	//对sender_id字段进行校验
	senderIdBool, _ := regexp.MatchString("[^0-9,]", data.SenderId)
	if data.SenderId != "" && senderIdBool {
		return !senderIdBool, "alias_name字段只能为空值和数字"
	}
	//对sender_id字段中的每个元素都检验，看是否符合号码规则
	if data.SenderId != "" {
		senderIdList := strings.Split(data.SenderId, ",")
		senderIdPrompt := ""
		for _, senderIdElement := range senderIdList {
			telResult := utils.PickTelephone(senderIdElement)
			if telResult == "" && !ValidateShortNumber(senderIdElement) && !ValidateSmsLongNumber(senderIdElement) {
				senderIdPrompt = senderIdElement + ","
			}
		}
		if senderIdPrompt != "" {
			return false, "sender_id字段中" + senderIdPrompt + "请检查sender_id字段号码格式"
		}
	}

	//对manual_phone字段进行校验
	manualPhoneBool, _ := regexp.MatchString(`[^0-9\-]`, data.ManualPhone)
	if manualPhoneBool {
		return !manualPhoneBool, "manual_phone字段只能由数字和-组成"
	}
	manualPhoneResult := utils.PickTelephone(data.ManualPhone)
	if manualPhoneResult == "" && !ValidateShortNumber(data.ManualPhone) && !ValidateShortTelephone(data.ManualPhone) &&
		!Validate400Number(data.ManualPhone) {
		return false, "请检查manual_phone字段号码格式"
	}
	//对website字段进行校验
	websiteBool, _ := regexp.MatchString(`[^a-z0-9:.\-/]`, data.Website)
	if websiteBool {
		return !websiteBool, "website字段只能由数字、字母、:、.、-和/组成"
	}

	//对domain字段进行校验
	domainBool, _ := regexp.MatchString("[^a-z0-9,]", data.Domain)
	if domainBool {
		return !domainBool, "domain字段只能为数字、字母和英文逗号','"
	}

	//判断更新的name、alias_name、manualphone、website字段在哪数据库中是否已经存在
	items, err := model.Reference{}.GetItems(utils.GetMysqlClient())
	if err != nil {
		return false, "从数据库获取数据失败"
	}
	var existName []string
	var existManualPhone []string
	var existWebsite []string
	for _, v := range items {
		existName = append(existName, v.Name)
		localManualPhone := strings.Split(v.ManualPhone, ",")
		localAliasName := strings.Split(v.AliasNames, ",")
		existName = append(existName, localAliasName...)
		existName = append(existName, localManualPhone...)
		existManualPhone = append(existManualPhone, v.ManualPhone)
		existWebsite = append(existWebsite, v.Website)
	}
	//检验name字段是否已经在数据库中存在
	if validateString(data.Name, existName) {
		return false, "name字段已经存在"
	}
	//检验alias_names字段是否已经在数据库中存在
	if data.AliasNames != "" {
		aliasNameList := strings.Split(data.AliasNames, ",")
		aliasNamePrompt := ""
		for _, aliasNameElement := range aliasNameList {
			if validateString(aliasNameElement, existName) {
				aliasNamePrompt = aliasNameElement + ","
			}
		}
		if aliasNamePrompt != "" {
			return false, "alias_names字段中" + aliasNamePrompt + "已经存在"
		}
	}
	//检验manualphone字段是否已经在数据库中存在
	if validateString(data.ManualPhone, existManualPhone) {
		return false, "manualphone字段已经存在"
	}
	//检验website字段是否已经在数据库中存在
	if validateString(data.Website, existWebsite) {
		return false, "Website字段已经存在"
	}

	return true, "校验成功"
}

//ValidateShortNumber   ...校验是否为95、96和10开头的短号
func ValidateShortNumber(str string) bool {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "-", "", -1)
	regShortNumber := `\A(95|96|10)[0-9]{3,6}`
	reg := regexp.MustCompile(regShortNumber)
	if reg.MatchString(str) {
		return true
	}
	return false
}

//ValidateShortTelephone   ...校验是否是区号+96开头短号
func ValidateShortTelephone(str string) bool {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "-", "", -1)
	regShortTelephone := `\A(96)[0-9]{3-4}`
	reg := regexp.MustCompile(regShortTelephone)
	if utils.AreaCode[str[0:3]] {
		if reg.MatchString(str[3:]) {
			return true
		}
	}
	if utils.AreaCode[str[0:4]] {
		if reg.MatchString(str[4:]) {
			return true
		}
	}
	return false
}

//Validate400Number   ...校验400开头的号码
func Validate400Number(str string) bool {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "-", "", -1)
	reg400 := `\A(400)[0-9]{7}`
	reg := regexp.MustCompile(reg400)
	if reg.MatchString(str) {
		return true
	}
	return false
}

//ValidateSmsLongNumber ...匹配10开头的短信长号码，号码位数8-19
func ValidateSmsLongNumber(str string) bool {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "-", "", -1)
	regLongNumber := `\A(10)[06][0-9]{5,16}`
	reg := regexp.MustCompile(regLongNumber)
	if reg.MatchString(str) {
		return true
	}
	return false
}

//validateString   ...判断字符串是否在列表中
func validateString(target string, strArray []string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

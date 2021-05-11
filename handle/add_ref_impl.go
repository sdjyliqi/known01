package handle

import (
	"known01/model"
	"known01/utils"
	"regexp"
	"strings"
)

func DataValidation(data model.Reference) (bool, string) {
	//首先对基准数据本身进行校验
	//对name字段进行校验，判断是否全是中文字符、数字
	nameBool, _ := regexp.MatchString("[^\u4e00-\u9fa50-9]", data.Name)
	if nameBool {
		return !nameBool, "name字段只能为中文和数字"
	}
	//对alias_name字段进行校验，判断是否全是中文字符、数字和英文逗号','
	aliasnameBool, _ := regexp.MatchString("[^\u4e00-\u9fa50-9,]", data.AliasNames)
	if data.AliasNames != "" && aliasnameBool {
		return !aliasnameBool, "alias_name字段只能为空值、中文、数字和英文逗号','"
	}

	//对phone字段进行校验，判断是否全是数字和英文逗号','
	phoneBool, _ := regexp.MatchString("[^0-9,]", data.Phone)
	if phoneBool {
		return !phoneBool, "phone字段只能为数字和英文逗号','"
	}
	//对phone字段中的每个元素都检验，看是否符合号码规则
	phoneList := strings.Split(data.Phone, ",")
	phonePrompt := ""
	for _, phoneElement := range phoneList {
		telResult := utils.PickTelephone(phoneElement)
		if telResult == "" || !ValidateShortNumber(phoneElement) || !ValidateShortTelephone(phoneElement) ||
			!Validate400Number(phoneElement) || !ValidateSmsLongNumber(phoneElement) {
			phonePrompt = phoneElement + ","
		}
	}
	if phonePrompt != "" {
		return false, "请检查phone字段号码格式"
	}
	//对sender_id字段进行校验
	senderIdBool, _ := regexp.MatchString("[^0-9]", data.SenderId)
	if data.SenderId != "" && senderIdBool {
		return !aliasnameBool, "alias_name字段只能为空值和数字"
	}

	//对manual_phone字段进行校验
	manualPhoneBool, _ := regexp.MatchString(`[^0-9\-]`, data.ManualPhone)
	if manualPhoneBool {
		return !phoneBool, "manual_phone字段只能为数字和'-'"
	}
	manualPhoneResult := utils.PickTelephone(data.ManualPhone)
	if manualPhoneResult == "" || !ValidateShortNumber(data.ManualPhone) || !ValidateShortTelephone(data.ManualPhone) ||
		!Validate400Number(data.ManualPhone) || !ValidateSmsLongNumber(data.ManualPhone) {
		return false, "请检查manual_phone字段号码格式"
	}

	items, err := model.Reference{}.GetItems(utils.GetMysqlClient())
	if err != nil {
		return false, "获取数据失败"
	}
	var existDatas []map[string]interface{}
	for _, v := range items {
		localData := make(map[string]interface{})
		data := make(map[string]interface{})
		var name []string
		name = strings.Split(v.AliasNames, ",")
		name = append(name, v.Name)
		localData["name"] = name
		localData["phone"] = strings.Split(v.Phone, ",")
		localData["sender_id"] = v.SenderId
		localData["manualphone"] = v.ManualPhone
		localData["website"] = v.Website
		localData["domain"] = strings.Split(v.Phone, ",")
		data[v.Name] = localData
		existDatas = append(existDatas, data)
	}
	/*

		后面要删除

	*/
	return true, "aaaa"
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

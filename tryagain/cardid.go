package tryagain

import (
	"regexp"
)

/*
* 13开头排序：(0-9)（134 135 136 137 138 139 130 131 132 133）
* 14开头排序：(5-9)（147 148 145 146 149）
* 15开头排序：(0-3|5-9)（150 151 152 157 158 159 155 156 153）
* 16开头排序：(6-7)（166 167）
* 17开头排序：(1-8)（172 178 171 175 176 173 174 177）
* 18开头排序：(0-9)（182 183 184 187 188 185 186 180 181 189）
* 19开头排序：(1|3|5|6|8|9)（195 198 196 191 193 199）
 */
var MobilePhonePrefix = map[string]bool{
	"130": true,
	"131": true,
	"132": true,
	"133": true,
	"134": true,
	"135": true,
	"136": true,
	"137": true,
	"138": true,
	"139": true,
	"145": true,
	"146": true,
	"147": true,
	"148": true,
	"149": true,
	"150": true,
	"151": true,
	"152": true,
	"153": true,
	"155": true,
	"156": true,
	"157": true,
	"158": true,
	"159": true,
	"166": true,
	"167": true,
	"171": true,
	"172": true,
	"173": true,
	"174": true,
	"175": true,
	"176": true,
	"177": true,
	"178": true,
	"181": true,
	"182": true,
	"183": true,
	"184": true,
	"185": true,
	"186": true,
	"187": true,
	"188": true,
	"189": true,
	"191": true,
	"193": true,
	"195": true,
	"196": true,
	"198": true,
	"199": true,
}

/*
 * 13开头排序：(0-9)（134 135 136 137 138 139 130 131 132 133）
 * 14开头排序：(5-9)（147 148 145 146 149）
 * 15开头排序：(0-3|5-9)（150 151 152 157 158 159 155 156 153）
 * 16开头排序：(6-7)（166 167）
 * 17开头排序：(1-8)（172 178 171 175 176 173 174 177）
 * 18开头排序：(0-9)（182 183 184 187 188 185 186 180 181 189）
 * 19开头排序：(1|3|5|6|8|9)（195 198 196 191 193 199）
 */

func ExtractMobilePhoneDs(txt string) []string {
	var phoneIDs []string
	var PhoneFormat = "[^0-9.-_](1[3-9]\\d{9})[^0-9]"
	phoneRegx := regexp.MustCompile(PhoneFormat)
	phoneNums := phoneRegx.FindAllString(txt, -1)
	for _, v := range phoneNums {
		phoneInfo := []rune(v)
		vLen := len(phoneInfo)
		startIndex, stopIdx := 0, len(phoneInfo)
		if phoneInfo[0] < '0' || phoneInfo[0] > '9' {
			startIndex = 1
		}
		if phoneInfo[vLen-1] < '0' || phoneInfo[vLen-1] > '9' {
			stopIdx = stopIdx - 1
		}
		phoneID := string(phoneInfo[startIndex:stopIdx])
		phoneIDPrefix := phoneID[0:3]
		_, ok := MobilePhonePrefix[phoneIDPrefix]
		if ok {
			phoneIDs = append(phoneIDs, phoneID)
		}
	}
	return phoneIDs
}

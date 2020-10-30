package control

import (
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
	"strconv"
	"strings"
)

func chkCity(city string, rule *models.NewsSetting) bool {
	if city == "" {
		return false
	}
	if rule.CitySel == 0 {
		return true
	}
	//城市正选
	cityDic := map[string]bool{}
	cityInfo := strings.Replace(rule.CityList, "，", ",", -1)
	cityInfo = strings.ToLower(cityInfo)
	if rule.CitySel == 1 {
		if cityInfo == "" || cityInfo == "all" {
			return true
		}
		cityList := strings.Split(cityInfo, ",")
		for _, v := range cityList {
			cityDic[v] = true
		}
		return cityDic[city]
	}
	//针对反选的部分
	if cityInfo == "" || cityInfo == "all" {
		return false
	}
	cityList := strings.Split(cityInfo, ",")
	for _, v := range cityList {
		cityDic[v] = true
	}
	return !cityDic[city]
}

func chkHID(hid string, rule *models.NewsSetting) bool {
	if hid == "" {
		return false
	}
	if rule.HidSel == 0 {
		return true
	}
	//hid正选
	hidDic := map[string]bool{}
	hidInfo := strings.Replace(rule.HidList, "，", ",", -1)
	hidInfo = strings.ToLower(hidInfo)
	if rule.HidSel == 1 {
		if hidInfo == "" || hidInfo == "all" {
			return true
		}
		cityList := strings.Split(hidInfo, ",")
		for _, v := range cityList {
			hidDic[v] = true
		}
		return hidDic[hid]
	}
	//针对反选的部分
	if hidInfo == "" || hidInfo == "all" {
		return false
	}
	cityList := strings.Split(hidInfo, ",")
	for _, v := range cityList {
		hidDic[v] = true
	}
	return !hidDic[hid]
}

func chkChannel(channel string, rule *models.NewsSetting) bool {
	if rule.ChnSel == 0 {
		return true
	}
	chnInfo := strings.Replace(rule.ChnList, "，", ",", -1)
	chnInfo = strings.ToLower(chnInfo)
	dicDic := map[string]bool{}
	chnList := strings.Split(chnInfo, ",")
	for _, v := range chnList {
		dicDic[v] = true
	}
	if rule.ChnSel == 1 {
		if rule.ChnList == "" || rule.ChnList == "all" {
			return true
		}
		return dicDic[channel]
	}
	//针对反选的部分
	if chnInfo == "" || chnInfo == "all" {
		return false
	}
	return !dicDic[channel]
}

//compareVersion ...笔记两个版本，前大返回1，一样大返回0，后者大返回-1
func compareVersion(v1, v2 string) (int, error) {
	v1List := strings.Split(v1, ".")
	v2List := strings.Split(v2, ".")
	minLen := len(v1List)
	if len(v2List) > minLen {
		minLen = len(v2List)
	}
	for i := 0; i < minLen; i++ {
		v1Value, err := strconv.Atoi(v1List[i])
		if err != nil {
			return 0, err
		}
		v2Value, err := strconv.Atoi(v2List[i])
		if err != nil {
			return 0, err
		}
		if v1Value > v2Value {
			return 1, nil
		}
		if v1Value < v2Value {
			return -1, nil
		}
	}
	if len(v1List) == len(v2List) {
		return 0, nil
	}

	if len(v1List) > len(v2List) {
		return 1, nil
	}
	return -1, nil
}

func chkVersion(ver string, rule *models.NewsSetting) bool {
	//如果未开启版本控制，返回true
	if rule.VersionSel == 0 {
		return true
	}
	//如果版本为空，同时开启版本选择，直接false
	if ver == "" {
		return false
	}
	minVersion := rule.VersionMin
	if minVersion == "" {
		minVersion = "0.0.0.0"
	}
	maxVersion := rule.VersionMax
	if maxVersion == "" {
		maxVersion = "0.0.0.0"
	}
	result, err := compareVersion(ver, minVersion)
	if err != nil {
		return false
	}
	if result == -1 {
		return false
	}

	result, err = compareVersion(ver, maxVersion)
	if err != nil {
		return false
	}
	if result == 1 {
		return false
	}
	return true
}

func (s *settingCenter) ClientUpdate(request *utils.UpdateArgs) (string, error) {
	items, err := s.newsSetting.FindItems(s.db)
	if err != nil {
		return "", err
	}
	for _, v := range items {
		cityFlag := chkCity(request.City, v)
		if cityFlag == false {
			continue
		}
		hidFlag := chkHID(request.UID, v)
		if !hidFlag {
			continue
		}
		versionFlag := chkVersion(request.Ver, v)
		if !versionFlag {
			continue
		}
		chnFlag := chkChannel(request.Frm, v)
		if !chnFlag {
			continue
		}
		return v.Response, nil

	}
	return "{\"code\":-1}", nil
}

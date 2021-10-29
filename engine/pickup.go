package main

import (
	"github.com/sdjyliqi/known01/reader"
	"github.com/sdjyliqi/known01/utils"
	"strings"
)

func pickupFiles(fileNames string) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	files := strings.Split(fileNames, ";")
	for _, v := range files {
		filePickup, err := pickupOnFile(v)
		if err != nil {
			return nil, err
		}
		result[v] = filePickup
	}
	return result, nil
}

func pickupOnFile(f string) (map[string]int, error) {
	content, err := reader.FReader.ReadTxt(f)
	if err != nil {
		return nil, err
	}
	//识别内容中的身份证号
	cardIDs, _ := utils.ExtractMobilePhone(string(content))
	phoneIDs := utils.PickTelephone(string(content))
	result := map[string]int{
		"card":  len(cardIDs),
		"phone": len(phoneIDs),
	}
	return result, nil

}

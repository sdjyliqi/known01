package utils

import (
	"errors"
	"reflect"
	"strings"
)

func GetStructByName(elem interface{}, index string) (interface{}, error) {
	immutable := reflect.ValueOf(elem).Elem()
	for i := 0; i < immutable.NumField(); i++ {
		idx := immutable.Type().Field(i).Name
		if strings.ToLower(idx) == strings.ToLower(index) {
			//fmt.Println("==========", immutable.FieldByName(idx).Elem())
			return immutable.FieldByName(idx).Elem().String(), nil
		}
	}
	return "", errors.New("not-find")
}

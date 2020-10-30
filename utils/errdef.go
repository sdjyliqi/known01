package utils

//这里定义全部的err
import "errors"

var (
	UserCenterRegisterPhoneExist = errors.New("uc-phone-existed")
)

//定义错误描述
var Descriptions = map[error]string{
	UserCenterRegisterPhoneExist: "The phone has already existed",
}

// 定义异常对应错误码
var StatusCodes = map[error]int{
	UserCenterRegisterPhoneExist: 400001,
}

//GetErrDesc ... 获取错误描述
func GetErrDesc(err error) string {
	v, ok := Descriptions[err]
	if ok {
		return v
	}
	return "unknown"
}

//GetErrCode ...获取错误码
func GetErrCode(err error) int {
	v, ok := StatusCodes[err]
	if ok {
		return v
	}
	return 999999
}

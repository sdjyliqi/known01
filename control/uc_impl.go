package control

import (
	"errors"
	"strings"
)

func (uc *userCenter) Login(userID, passport string) error {
	isValid, err := uc.UserBasic.ChkPassportValid(uc.db, userID, passport)
	if err != nil {
		return err
	}
	if !isValid {
		return errors.New("passport-invalid")
	}
	return nil
}

func (uc *userCenter) Logout(userID string) error {
	return nil
}

func (uc *userCenter) UserChn(userID string) ([]string, error) {
	item, err := uc.UserBasic.GetUserBasic(uc.db, userID)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return []string{}, nil
	}
	return strings.Split(item.Chn, ","), nil
}

func (uc *userCenter) UserAuthChn(userID, requestChn string) string {
	item, err := uc.UserBasic.GetUserBasic(uc.db, userID)
	if err != nil {
		return "XXX"
	}
	//如果该用户的授权渠道为空，证明全部放开
	if item.Chn == "" {
		return requestChn
	}
	//如果用户请求的渠道为空，直接返回已经授权的渠道
	if requestChn == "" {
		return item.Chn
	}

	chnMap := map[string]bool{}
	chnListInDB := strings.Split(item.Chn, ",")
	for _, v := range chnListInDB {
		chnMap[v] = true
	}
	authChnList := []string{}
	chnListRequest := strings.Split(requestChn, ",")
	for _, v := range chnListRequest {
		_, ok := chnMap[v]
		if ok {
			authChnList = append(authChnList, v)
		}
	}
	if len(authChnList) == 0 {
		return "XXX"
	}
	return strings.Join(authChnList, ",")
}

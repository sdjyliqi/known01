package utils

import (
	"crypto/md5"
	"encoding/hex"
)

const salt = "feirar1234567890!#"

func CreatePassportSec(passport string) string {
	return passport
}

func CreateToken(userName, passport string) string {
	h := md5.New()
	h.Write([]byte(userName + "_" + passport))
	return hex.EncodeToString(h.Sum(nil))
}

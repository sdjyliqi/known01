package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"testing"
)

func Test_GetUserBasicByID(t *testing.T) {
	testUserBasic := UserBasic{}
	item, err := testUserBasic.GetUserBasic(testutil.TestMysql, "liqi")
	t.Log(item, err)
	isValid, err := testUserBasic.ChkPassportValid(testutil.TestMysql, "liqi", "liqi123456")
	t.Log(isValid, err)
}

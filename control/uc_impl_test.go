package control

import (
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

//
//func Test_Login(t *testing.T) {
//
//}
//
//func Test_Register(t *testing.T) {
//
//}
//
func Test_UserAuthChn(t *testing.T) {
	ucUtil := CreateUserCenter(&testutil.TestCfg)
	chn := ucUtil.UserAuthChn("admin", "all")
	t.Log(chn)
	chn = ucUtil.UserAuthChn("liqi", "all,BZ")
	t.Log(chn)

	chn = ucUtil.UserAuthChn("liqi", "")
	t.Log(chn)
}

func Test_UserChn(t *testing.T) {
	ucUtil := CreateUserCenter(&testutil.TestCfg)
	chn, err := ucUtil.UserChn("admin")
	assert.Nil(t, err)
	t.Log(chn)
	chn, err = ucUtil.UserChn("liqi")
	assert.Nil(t, err)
	t.Log(chn)

}

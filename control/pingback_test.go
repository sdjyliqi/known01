package control

//相关封装函数的UT
import (
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var TestUtil = CreatePingbackCenter(&testutil.TestCfg)

func Test_GetActiveDetailItems(t *testing.T) {
	item, count, err := TestUtil.GetActiveDetailItems("", 1, 10, 0, time.Now().Unix())
	assert.Nil(t, err)
	t.Log(item, count)
}

func Test_GetActiveDetailChannels(t *testing.T) {
	item, err := TestUtil.GetActiveChannel("admin")
	assert.Nil(t, err)
	t.Log(item)
}

func Test_GetFeirarDetailItems(t *testing.T) {
	items, count, err := TestUtil.GetFeirarDetailItems("", 1, 10, 0, time.Now().Unix())
	assert.Nil(t, err)
	t.Log(items, count)
}

func Test_GetInstallDetailItems(t *testing.T) {
	items, count, err := TestUtil.GetInstallDetailItems("", 1, 10, 0, time.Now().Unix())
	assert.Nil(t, err)
	t.Log(items, count)
}

func Test_GetInstallDetailChannels(t *testing.T) {
	item, err := TestUtil.GetInstallChannel("admin")
	assert.Nil(t, err)
	t.Log(item)
}

func Test_GetUninstallDetailItems(t *testing.T) {
	items, count, err := TestUtil.GetUninstallDetailItems("", 1, 10, 0, time.Now().Unix())
	assert.Nil(t, err)
	t.Log(items, count)
}

//func Test_GetNewsDetailItems(t *testing.T) {
//	items, count, err := TestUtil.GetNewsDetailItems("", 1, 10, 0, time.Now().Unix())
//	assert.Nil(t, err)
//	t.Log(items, count)
//}

func Test_GetPreserveDetailItems(t *testing.T) {
	items, count, err := TestUtil.GetPreserveDetailItems("", 1, 10, 0, time.Now().Unix())
	assert.Nil(t, err)
	t.Log(items, count)
}

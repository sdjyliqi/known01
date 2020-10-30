package control

import (
	"encoding/json"
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//GetHistoryCalculator ...基于历史留存数据
func Test_GetHistoryCalculator(t *testing.T) {
	util := CreatePingbackCenter(&testutil.TestCfg)
	day := time.Now().Add(time.Duration(-1*7*24) * time.Hour)
	items, err := util.GetInstallHistoryCalculator("all", day.Unix(), 9)
	assert.Nil(t, err)
	t.Log(items)
	co, err := json.Marshal(items)
	t.Log(string(co))
}

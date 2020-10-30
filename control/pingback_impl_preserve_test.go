package control

import (
	"encoding/json"
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_GetPreserveHistoryCalculator(t *testing.T) {
	util := CreatePingbackCenter(&testutil.TestCfg)
	day := time.Now().Add(time.Duration(-1*7*24) * time.Hour)
	items, err := util.GetPreserveHistoryCalculator("all", day.Unix(), 7)
	assert.Nil(t, err)
	t.Log(items)
	co, err := json.Marshal(items)
	t.Log(string(co))
}

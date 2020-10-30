package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/sdjyliqi/feirars/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_PreserveDetailGetItemsByPage(t *testing.T) {
	testPreserveDetail := PreserveDetail{}
	items, count, err := testPreserveDetail.GetItemsByPage(testutil.TestMysql, "all,BZ", 1, 10, 0, time.Now().Unix())
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.LastUpdate.Format(utils.FullTime), v.Uv)
	}
}
func Test_PreserveDetailGetAllChannels(t *testing.T) {
	testActiveDetail := PreserveDetail{}
	items, err := testActiveDetail.GetAllChannels(testutil.TestMysql)
	assert.Nil(t, err)
	t.Log(items, err)
}

func Test_PreserveDetailGetChartItems(t *testing.T) {
	testActiveDetail := PreserveDetail{}
	items, err := testActiveDetail.GetChartItems(testutil.TestMysql, "", 0, time.Now().Unix())
	assert.Nil(t, err)
	t.Log(items, err)
}

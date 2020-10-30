package models

import (
	"encoding/json"
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_FeirarDetailGetItemsByPage(t *testing.T) {
	testFeirarDetail := FeirarDetail{}
	items, count, err := testFeirarDetail.GetItemsByPage(testutil.TestMysql, "", "", 1, 100, 0, time.Now().Unix())
	t.Log(items, count, err)

	items, count, err = testFeirarDetail.GetItemsByPage(testutil.TestMysql, "/api/update", "", 1, 100, 0, time.Now().Unix())
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.EventDay, v.EventKey, v.Pv)
	}
	//items, count, err = testFeirarDetail.GetItemsByPage(testutil.TestMysql, "/api/FeiRarNews", "", 1, 100, 0, time.Now().Unix())
	//t.Log(items, count, err)

}

func Test_FeirarDetailGetChartItems(t *testing.T) {
	testFeirarDetail := FeirarDetail{}
	item, err := testFeirarDetail.GetChartItems(testutil.TestMysql, "", "all,BZ", 0, time.Now().Unix())
	strItems, err := json.Marshal(item)
	t.Log(string(strItems), err)

	item, err = testFeirarDetail.GetChartItems(testutil.TestMysql, "/api/update", "all,BZ", 0, time.Now().Unix())
	assert.Nil(t, err)
	strItems, err = json.Marshal(item)
	t.Log(string(strItems), err)

	item, err = testFeirarDetail.GetChartItems(testutil.TestMysql, "/api/FeiRarNews", "all,BZ", 0, time.Now().Unix())
	assert.Nil(t, err)
	strItems, err = json.Marshal(item)
	t.Log(string(strItems), err)
}

func Test_FeirarDetailGetAllChannels(t *testing.T) {
	testActiveDetail := FeirarDetail{}
	items, err := testActiveDetail.GetAllChannels(testutil.TestMysql, "")
	assert.Nil(t, err)
	t.Log(items, err)

	items, err = testActiveDetail.GetAllChannels(testutil.TestMysql, "/api/update")
	assert.Nil(t, err)
	t.Log(items, err)

	items, err = testActiveDetail.GetAllChannels(testutil.TestMysql, "/api/FeiRarNews")
	assert.Nil(t, err)
	t.Log(items, err)
}

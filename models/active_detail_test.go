package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_ActiveDetailGetItemsByPage(t *testing.T) {
	testActiveDetail := ActiveDetail{}
	items, count, err := testActiveDetail.GetItemsByPage(testutil.TestMysql, "all,BZ", 1, 20, 0, time.Now().Unix())
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.EventDay, v.Pv)
	}
}

func Test_ActiveDetailGetAllChannels(t *testing.T) {
	testActiveDetail := ActiveDetail{}
	items, err := testActiveDetail.GetAllChannels(testutil.TestMysql)
	assert.Nil(t, err)
	t.Log(items, err)
}

func Test_ActiveDetailGetChartItems(t *testing.T) {
	testActiveDetail := ActiveDetail{}
	items, err := testActiveDetail.GetChartItems(testutil.TestMysql, "", 0, time.Now().Unix())
	assert.Nil(t, err)
	t.Log(items, err)
}

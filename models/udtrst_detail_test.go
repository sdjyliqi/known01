package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_UdtrstDetailGetItemsByPage(t *testing.T) {
	testMod := UdtrstDetail{}
	items, count, err := testMod.GetItemsByPage(testutil.TestMysql, "", 1, 10, 0, time.Now().Unix())
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.EventDay, v.Channel)
	}
}

func Test_UdtrstDetailGetAllChannels(t *testing.T) {
	testMod := UdtrstDetail{}
	items, err := testMod.GetAllChannels(testutil.TestMysql)
	assert.Nil(t, err)
	t.Log(items, err)
}

func Test_UdtrstDetailGetChartItems(t *testing.T) {
	testMod := UdtrstDetail{}
	items, err := testMod.GetChartItems(testutil.TestMysql, "", 0, time.Now().Unix())
	assert.Nil(t, err)
	t.Log(items, err)
}

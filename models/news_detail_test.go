package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_NewsDetailGetItemsByPage(t *testing.T) {
	testNewsDetail := NewsDetail{}
	items, count, err := testNewsDetail.GetItemsByPage(testutil.TestMysql, "", 1, 10, 0, time.Now().Unix(),"newsshow")
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.EventDay, v.Pv)
	}
}

func Test_NewsDetailGetAllChannels(t *testing.T) {
	testNewsDetail := NewsDetail{}
	items, err := testNewsDetail.GetAllChannels(testutil.TestMysql,"newsshow")
	assert.Nil(t, err)
	t.Log(items, err)
}

func Test_NewsDetailGetChartItems(t *testing.T) {
	testNewsDetail := NewsDetail{}
	items, err := testNewsDetail.GetChartItems(testutil.TestMysql, "", 0, time.Now().Unix(),"newsshow")
	assert.Nil(t, err)
	t.Log(items, err)

	items, err = testNewsDetail.GetChartItems(testutil.TestMysql, "", 0, time.Now().Unix(),"traygametipsshow")
	assert.Nil(t, err)
	t.Log(items, err)
}

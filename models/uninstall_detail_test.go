package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_UninstallDetailGetItemsByPage(t *testing.T) {
	testUninstallDetail := UninstallDetail{}
	items, count, err := testUninstallDetail.GetItemsByPage(testutil.TestMysql, "BZ", 1, 10, 0, time.Now().Unix())
	assert.Nil(t, err)
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.Channel)
	}
}

func Test_UninstallDetailGetAllChannels(t *testing.T) {
	testUninstallDetail := UninstallDetail{}
	items, err := testUninstallDetail.GetAllChannels(testutil.TestMysql)
	assert.Nil(t, err)
	t.Log(items, err)
}

func Test_UninstallDetailGetChartItems(t *testing.T) {
	testUninstallDetail := UninstallDetail{}
	items, err := testUninstallDetail.GetChartItems(testutil.TestMysql, "all", 0, time.Now().Unix())
	assert.Nil(t, err)
	t.Log(items, err)
}

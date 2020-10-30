package models

import (
	"fmt"
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_InstallDetailGetItemsByPage(t *testing.T) {
	testInstallDetail := InstallDetail{}
	items, count, err := testInstallDetail.GetItemsByPage(testutil.TestMysql, "all,BZ", 1, 10, 0, time.Now().Unix())
	t.Log(items, count, err)
	for _, v := range items {
		t.Log(v.EventDay, v.Pv, v.Day1Active, v.Day2Active, v.Day3Active, v.WeekActive)
	}
}

func Test_InstallDetailGetAllChannels(t *testing.T) {
	testActiveDetail := InstallDetail{}
	items, err := testActiveDetail.GetAllChannels(testutil.TestMysql)
	assert.Nil(t, err)
	t.Log(items, err)
}

func Test_InstallDetailGetChartItems(t *testing.T) {
	testActiveDetail := InstallDetail{}
	items, err := testActiveDetail.GetChartItems(testutil.TestMysql, "", 0, time.Now().Unix())
	assert.Nil(t, err)
	t.Log(items, err)
}

func Test_GetItemsForHistory(t *testing.T) {
	day := time.Now().Add(time.Duration(-1*7*24) * time.Hour)
	fmt.Println(day)
	testActiveDetail := InstallDetail{}
	items, err := testActiveDetail.GetItemsForHistory(testutil.TestMysql, "all", day.Unix(), 3)
	assert.Nil(t, err)
	t.Log(items, err)
}

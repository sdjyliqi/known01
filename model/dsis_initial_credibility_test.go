package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ScoreGetItems(t *testing.T) {
	items, err := InitialCredibilityModel.GetItems()
	assert.Nil(t, err)
	t.Logf("数据列表:%+v", items)
}

func Test_SGetItemByIdx(t *testing.T) {
	item, err := InitialCredibilityModel.GetItemByIdx("D1M1P1")
	assert.Nil(t, err)
	t.Log(item)
}

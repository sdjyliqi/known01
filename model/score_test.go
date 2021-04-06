package model

import (
	"github.com/sdjyliqi/known01/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ScoreGetItems(t *testing.T) {
	items, err := Score{}.GetItems(testutils.DBEngineTest)
	assert.Nil(t, err)
	t.Logf("数据列表:%+v", items)
}

func Test_SGetItemByIdx(t *testing.T) {
	item, err := Score{}.GetItemByIdx("D1M1P1", testutils.DBEngineTest)
	assert.Nil(t, err)
	t.Log(item)
}

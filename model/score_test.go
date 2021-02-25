package model

import (
	"github.com/stretchr/testify/assert"
	"known01/utils"
	"testing"
)

func Test_ScoreGetItems(t *testing.T) {
	items, err := Score{}.GetItems(utils.GetMysqlClient())
	assert.Nil(t, err)
	t.Logf("数据列表:%+v", items)
}

func Test_SGetItemByIdx(t *testing.T) {
	item, err := Score{}.GetItemByIdx("D1M1P1", utils.GetMysqlClient())
	assert.Nil(t, err)
	t.Log(item)
}

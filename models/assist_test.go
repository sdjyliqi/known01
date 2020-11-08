package models

import (
	"github.com/stretchr/testify/assert"
	"known01/utils"
	"testing"
)

func Test_AssistGetItems(t *testing.T) {
	items, err := Assist{}.GetItems(utils.GetMysqlClient())
	assert.Nil(t, err)
	t.Logf("副助词列表如:%+v", items)
}

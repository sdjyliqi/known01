package model

import (
	"github.com/stretchr/testify/assert"
	"known01/utils"
	"testing"
)

func Test_NewsGetItems(t *testing.T) {
	items, err := News{}.GetItems(utils.GetMysqlClient(), 0)
	assert.Nil(t, err)
	t.Logf("信息流列表如:%+v", items)
}

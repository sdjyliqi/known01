package models

import (
	"github.com/stretchr/testify/assert"
	"known01/utils"
	"testing"
)

func Test_TemplatesGetItems(t *testing.T) {
	items, err := Templates{}.GetItems(utils.GetMysqlClient())
	assert.Nil(t, err)
	t.Logf("模板列表如:%+v", items)
}

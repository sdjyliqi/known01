package model

import (
	"github.com/sdjyliqi/known01/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_TemplatesGetItems(t *testing.T) {
	items, err := MessageTMPModel.GetItems(testutils.DBEngineTest)
	assert.Nil(t, err)
	t.Logf("模板列表如:%+v", items)

	for _, v := range items {
		t.Log(v)
	}
}

package model

import (
	"github.com/sdjyliqi/known01/testutils"
	"github.com/stretchr/testify/assert"

	"testing"
)

func Test_AssistGetItems(t *testing.T) {
	items, err := Assist{}.GetItems(testutils.DBEngineTest)
	assert.Nil(t, err)
	t.Logf("副助词列表如:%+v", items)
}

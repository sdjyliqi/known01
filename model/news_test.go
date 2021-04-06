package model

import (
	"github.com/sdjyliqi/known01/testutils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewsGetItems(t *testing.T) {
	items, err := News{}.GetItems(testutils.DBEngineTest, 0)
	assert.Nil(t, err)
	t.Logf("信息流列表如:%+v", items)
}

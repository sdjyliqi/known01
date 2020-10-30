package models

import (
	"github.com/sdjyliqi/feirars/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FindItems(t *testing.T) {
	newsSettingUtil := NewsSetting{}
	items, err := newsSettingUtil.FindItems(testutil.TestMysql)
	assert.Nil(t, err)
	t.Log(items, err)
}

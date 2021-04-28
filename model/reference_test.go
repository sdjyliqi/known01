package model

import (
	"github.com/stretchr/testify/assert"
	"known01/utils"
	"testing"
	"time"
)

func Test_ReferenceGetItemsAndGetItemByID(t *testing.T) {
	items, err := Reference{}.GetPages(0, 10)
	assert.Nil(t, err)
	t.Logf("Items:%+v", items)

	if items != nil {
		chkItem := items[0]
		item, err := Reference{}.GetItemByID(chkItem.Id)
		assert.Nil(t, err)
		assert.Equal(t, chkItem.Name, item.Name)
		assert.Equal(t, chkItem.Domain, item.Domain)
	}
}

func Test_UpdateInsertItem(t *testing.T) {
	item := Reference{
		Name:         "test1922",
		CategoryId:   0,
		AliasNames:   "test",
		Phone:        "test",
		ManualPhone:  "test",
		Website:      "",
		Domain:       "",
		LastModified: time.Now(),
	}
	err := Reference{}.InsertItemByID(item)
	assert.Nil(t, err)
	count, err := utils.GetMysqlClient().Where("name=?", "test").Delete(Reference{})
	assert.Equal(t, int64(1), count)

}

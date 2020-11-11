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

	for _, v := range items {
		t.Log(v)
		simVal := utils.SimHashTool.Hash(v.Detail)
		v.SimHash = simVal
		t.Log(simVal)
		count, err := utils.GetMysqlClient().Table(v.TableName()).ID(v.Id).Cols("sim_hash").Update(v)
		t.Log(count, err)
	}
}

func Test_TemplatesUpdateSimHash(t *testing.T) {
	items, err := Templates{}.GetItems(utils.GetMysqlClient())
	assert.Nil(t, err)
	for _, v := range items {
		//simVal:= utils.SimHashTool.Hash(v.Detail)
		//v.SimHash = simVal

		t.Log(v)
		//c,err:=utils.GetMysqlClient().Table(v.TableName()).ID(v.Id).Cols("simhash").Update(v)
		//t.Log(c,err)
	}
}

package model

import (
	"github.com/sdjyliqi/known01/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_ResultsStore(t *testing.T) {
	err := ResultTool.Store(&Results{
		CategoryId:   utils.EngineBank,
		Detail:       "xxxxxxxxxxxxxx",
		Extract:      "sss",
		Compare:      "sssssssssssssss",
		Flag:         1,
		Suggest:      "aaaaaaaaaaaaaaa",
		LastModified: time.Now(),
	})
	assert.Nil(t, err)

}

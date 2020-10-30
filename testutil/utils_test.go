package testutil

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

//Test_InitMysql ...测试mysql初始化功能
func Test_InitMysql(t *testing.T) {
	tables, err := TestMysql.Exec("show tables")
	assert.Nil(t, err)
	t.Log(tables)
}

//Test_InitRedis  ...测试redis初始化功能
func Test_InitRedis(t *testing.T) {
	key, value := "test-key", fmt.Sprintf("%d", time.Now().Unix())
	result := TestRedis.Set(key, value, 0)
	assert.Nil(t, result.Err())
	valDB, err := TestRedis.Get(key).Result()
	assert.Nil(t, err)
	assert.Equal(t, value, valDB)
}

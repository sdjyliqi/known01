package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Student struct {
	ID   interface{}
	Name string
}

func Test_SetValue(t *testing.T) {
	s := Student{
		ID:   "999",
		Name: "liqi",
	}

	v, err := GetStructByName(&s, "ID")
	fmt.Println(v, err)
	assert.Nil(t, err)
	assert.Equal(t, "999", v)

}

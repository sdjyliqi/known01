package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReadExcel(t *testing.T) {
	var reader FileReader
	//content, err := reader.ReadExcel("")
	//assert.Nil(t, err)
	//t.Log(content)
	//
	b, err := reader.ReadTxt("D:\\gowork\\src\\known01\\data\\city.txt")
	t.Log(string(b), err)
}

func Test_ReadPDF(t *testing.T) {
	var reader FileReader
	content, err := reader.ReadDoc("")
	assert.Nil(t, err)
	t.Log(content)
}

func Test_ReadPPT(t *testing.T) {
	//D:\gowork\src\known01\data\city.txt

}

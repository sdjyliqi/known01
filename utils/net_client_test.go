package utils

import (
	"fmt"
	"io/ioutil"
	"testing"
)

//todo 需要实现该功能 curl "http://t.cn/E9t6FgQ" -i
func Test_GetHTTPClient(t *testing.T) {
	url := "http://t.cn/E9t6FgQ"
	resp, err := GetHTTPClient().Get(url)
	t.Log(resp.Header, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.Status)
	t.Log(string(body))
}

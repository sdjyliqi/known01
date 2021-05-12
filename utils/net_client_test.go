package utils

import (
	"fmt"
	"testing"
)

//todo 需要实现该功能 curl "http://t.cn/E9t6FgQ" -i
func Test_GetHTTPClient(t *testing.T) {
	url := "http://t.cn/E9t6FgQ"
	resp, err := GetHTTPClient().Get(url)
	t.Log(resp.Header, err)
	fmt.Println(resp.Status)

	fmt.Println("------------------------------------")
	/*
		client := &http.Client{}
		reqest, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return fmt.Errorf("first response")
		}
		response, _ := client.Do(reqest)
		fmt.Println(response.StatusCode)
		//t.Log(response.Header)

	*/

}

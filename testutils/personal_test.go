package testutils

import (
	"fmt"
	"github.com/prometheus/common/log"
	"net/http"
	"regexp"
	"testing"
)

func Test_HTTP302(t *testing.T) {
	client := &http.Client{}
	url := "http://www.qq.com"
	url = "http://t.cn/E9t6FgQ"
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Get url %s failed,err:%+v", url, err)
		return
	}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return fmt.Errorf("first response")
	}
	response, _ := client.Do(reqest)
	t.Log(response.StatusCode, response.Header)
	if response != nil && response.Header != nil {
		t.Log(response.Header.Get("location"))
	}
}

func Test_CardID(t *testing.T) {
	text := "我的身份证是152102198510281834"
	//pattern := "^([1-9]\\d{5}[12]\\d{3}(0[1-9]|1[012])(0[1-9]|[12][0-9]|3[01])\\d{3}[0-9xX])"

	pattern := `[0-9]{6,18}`

	reg := regexp.MustCompile(pattern)
	cardIDList := reg.FindAllString(text, -1)
	fmt.Println(cardIDList)
}

package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sdjyliqi/known01/utils"
	"io/ioutil"
	"log"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		postBody, _ := c.GetRawData()                                // 获取POST BODY体后，默认会自动清理body体中的内容
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(postBody)) //  重新给POST BODY赋值
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		start := time.Now()
		c.Next()
		responseBody := bodyLogWriter.body.String()
		//日志格式
		accessLogMap := make(map[string]interface{})
		accessLogMap["start"] = start.Format(utils.FullTime)
		accessLogMap["method"] = c.Request.Method
		accessLogMap["uri"] = c.Request.RequestURI
		accessLogMap["proto"] = c.Request.Proto
		accessLogMap["ua"] = c.Request.UserAgent()
		accessLogMap["referer"] = c.Request.Referer()
		accessLogMap["post_body"] = string(postBody)
		accessLogMap["client_ip"] = c.ClientIP()
		accessLogMap["response"] = responseBody
		accessLogMap["latency"] = time.Now().Sub(start).Microseconds()
		logContent, _ := json.Marshal(accessLogMap)
		log.Println(string(logContent))
	}
}

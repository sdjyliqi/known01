package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/prometheus/client_golang/prometheus"
	uuid "github.com/satori/go.uuid"
	"known01/handle"
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
		headIdx := "X-Request-ID"
		c.Set(headIdx, uuid.NewV4().String())
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		start := time.Now()
		c.Next()
		responseBody := bodyLogWriter.body.String()
		//日志格式
		uuid, _ := c.Get(headIdx)
		glog.Infoln("INFO", uuid, c.ClientIP(), start, c.Request.Method, c.Request.RequestURI, c.Request.UserAgent(), len(responseBody), ",cost:", time.Now().Sub(start).Microseconds(), "response:", c.Writer.Status())
		fmt.Println("INFO", uuid, c.ClientIP(), start, c.Request.Method, c.Request.RequestURI, c.Request.UserAgent(), len(responseBody), ",cost:", time.Now().Sub(start).Microseconds(), "response:", c.Writer.Status())
		// prometheus监控增加标签，包括服务名、请求地址和响应时间
		{
			prometheusLables := prometheus.Labels{"application": "data-audit", "uri": c.Request.RequestURI, "status": "200"}
			handle.WebRequestSecondsBucket.With(prometheusLables).Observe(time.Now().Sub(start).Seconds())
		}
	}
}

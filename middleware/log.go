package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
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
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		start := time.Now()
		c.Next()
		responseBody := bodyLogWriter.body.String()
		//日志格式
		glog.Infoln(c.ClientIP(), start, c.Request.Method, c.Request.RequestURI, c.Request.UserAgent(), len(responseBody), ",cost:", time.Now().Sub(start).Microseconds())
		fmt.Println(c.ClientIP(), start, c.Request.Method, c.Request.RequestURI, c.Request.UserAgent(), len(responseBody), ",cost:", time.Now().Sub(start).Microseconds())
	}
}

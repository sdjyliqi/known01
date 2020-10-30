package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		path := c.Request.URL.Path
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		requestID := c.GetHeader("X-Request-Id")
		glog.Infof("| %3d | %13v | %15s | %s  %s | %5d| %s |",
			statusCode,
			latency,
			clientIP,
			method,
			path,
			c.Writer.Size(),
			requestID,
		)
	}
}

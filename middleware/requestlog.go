package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"

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
		fmt.Printf("| Header:%+v | %13v | %15s | %s  %s | %5d \n",
			c.Request.Header,
			latency,
			clientIP,
			method,
			path,
			c.Writer.Size(),
		)
	}
}

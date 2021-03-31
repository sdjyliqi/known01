package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"known01/utils"

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
		fmt.Printf("Time at:%s %+v |cost:%v |%s %s |content-length:%d\n",
			start.Format(utils.FullTime),
			clientIP,
			latency,
			method,
			path,
			c.Writer.Size(),
		)
	}
}

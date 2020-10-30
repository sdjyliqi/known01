package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
)

//... RequestIDMiddleware  add request id into header
func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuidReq := uuid.NewV4()
		c.Writer.Header().Set("X-Request-Id", uuidReq.String())
		c.Request.Header.Add("X-Request-Id", uuidReq.String())
		c.Next()
	}
}

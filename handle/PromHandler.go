package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PromHandler(handler http.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler.ServeHTTP(c.Writer, c.Request)
	}
}

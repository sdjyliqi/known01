package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//UCLogin ...用户登录
func UCLogin(c *gin.Context) {
	token := "000011111122222"
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": token})
}

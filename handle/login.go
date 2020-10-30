package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/sdjyliqi/known01/utils"
	"net/http"
)

//LoginArgs ...登录的请求体
type LoginArgs struct {
	UserName string `json:"name" form:"name" binding:"required"`
	Passport string `json:"passport" form:"passport" binding:"required"`
}

func UCLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "token": "1234566"})
}

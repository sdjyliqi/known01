package router

import (
	"github.com/gin-gonic/gin"
	"known01/handle"
	"known01/middleware"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())
	r.POST("/brain", handle.JudgeMessage) //获取信息详情
}

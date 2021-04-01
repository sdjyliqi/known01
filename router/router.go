package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sdjyliqi/known01/handle"
	"github.com/sdjyliqi/known01/middleware"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())
	r.POST("/brain", handle.JudgeMessage)   //获取信息详情
	r.GET("/brain", handle.JudgeMessageGET) //获取信息详情
}

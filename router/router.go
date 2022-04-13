package router

import (
	"github.com/gin-gonic/gin"
	"known01/handle"
	"known01/middleware"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())
	r.POST("/message/brain", handle.JudgeMessage)   //识别诈骗短消息
	r.GET("/message/brain", handle.JudgeMessageGET) //识别诈骗短消息
}

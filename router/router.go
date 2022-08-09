package router

import (
	"github.com/gin-gonic/gin"
	"known01/handle"
	"known01/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())
	r.GET("/message/ping", handle.Ping)           //探活接口
	r.POST("/message/brain", handle.JudgeMessage) //识别诈骗短消息

	r.GET("/metrics", handle.PromHandler(promhttp.Handler())) //prometheus监控Url
}

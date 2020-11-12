package router

import (
	"github.com/gin-gonic/gin"
	"known01/handle"
	"known01/middleware"
)

func InitRouter(r *gin.Engine) {
	//r.Use(middleware.Cors())
	r.Use(middleware.Logger())

	GroupV2 := r.Group("/v1")
	{
		GroupV2.GET("/news", handle.GetNews)        //获取列表
		GroupV2.POST("/news", handle.GetNews)       //获取列表
		GroupV2.GET("/brain", handle.JudgeMessage)  //获取信息详情
		GroupV2.POST("/brain", handle.JudgeMessage) //获取信息详情
	}

}

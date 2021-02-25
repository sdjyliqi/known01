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

	//管理后台相关的接口
	GroupV1 := r.Group("/admin")
	{
		GroupV1.GET("/login", handle.UCLogin) //用户登录
		GroupV1.POST("/login", handle.UCLogin)
		GroupV1.POST("/logout", handle.GetNews) //用户注销
	}

}

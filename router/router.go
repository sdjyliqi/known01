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
		GroupV1.POST("/login", handle.UCLogin)                       //用户登录
		GroupV1.GET("/users/list", handle.UCUsers)                   //获取全部用户列表
		GroupV1.POST("/users/status", handle.UCUsersStatus)          //用户禁用激活状态设置
		GroupV1.POST("/users/add", handle.UCAddUsers)                //添加用户
		GroupV1.POST("/users/resetpas", handle.UCResetPassword)      //重置用户密码
		GroupV1.POST("/users/information", handle.UCShowInformation) //展示用户详细信息
		GroupV1.POST("/users/chapas", handle.UCChangePassword)       //用户修改密码
	}

}

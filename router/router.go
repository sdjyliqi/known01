package router

import (
	"github.com/gin-gonic/gin"
	"known01/handle"
	"known01/middleware"
)

func InitRouter(r *gin.Engine) {
	//middleware.Cors   ... 跨域设置
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())

	r.POST("/login", handle.UCLogin)       //用户登录
	r.POST("/login/roles", handle.UCRoles) //用户角色判断

	//管理后台相关的接口
	GroupV1 := r.Group("/admin")
	{
		GroupV1.GET("/users/list", handle.UCUsers)                  //获取全部用户列表
		GroupV1.GET("/users/status", handle.UCUsersStatus)          //用户禁用激活状态设置
		GroupV1.POST("/users/add", handle.UCAddUsers)               //添加用户
		GroupV1.GET("/users/resetpas", handle.UCResetPassword)      //管理员重置用户密码
		GroupV1.GET("/users/information", handle.UCShowInformation) //展示用户详细信息
		GroupV1.POST("/update", handle.UCAdminUpdateItem)           //管理员更新用户详细信息
	}
	//普通用户相关接口
	GroupV2 := r.Group("/editor")
	{
		GroupV2.POST("/changepas", handle.UCChangePassword)       //用户修改密码
		GroupV2.POST("/users/update", handle.UCUpdateInformation) //用户更新个人信息
	}
	//短信鉴别相关接口
	GroupV3 := r.Group("/v1")
	{
		GroupV3.GET("/news", handle.GetNews)        //获取列表
		GroupV3.POST("/news", handle.GetNews)       //获取列表
		GroupV3.GET("/brain", handle.JudgeMessage)  //获取信息详情
		GroupV3.POST("/brain", handle.JudgeMessage) //获取信息详情
	}
	//管理后台添加参考数据接口
	GroupV4 := r.Group("/reference")
	{
		GroupV4.GET("/lists", handle.GetRefList)          //获取列表
		GroupV4.GET("/getitem", handle.GetRefItem)        //获取单条详细信息
		GroupV4.POST("/updateitem", handle.UpdateRefItem) //修改数据
		GroupV4.POST("/additem", handle.AddRefItem)       //插入数据
	}
}

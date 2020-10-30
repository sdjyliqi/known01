package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sdjyliqi/known01/handle"
	"github.com/sdjyliqi/known01/middleware"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())
	//r.Use(middleware.RequestAddIPLoc())
	// uc先关接口
	GroupV1 := r.Group("/admin")
	{
		GroupV1.GET("/pingback", handle.UCLogin)
	}

	GroupV2 := r.Group("/api")
	{
		GroupV2.GET("/update", handle.UCLogin)
		GroupV2.GET("/city", handle.UCLogin)
	}

}

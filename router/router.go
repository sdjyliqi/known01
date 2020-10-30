package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sdjyliqi/feirars/handle"
	"github.com/sdjyliqi/feirars/middleware"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())
	//r.Use(middleware.RequestAddIPLoc())
	// uc先关接口
	GroupV1 := r.Group("/admin")
	{
		GroupV1.GET("/pingback", handle.HandlePingbak)
		GroupV1.GET("/login", handle.UCLogin)
		GroupV1.POST("/login", handle.UCLogin)
		GroupV1.GET("/chart", handle.HandleChart)
		GroupV1.GET("/chn", handle.HandleChannels)
		GroupV1.GET("/export", handle.Export)
		GroupV1.GET("/history", handle.HistoryCalculator)
	}

	GroupV2 := r.Group("/api")
	{
		GroupV2.GET("/update", handle.HandleUpdate)
		GroupV2.GET("/city", handle.HandleCity)
	}

}

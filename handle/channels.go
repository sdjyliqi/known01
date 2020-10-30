package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

//ChannelsArgs ...ChannelsArgs统计体
type ChannelsArgs struct {
	ModuleName string `json:"type" form:"type" binding:"required"`
	UserName   string `json:"type" form:"name" binding:"required"`
}

func HandleChannels(c *gin.Context) {
	header := c.Request.Header
	glog.Info(header)
	var reqArgs ChannelsArgs
	err := c.ShouldBind(&reqArgs)
	if err != nil || reqArgs.ModuleName == "" || reqArgs.UserName == "" {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	switch reqArgs.ModuleName {
	case "install":
		items, err := PingbackCenter.GetInstallChannel(reqArgs.UserName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return
	case "uninstall":
		items, err := PingbackCenter.GetUninstallChannel(reqArgs.UserName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return

	case "active":
		items, err := PingbackCenter.GetActiveChannel(reqArgs.UserName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return
	case "news":
		items, err := PingbackCenter.GetNewsChannel(reqArgs.UserName, "newsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return

	case "preserve":
		items, err := PingbackCenter.GetPreserveChannel(reqArgs.UserName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return
	case "feirar":
		items, err := PingbackCenter.GetFeirarChannel(reqArgs.UserName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return

	case "feirar-news":
		items, err := PingbackCenter.GetFeirarNewsChannel(reqArgs.UserName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return

	case "feirar-update":
		items, err := PingbackCenter.GetFeirarUpdateChannel(reqArgs.UserName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return

		//s升级
	case "feirar-udtrst":
		items, err := PingbackCenter.GetUdtrstChannel(reqArgs.UserName)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return

	case "feirar-righttipsshow":
		items, err := PingbackCenter.GetNewsChannel(reqArgs.UserName, "righttipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return

	case "feirar-rightnewstipsshow":
		items, err := PingbackCenter.GetNewsChannel(reqArgs.UserName, "rightnewstipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return

	case "feirar-taskbartipsshow":
		items, err := PingbackCenter.GetNewsChannel(reqArgs.UserName, "taskbartipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})

		return

	case "feirar-msgcentershow":
		items, err := PingbackCenter.GetNewsChannel(reqArgs.UserName, "msgcentershow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})

		return

	case "feirar-topcentertipsshow":
		items, err := PingbackCenter.GetNewsChannel(reqArgs.UserName, "topcentertipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})

		return

	case "feirar-traygametipsshow":
		items, err := PingbackCenter.GetNewsChannel(reqArgs.UserName, "traygametipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "items": items})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "type参数错误"})
}

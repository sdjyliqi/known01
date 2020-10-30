package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//ChartArgs ... 曲线区统计体
type ChartArgs struct {
	ModuleName string `json:"type" form:"type" binding:"required"`
	TimeStart  int64  `json:"ts" form:"ts" `
	TimeEnd    int64  `json:"te" form:"te"`
	Channels   string `json:"chn" form:"chn"`
}

func HandleChart(c *gin.Context) {
	var reqArgs ChartArgs
	err := c.ShouldBind(&reqArgs)
	if err != nil || reqArgs.ModuleName == "" || reqArgs.TimeStart < 0 || reqArgs.TimeEnd < 0 {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误1111"})
		return
	}
	switch reqArgs.ModuleName {
	case "install":
		items, err := PingbackCenter.GetInstallChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return

	case "uninstall":
		items, err := PingbackCenter.GetUninstallChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return

	case "active":
		items, err := PingbackCenter.GetActiveChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return
	case "news":
		items, err := PingbackCenter.GetNewsChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd,"newsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return

	case "preserve":
		items, err := PingbackCenter.GetPreserveChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return

	case "feirar":
		items, err := PingbackCenter.GetFeirarChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return
	case "feirar-news":
		items, err := PingbackCenter.GetFeirarNewsChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return

	case "feirar-update":
		items, err := PingbackCenter.GetFeirarUpdateChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return

	case "feirar-udtrst":
		items, err := PingbackCenter.GetUdtrstChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return


	case "feirar-righttipsshow":
		items, err := PingbackCenter.GetNewsChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd,"righttipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return
	case "feirar-rightnewstipsshow":
		items, err := PingbackCenter.GetNewsChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd,"rightnewstipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return

	case "feirar-taskbartipsshow":
		items, err := PingbackCenter.GetNewsChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd,"taskbartipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return
	case "feirar-msgcentershow":
		items, err := PingbackCenter.GetNewsChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd,"msgcentershow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return
	case "feirar-topcentertipsshow":
		items, err := PingbackCenter.GetNewsChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd,"topcentertipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return
	case "feirar-traygametipsshow":
		items, err := PingbackCenter.GetNewsChart(reqArgs.Channels, reqArgs.TimeStart, reqArgs.TimeEnd,"traygametipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": items})
		return
	}


	c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "type参数错误"})
}

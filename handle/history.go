package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"net/http"
)

type HistoryArgs struct {
	ModuleName string `json:"type" form:"type" binding:"required"`
	TimeStart  int64  `json:"ts"   form:"ts" `
	Days       int    `json:"days" form:"days"`
	Channel    string `json:"chn"  form:"chn"`
	Name       string `json:"name" form:"name"`
}

func HistoryCalculator(c *gin.Context) {
	header := c.Request.Header
	glog.Info(header)
	var reqArgs HistoryArgs
	err := c.ShouldBind(&reqArgs)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	switch reqArgs.ModuleName {
	case "install":
		cols := utils.HistoryCalculatorCols
		items, err := PingbackCenter.GetInstallHistoryCalculator(reqArgs.Channel, reqArgs.TimeStart, reqArgs.Days)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items})
		return

	case "preserve":
		cols := utils.HistoryCalculatorCols
		items, err := PingbackCenter.GetPreserveHistoryCalculator(reqArgs.Channel, reqArgs.TimeStart, reqArgs.Days)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "cols": cols, "items": items})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误"})
}

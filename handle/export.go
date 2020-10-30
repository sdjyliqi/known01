package handle

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/models"
	"github.com/sdjyliqi/feirars/utils"
	"net/http"
	"time"
)

func ExportInstall(c *gin.Context, cols []map[string]string, items []models.InstallDetailWeb) {
	excelTitleLine := utils.CreateExcelTitle(cols)
	var excelItems [][]string
	for _, v := range items {
		oneLine := []string{v.EventDay, v.Channel, v.Pv, v.Uv, v.Day1Active, v.Day2Active, v.Day3Active, v.Day4Active, v.Day5Active, v.Day6Active, v.WeekActive}
		excelItems = append(excelItems, oneLine)
	}
	filePath, err := utils.CreateExcelFile(excelTitleLine, excelItems)
	if err != nil {
		c.Error(errors.New("暂时无法导出excel，请稍后重新"))
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filePath)
	c.Header("Content-Transfer-Encoding", "binary")
	time.Sleep(1 * time.Second)
	c.File(filePath)
}

func ExportUninstall(c *gin.Context, cols []map[string]string, items []models.UninstallDetailWeb) {
	excelTitleLine := utils.CreateExcelTitle(cols)
	var excelItems [][]string
	for _, v := range items {
		oneLine := []string{v.EventDay, v.Channel, v.Pv, v.Uv, v.LastUpdate}
		excelItems = append(excelItems, oneLine)
	}
	filePath, err := utils.CreateExcelFile(excelTitleLine, excelItems)
	if err != nil {
		c.Error(errors.New("暂时无法导出excel，请稍后重新"))
		return
	}
	time.Sleep(1 * time.Second)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filePath)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
}

func ExportNewsDetail(c *gin.Context, cols []map[string]string, items []models.NewsDetailWeb) {
	excelTitleLine := utils.CreateExcelTitle(cols)
	var excelItems [][]string
	for _, v := range items {
		oneLine := []string{v.EventDay, v.Channel, v.ShowPv, v.ShowUv, v.ClickPv, v.ClickUv, v.LastUpdate}
		excelItems = append(excelItems, oneLine)
	}
	filePath, err := utils.CreateExcelFile(excelTitleLine, excelItems)
	if err != nil {
		c.Error(errors.New("暂时无法导出excel，请稍后重新"))
		return
	}
	time.Sleep(1 * time.Second)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filePath)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
}

func ExportActiveDetail(c *gin.Context, cols []map[string]string, items []models.ActiveDetailWeb) {
	excelTitleLine := utils.CreateExcelTitle(cols)
	var excelItems [][]string
	for _, v := range items {
		oneLine := []string{v.EventDay, v.Channel, v.ActiveMode, v.Pv, v.Uv, v.LastUpdate}
		excelItems = append(excelItems, oneLine)
	}
	filePath, err := utils.CreateExcelFile(excelTitleLine, excelItems)
	if err != nil {
		c.Error(errors.New("暂时无法导出excel，请稍后重新"))
		return
	}
	time.Sleep(1 * time.Second)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filePath)
	c.Header("Content-Transfer-Encoding", "binary")

	c.File(filePath)
}

func ExportPreserveDetail(c *gin.Context, cols []map[string]string, items []models.PreserveDetailWeb) {
	excelTitleLine := utils.CreateExcelTitle(cols)
	var excelItems [][]string
	for _, v := range items {
		oneLine := []string{v.EventTime, v.Channel, v.PassedWeekActive, v.Uv, v.NewUv, v.Day1Active, v.Day2Active, v.Day3Active, v.Day4Active, v.Day5Active, v.Day6Active, v.WeekActive, v.LastUpdate}
		excelItems = append(excelItems, oneLine)
	}
	filePath, err := utils.CreateExcelFile(excelTitleLine, excelItems)
	if err != nil {
		c.Error(errors.New("暂时无法导出excel，请稍后重新"))
		return
	}
	time.Sleep(1 * time.Second)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filePath)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
}

func ExportudtrstDetail(c *gin.Context, cols []map[string]string, items []models.UdtrstDetailWeb) {
	excelTitleLine := utils.CreateExcelTitle(cols)
	var excelItems [][]string
	for _, v := range items {
		oneLine := []string{v.EventDay, v.Channel, v.Pv, v.Uv, v.Rst0Pv, v.Rst0Uv, v.Rst1Pv, v.Rst1Uv, v.Rst3Pv, v.Rst3Uv, v.Rst4Pv, v.Rst4Uv, v.Rst7Pv, v.Rst7Uv, v.LastUpdate}
		excelItems = append(excelItems, oneLine)
	}
	filePath, err := utils.CreateExcelFile(excelTitleLine, excelItems)
	if err != nil {
		c.Error(errors.New("暂时无法导出excel，请稍后重新"))
		return
	}
	time.Sleep(1 * time.Second)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filePath)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
}

func ExportFeirarDetail(c *gin.Context, cols []map[string]string, items []models.FeirarDetailWeb) {
	excelTitleLine := utils.CreateExcelTitle(cols)
	var excelItems [][]string
	for _, v := range items {
		oneLine := []string{v.EventDay, v.Channel, v.EventKey, v.Pv, v.Uv, v.LastUpdate}
		excelItems = append(excelItems, oneLine)
	}
	filePath, err := utils.CreateExcelFile(excelTitleLine, excelItems)
	if err != nil {
		c.Error(errors.New("暂时无法导出excel，请稍后重新"))
		return
	}
	time.Sleep(1 * time.Second)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filePath)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(filePath)
}

type ExportArgs struct {
	ModuleName string `json:"type" form:"type" binding:"required"`
	PageID     int    `json:"page" form:"page" `
	PageCount  int    `json:"pcount" form:"pcount"`
	TimeStart  int64  `json:"ts" form:"ts" `
	TimeEnd    int64  `json:"te" form:"te"`
	Channels   string `json:"chn" form:"chn"`
	Name       string `json:"name" form:"name"`
}

func Export(c *gin.Context) {
	header := c.Request.Header
	glog.Info(header)
	var reqArgs ExportArgs
	err := c.ShouldBind(&reqArgs)
	if err != nil || reqArgs.TimeEnd <= 0 {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	switch reqArgs.ModuleName {
	case "install":
		cols := PingbackCenter.GetInstallDetailCols()
		items, _, err := PingbackCenter.GetInstallDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportInstall(c, cols, items)
		return
	case "uninstall":
		cols := PingbackCenter.GetUninstallDetailCols()
		items, _, err := PingbackCenter.GetUninstallDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportUninstall(c, cols, items)
		return

	case "active":
		cols := PingbackCenter.GetActiveDetailCols()
		items, _, err := PingbackCenter.GetActiveDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportActiveDetail(c, cols, items)
		return

	case "news":
		cols := PingbackCenter.GetNewsDetailCols()
		items, _, err := PingbackCenter.GetNewsDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd, "newsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportNewsDetail(c, cols, items)
		return

	case "feirar-righttipsshow":
		cols := PingbackCenter.GetNewsDetailCols()
		items, _, err := PingbackCenter.GetNewsDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd, "righttipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportNewsDetail(c, cols, items)
		return

	case "feirar-rightnewstipsshow":
		cols := PingbackCenter.GetNewsDetailCols()
		items, _, err := PingbackCenter.GetNewsDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd, "rightnewstipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportNewsDetail(c, cols, items)
		return

	case "feirar-taskbartipsshow":
		cols := PingbackCenter.GetNewsDetailCols()
		items, _, err := PingbackCenter.GetNewsDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd, "taskbartipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportNewsDetail(c, cols, items)
		return

	case "feirar-msgcentershow":
		cols := PingbackCenter.GetNewsDetailCols()
		items, _, err := PingbackCenter.GetNewsDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd, "msgcentershow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportNewsDetail(c, cols, items)
		return

	case "feirar-topcentertipsshow":
		cols := PingbackCenter.GetNewsDetailCols()
		items, _, err := PingbackCenter.GetNewsDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd, "topcentertipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportNewsDetail(c, cols, items)
		return

	case "feirar-traygametipsshow":
		cols := PingbackCenter.GetNewsDetailCols()
		items, _, err := PingbackCenter.GetNewsDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd, "traygametipsshow")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportNewsDetail(c, cols, items)
		return

	case "preserve":
		cols := PingbackCenter.GetPreserveDetailCols()
		items, _, err := PingbackCenter.GetPreserveDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportPreserveDetail(c, cols, items)
		return

	case "feirar":
		cols := PingbackCenter.GetFeirarDetailCols()
		items, _, err := PingbackCenter.GetFeirarDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportFeirarDetail(c, cols, items)
		return

	case "feirar-news":
		cols := PingbackCenter.GetFeirarNewsDetailCols()
		items, _, err := PingbackCenter.GetFeirarNewsDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportFeirarDetail(c, cols, items)
		return

	case "feirar-update":
		cols := PingbackCenter.GetFeirarDetailCols()
		items, _, err := PingbackCenter.GetFeirarUpdateDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportFeirarDetail(c, cols, items)
		return
	case "feirar-udtrst":
		cols := PingbackCenter.GetUdtrstDetailCols()
		items, _, err := PingbackCenter.GetUdtrstDetailItems(reqArgs.Channels, reqArgs.PageID, reqArgs.PageCount, reqArgs.TimeStart, reqArgs.TimeEnd)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
			return
		}
		ExportudtrstDetail(c, cols, items)
		return
	}

}

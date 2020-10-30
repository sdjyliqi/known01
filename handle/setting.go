package handle

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/sdjyliqi/feirars/utils"
	"net/http"
)

func HandleUpdate(c *gin.Context) {
	header := c.Request.Header
	glog.Info(header)
	var reqArgs utils.UpdateArgs
	err := c.ShouldBind(&reqArgs)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	reqArgs.City = c.GetHeader("IPCITY")
	strContent, err := SettingCenter.ClientUpdate(&reqArgs)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
		return
	}
	var item map[string]interface{}
	err = json.Unmarshal([]byte(strContent), &item)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": -1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

func HandleCity(c *gin.Context) {
	header := c.Request.Header
	glog.Info(header)
	fmt.Println(header)
	city := c.GetHeader("IPCITY")
	c.JSON(http.StatusOK, gin.H{"city": city})
}

package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

func JudgeMessage(c *gin.Context) {
	minLevelScore := 50
	type SubmitContent struct {
		Content string `json:"content"`
		Sender  string `json:"sender"`
	}
	reqJson := SubmitContent{}
	err := c.ShouldBindJSON(&reqJson)
	if err != nil {
		glog.Errorf("The request %+v is invalid,please check.", c.Request)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "bind json failed."})
		return
	}
	if len(reqJson.Content) < 1 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "You must submit the invalid message11"})
		return
	}
	score, suggest := baCenter.JudgeMessage(reqJson.Content, reqJson.Sender)
	flag := 0
	if score > minLevelScore {
		flag = 1
	}
	if score <= minLevelScore && score > 0 {
		flag = 2
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": gin.H{"suggest": suggest, "flag": flag}})
}

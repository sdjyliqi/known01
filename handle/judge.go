package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"known01/utils"
	"net/http"
	"time"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "time": time.Now()})
}

//JudgeMessage ...判断诈骗短信，如果score为未识别的分值，返回值中的score设置为0，flag需要设置为0
func JudgeMessage(c *gin.Context) {
	minLevelScore := 50
	var scoreRate = 0.0
	customPhone, website := "", ""
	type submitContent struct {
		Content string `json:"content"`
		Sender  string `json:"sender"`
	}
	reqJson := submitContent{}
	err := c.ShouldBindJSON(&reqJson)
	if err != nil {
		glog.Errorf("The request %+v is invalid,please check.", c.Request)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "bind json failed."})
		return
	}
	if len(reqJson.Content) < 1 || len(reqJson.Content) > 2048 || len(reqJson.Sender) > 128 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "You must submit the valid message(the length of sender or content is invalid.)"})
		return
	}
	score, reference := baCenter.JudgeMessage(reqJson.Content, reqJson.Sender)
	flag := 0
	if score > minLevelScore {
		flag = 1
	}
	if score <= minLevelScore && score != utils.OutsideKnown {
		flag = 2
	}
	if reference != nil {
		customPhone, website = reference.ManualPhone, reference.Website
	}
	if score != utils.OutsideKnown {
		scoreRate = float64(score) / 100
	}
	box := gin.H{"code": 0,
		"msg": "succ",
		"data": gin.H{
			"flag":    flag,
			"score":   scoreRate,
			"website": website,
			"hotline": customPhone,
		}}
	c.JSON(http.StatusOK, box)
}

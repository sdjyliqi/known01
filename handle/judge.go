package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

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
	if len(reqJson.Content) < 1 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "You must submit the invalid message11"})
		return
	}
	score, reference, suggest := baCenter.JudgeMessage(reqJson.Content, reqJson.Sender)
	flag := 0
	if score > minLevelScore {
		flag = 1
	}
	if score <= minLevelScore && score > 0 {
		flag = 2
	}
	if score > 0 && reference != nil {
		customPhone, website = reference.ManualPhone, reference.Website
	}

	scoreRate = float64(score) / 100
	box := gin.H{"code": 0,
		"msg": "succ",
		"data": gin.H{
			"suggest": suggest,
			"flag":    flag,
			"score":   scoreRate,
			"website": website,
			"hotline": customPhone,
		}}
	fmt.Printf("Request:%+v,Response:%+v\n", reqJson, box)
	c.JSON(http.StatusOK, box)
}

func JudgeMessageGET(c *gin.Context) {
	minLevelScore := 50
	var scoreRate = 0.0
	customPhone, website := "", ""
	submitContent := c.DefaultQuery("content", "")
	submitSender := c.DefaultQuery("sender", "")
	if len(submitContent) < 1 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "You must submit the invalid message"})
		return
	}
	score, reference, suggest := baCenter.JudgeMessage(submitContent, submitSender)
	flag := 0
	if score > minLevelScore {
		flag = 1
	}
	if score <= minLevelScore && score > 0 {
		flag = 2
	}
	if score > 0 && reference != nil {
		customPhone, website = reference.ManualPhone, reference.Website
	}

	scoreRate = float64(score) / 100
	box := gin.H{"code": 0,
		"msg": "succ",
		"data": gin.H{
			"suggest": suggest,
			"flag":    flag,
			"score":   scoreRate,
			"website": website,
			"hotline": customPhone,
		}}
	fmt.Printf("Request:%+v,Response:%+v\n", c.Request.URL, box)
	c.JSON(http.StatusOK, box)
}

package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JudgeMessage(c *gin.Context) {
	minLevelScore := 50
	message := c.DefaultQuery("content", "") //page 编码id
	sender := c.DefaultQuery("sender", "")   //page 编码id
	if message == "" {
		message = c.PostForm("content")
		sender = c.PostForm("sender")
	}
	if len(message) < 1 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "You must submit the invalid message"})
		return
	}
	score, suggest := baCenter.JudgeMessage(message, sender)
	flag := 0
	if score > minLevelScore {
		flag = 1
	}
	if score <= minLevelScore && score > 0 {
		flag = 2
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": gin.H{"suggest": suggest, "flag": flag}})
}

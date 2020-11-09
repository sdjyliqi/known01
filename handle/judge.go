package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JudgeMessage(c *gin.Context) {
	message := c.DefaultQuery("content", "") //page 编码id
	if len(message) < 1 {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "You must submit the invalid message"})
		return
	}
	score, suggest := baCenter.JudgeMessage(message)
	flag := 0
	if score > 0.5 {
		flag = 1
	}
	if score <= 0.5 && score > 0 {
		flag = 2
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": gin.H{"suggest": suggest, "flag": flag}})
}

package handle

import (
	"github.com/gin-gonic/gin"
	"known01/models"
	"known01/utils"
	"net/http"
)

func GetNews(c *gin.Context) {
	pageID := c.GetInt("page") //page 编码id
	items, err := models.News{}.GetItems(utils.GetMysqlClient(), pageID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": items, "is_end": len(items) < utils.PageEntry})
}

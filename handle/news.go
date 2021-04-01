package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/sdjyliqi/known01/model"
	"github.com/sdjyliqi/known01/utils"
	"net/http"
	"strconv"
)

//GetNews ...获取信息列表，优先从GET中获取页码，如果GET参数中，无发现，从POST中获取参数
func GetNews(c *gin.Context) {
	strPage := c.GetString("page")
	if strPage == "" {
		strPage = c.PostForm("page")
	}
	pageID, _ := strconv.Atoi(strPage)
	items, err := model.News{}.GetItems(utils.GetMysqlClient(), pageID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items, "is_end": len(items) < utils.PAGE_ENTRY})
}

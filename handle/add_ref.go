package handle

import (
	"github.com/gin-gonic/gin"
	"known01/model"
	"net/http"
	"strconv"
)

//GetRefList   ...分页查询短信鉴别参考数据表中的数据
func GetRefList(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	entry := c.DefaultQuery("entry", "0")
	page1, _ := strconv.Atoi(page)
	entry1, _ := strconv.Atoi(entry)
	if page1 <= 0 || entry1 <= 0 {
		c.JSON(http.StatusOK, gin.H{"code": 4004, "msg": "Page and entry must be integers greater than 0."})
		return
	}
	items, total, err := model.Reference{}.GetPages(page1, entry1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4005, "msg": "Failed to get list from table"})
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items, "number": total})
}

//GetRefItem ... 展示表中的一条记录的详细信息
func GetRefItem(c *gin.Context) {
	id := c.DefaultQuery("id", "0")
	id1, _ := strconv.Atoi(id)
	if id1 == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 4002, "msg": "the id must not be empty."})
		return
	}
	item, err := model.Reference{}.GetItemByID(id1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4005, "msg": "Failed to get list from table"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": item})

}

//UpdateRefItem   ...更新表中数据
func UpdateRefItem(c *gin.Context) {
	json := model.Reference{}
	err1 := c.BindJSON(&json)
	//log.Printf("%v", &json)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "parse error"})
		return
	}
	if json.Id == 0 {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Id can't be empty"})
		return
	}
	err2 := model.Reference{}.UpdateItemByID(json.Id, &json)
	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "Id doesn't exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})

}

//AddRefItem ...添加数据
func AddRefItem(c *gin.Context) {
	json := model.Reference{}
	err1 := c.BindJSON(&json)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "parse error"})
		return
	}
	if json.Name == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Name can't be empty"})
		return
	}
	err2 := model.Reference{}.InsertItemByID(json)
	//if errors.Is(err2, errors.New("already-existed")) {
	//	c.JSON(http.StatusOK, gin.H{"code": 4010, "msg": "The item already exists"})
	//	return
	//}
	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4015, "msg": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
}

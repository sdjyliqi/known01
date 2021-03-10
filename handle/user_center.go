package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"known01/model"
	"log"
	"net/http"
	"strconv"
)

//UCLogin ...用户登录
func UCLogin(c *gin.Context) {
	token := "000011111122222"
	name := c.DefaultPostForm("name", "")
	//nil
	password := c.DefaultPostForm("password", "")
	if name == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "the user name or password must not be empty."})
		return
	}
	invalidFlag, err := model.User{}.ChkPassword(name, password)
	//如果错误的时候，返回前端异常
	if err != nil {
		//todo
		glog.Info(invalidFlag)

	}
	//如果
	if invalidFlag == true {
		//返回用户名密码正确，token
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": token})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "the username or password invalid"})
}

//UCUsers ...分页查询用户
func UCUsers(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	entry := c.DefaultQuery("entry", "0")
	page1, _ := strconv.Atoi(page)
	entry1, _ := strconv.Atoi(entry)
	if page1 <= 0 || entry1 <= 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Page and entry must be integers greater than 0."})
		return
	}
	items, err := model.User{}.GetItems(page1, entry1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "Failed to get list from table"})
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
}

//UsersStatus ...改变用户状态，传入参数为用户登录ID
func UCUsersStatus(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	if name == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Name cann't be empty"})
		return
	}
	res, _ := model.User{}.ModifyEnable(name)
	if res == true {
		//用户状态修改成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "user doesn't exist"})

}
func UCAddUsers(c *gin.Context) {
	json := model.AddUser{}
	err := c.BindJSON(&json)
	log.Printf("%v", &json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "parse error"})
		return
	}
	if json.Name == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4005, "msg": "Name cannot be empty"})
		return
	}
	res, _ := model.User{}.AddData(json)
	if res == true {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "User already exists"})

}

package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"known01/model"
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
	items, _ := model.User{}.GetItems(page1, entry1)
	//if err != errors.New("not-find") {
	//	c.JSON(http.StatusOK, gin.H{"code": 4002, "msg": err.Error()})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
}

//UsersStatus ...改变用户状态，传入参数为用户登录ID
func UsersStatus(c *gin.Context) {
	name := c.PostForm("name")
	enable := c.PostForm("enable")
	if enable == "1" || enable == "0" {
		res, _ := model.User{}.ModifyEnable(name)
		if res == true {
			//用户状态修改成功
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "user doesn't exist"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Unknown parameter.", "enable": enable})
		return
	}

}

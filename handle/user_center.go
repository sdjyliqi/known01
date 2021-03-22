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
	keyid := c.DefaultPostForm("keyid", "")
	//nil
	password := c.DefaultPostForm("password", "")
	if keyid == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "the user name or password must not be empty."})
		return
	}
	invalidFlag, err := model.User{}.ChkPassword(keyid, password)
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

//UCShowInformation ... 展示用户详细信息
func UCShowInformation(c *gin.Context) {
	keyid := c.DefaultPostForm("keyid", "")
	if keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "keyid cann't be empty"})
		return
	}
	res, err := model.User{}.ShowInf(keyid)
	if err != nil {
		//未找到该用户
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "user doesn't exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": res})
}

//UsersStatus ...改变用户状态，传入参数为用户登录ID
func UCUsersStatus(c *gin.Context) {
	keyid := c.DefaultPostForm("keyid", "")
	if keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Keyid cann't be empty"})
		return
	}
	res, _ := model.User{}.ModifyEnable(keyid)
	if res == true {
		//用户状态修改成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "user doesn't exist"})

}

//UCAddUsers ...添加用户
func UCAddUsers(c *gin.Context) {
	json := model.AddUser{}
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "parse error"})
		return
	}
	if json.Keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4005, "msg": "Keyid cannot be empty"})
		return
	}
	res, _ := model.User{}.AddData(json)
	if res == true {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "User already exists"})
}

//UCResetPassword  ... 用户重置密码
func UCResetPassword(c *gin.Context) {
	keyid := c.DefaultPostForm("keyid", "")
	if keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Keyid cann't be empty"})
		return
	}
	res, _ := model.User{}.ResetPas(keyid)
	if res == true {
		//用户密码重置成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "user doesn't exist"})
}

//UCChangePassword   ... 用户修改密码
func UCChangePassword(c *gin.Context) {
	keyid := c.DefaultPostForm("keyid", "")
	oldpas := c.DefaultPostForm("oldpas", "")
	newpas := c.DefaultPostForm("newpas", "")
	confirmpas := c.DefaultPostForm("confirmpas", "")
	if keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Keyid cann't be empty"})
		return
	}
	if oldpas == "" || newpas == "" || confirmpas == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Password cannot be empty"})
		return
	}
	if oldpas == newpas || oldpas == confirmpas {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "The new password and the old password cannot be the same"})
		return
	}
	if newpas != confirmpas {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "The new password is not the same as the confirmation password"})
		return
	}
	_, err := model.User{}.ShowInf(keyid) //通过keyid值查询用户是否在数据库中
	if err != nil {
		//未找到该用户
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "user doesn't exist"})
		return
	}
	invalidFlag, _ := model.User{}.ChkPassword(keyid, oldpas)
	if invalidFlag != true {
		//如果旧密码错误
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Wrong password"})
		return
	}
	res, _ := model.User{}.ChangePas(keyid, newpas)
	if res == true {
		//用户修改密码成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "user doesn't exist"})
}

//UCUpdateInformation   ...用户更新个人信息
func UCUpdateInformation(c *gin.Context) {
	json := model.UserUpdate{}
	err := c.BindJSON(&json)
	log.Printf("%v", &json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "parse error"})
		return
	}
	if json.Keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "Keyid cann't be empty"})
		return
	}
	res, _ := model.User{}.UpdateInf(json)
	if res == true {
		//用户个人信息更新成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "user doesn't exist"})
}

package handle

import (
	"github.com/gin-gonic/gin"
	"known01/model"
	"log"
	"net/http"
	"strconv"
)

//ChangePasInf  ... 用户自己修改密码提交内容
type ChangePasInf struct {
	Keyid      string `json:"keyid" xorm:"not null pk comment('api请求分配的账号id') unique VARCHAR(64)"`
	Oldpas     string `json:"oldpas"`
	Newpas     string `json:"newpas"`
	Confirmpas string `json:"confirmpas"`
}

//UCLogin ...用户登录
func UCLogin(c *gin.Context) {
	token := "000011111122222"
	frontmsg := model.User{} //前端传来的数据
	err1 := c.BindJSON(&frontmsg)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "parse error"})
		return
	}
	if frontmsg.Keyid == "" || frontmsg.Password == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4002, "msg": "the username or password must not be empty."})
		return
	}
	item, err2 := model.User{}.GetItemById(frontmsg.Keyid)
	//如果错误的时候，返回前端异常
	if err2 == nil && frontmsg.Password == item.Password {
		data := map[string]string{
			"roles": item.Roles,
			"token": token,
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": data})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4003, "msg": "the username or password wrong"})
}

//UCRoles  ...返回用户角色，管理员或普通用户
func UCRoles(c *gin.Context) {
	frontmsg := model.User{}
	err1 := c.BindJSON(&frontmsg)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "parse error"})
		return
	}
	if frontmsg.Keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid can't be empty"})
		return
	}
	item, err2 := model.User{}.GetItemById(frontmsg.Keyid)
	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4005, "msg": "Failed to get list from table"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": item.Roles})
}

//UCUsers ...分页查询用户
func UCUsers(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	entry := c.DefaultQuery("entry", "0")
	page1, _ := strconv.Atoi(page)
	entry1, _ := strconv.Atoi(entry)
	if page1 <= 0 || entry1 <= 0 {
		c.JSON(http.StatusOK, gin.H{"code": 4004, "msg": "Page and entry must be integers greater than 0."})
		return
	}
	items, err := model.User{}.GetItems(page1, entry1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4005, "msg": "Failed to get list from table"})
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": items})
}

//UCShowInformation ... 展示用户详细信息
func UCShowInformation(c *gin.Context) {
	keyid := c.DefaultQuery("keyid", "")
	if keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "keyid can not be empty"})
		return
	}
	res, err := model.User{}.ShowInf(keyid)
	if err != nil {
		//未找到该用户
		c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "user doesn't exist"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ", "data": res})
}

//UsersStatus ...改变用户状态，传入参数为用户登录ID
func UCUsersStatus(c *gin.Context) {
	keyid := c.DefaultQuery("keyid", "")
	if keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid can't be empty"})
		return
	}
	res, _ := model.User{}.ModifyEnable(keyid)
	if res == true {
		//用户状态修改成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "user doesn't exist"})

}

//UCAddUsers ...添加用户
func UCAddUsers(c *gin.Context) {
	json := model.User{}
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "parse error"})
		return
	}
	if json.Keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid can't be empty"})
		return
	}
	if json.Manager == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4008, "msg": "Manager can't be empty"})
		return
	}
	if json.Roles != "admin" && json.Roles != "editor" {
		c.JSON(http.StatusOK, gin.H{"code": 4009, "msg": "Roles must be admin or editor"})
		return
	}
	res, _ := model.User{}.InsertItem(json)
	if res == true {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4010, "msg": "User already exists"})
}

//UCResetPassword  ... 管理员重置用户密码
func UCResetPassword(c *gin.Context) {
	keyid := c.DefaultQuery("keyid", "")
	if keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid can't be empty"})
		return
	}
	res, _ := model.User{}.ResetPas(keyid)
	if res == true {
		//用户密码重置成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "user doesn't exist"})
}

//UCChangePassword   ... 用户修改密码
func UCChangePassword(c *gin.Context) {
	pasinf := ChangePasInf{} //用户输入的修改密码信息
	err1 := c.BindJSON(&pasinf)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "parse error"})
		return
	}
	if pasinf.Keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid can't be empty"})
		return
	}
	if pasinf.Oldpas == "" || pasinf.Newpas == "" || pasinf.Confirmpas == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4011, "msg": "Password can't be empty"})
		return
	}
	if pasinf.Oldpas == pasinf.Newpas || pasinf.Oldpas == pasinf.Confirmpas {
		c.JSON(http.StatusOK, gin.H{"code": 4012, "msg": "The new password and the old password can't be the same"})
		return
	}
	if pasinf.Newpas != pasinf.Confirmpas {
		c.JSON(http.StatusOK, gin.H{"code": 4013, "msg": "The new password is not the same as the confirmation password"})
		return
	}
	item, err2 := model.User{}.GetItemById(pasinf.Keyid) //通过keyid值查询用户是否在数据库中
	if err2 != nil {
		//未找到该用户
		c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "user doesn't exist"})
		return
	}
	if item.Password != pasinf.Oldpas {
		//如果旧密码错误
		c.JSON(http.StatusOK, gin.H{"code": 4014, "msg": "Wrong password"})
		return
	}
	res, _ := model.User{}.ChangePas(pasinf.Keyid, pasinf.Newpas)
	if res == true {
		//用户修改密码成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "user doesn't exist"})
}

//UCUpdateInformation   ...用户更新个人信息
func UCUpdateInformation(c *gin.Context) {
	json := model.User{}
	err := c.BindJSON(&json)
	log.Printf("%v", &json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "parse error"})
		return
	}
	if json.Keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid can't be empty"})
		return
	}
	res, _ := model.User{}.EditorUpdateItemById(json)
	if res == true {
		//用户个人信息更新成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "user doesn't exist"})
}

//UCAdminUpdateItem   ...管理员更新用户信息
func UCAdminUpdateItem(c *gin.Context) {
	json := model.User{}
	err := c.BindJSON(&json)
	log.Printf("%v", &json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "parse error"})
		return
	}
	if json.Keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid can't be empty"})
		return
	}
	if json.Roles != "admin" && json.Roles != "editor" {
		c.JSON(http.StatusOK, gin.H{"code": 4009, "msg": "Roles must be admin or editor"})
		return
	}
	res, _ := model.User{}.AdminUpdateItem(json)
	if res == true {
		//用户个人信息更新成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "succ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "user doesn't exist"})
}

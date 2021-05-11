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
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "解析错误"})
		return
	}
	if frontmsg.Keyid == "" || frontmsg.Password == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4002, "msg": "用户名或密码不能为空"})
		return
	}
	item, err2 := model.User{}.GetItemById(frontmsg.Keyid)
	//如果错误的时候，返回前端异常
	if err2 == nil && frontmsg.Password == item.Password {
		if item.Enable != 1 {
			c.JSON(http.StatusOK, gin.H{"code": 4015, "msg": "该用户已被禁用"})
			return
		} else {
			data := map[string]string{
				"roles": item.Roles,
				"token": token,
			}
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "登录成功", "data": data})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 4003, "msg": "用户名或密码错误"})
}

//UCRoles  ...返回用户角色，管理员或普通用户
func UCRoles(c *gin.Context) {
	frontmsg := model.User{}
	err1 := c.BindJSON(&frontmsg)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "解析错误"})
		return
	}
	if frontmsg.Keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid 不能为空"})
		return
	}
	item, err2 := model.User{}.GetItemById(frontmsg.Keyid)
	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4005, "msg": "获取数据失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "成功获取用户角色", "data": item.Roles})
}

//UCUsers ...分页查询用户
func UCUsers(c *gin.Context) {
	page := c.DefaultQuery("page", "0")
	entry := c.DefaultQuery("entry", "0")
	page1, _ := strconv.Atoi(page)
	entry1, _ := strconv.Atoi(entry)
	if page1 <= 0 || entry1 <= 0 {
		c.JSON(http.StatusOK, gin.H{"code": 4004, "msg": "页码和每页展示数必须为大于0的整数"})
		return
	}
	items, total, err := model.User{}.GetItems(page1, entry1)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4005, "msg": "获取数据失败"})
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "查询成功", "data": items, "number": total})
}

//UCShowInformation ... 展示用户详细信息
func UCShowInformation(c *gin.Context) {
	keyid := c.DefaultQuery("keyid", "")
	if keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "keyid 不能为空"})
		return
	}
	res, err := model.User{}.ShowInf(keyid)
	if err != nil {
		//未找到该用户
		c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "该用户不存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "用户详细信息获取成功", "data": res})
}

//UsersStatus ...改变用户状态，传入参数为用户登录ID
func UCUsersStatus(c *gin.Context) {
	keyid := c.DefaultQuery("keyid", "")
	if keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid 不能为空"})
		return
	}
	res, _ := model.User{}.ModifyEnable(keyid)
	item, _ := model.User{}.GetItemById(keyid)
	if res == true {
		//用户状态修改成功
		if item.Enable == 1 {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "该用户已激活"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "该用户已被禁用"})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "该用户不存在"})

}

//UCAddUsers ...添加用户
func UCAddUsers(c *gin.Context) {
	json := model.User{}
	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "解析错误"})
		return
	}
	if json.Keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid不能为空"})
		return
	}
	if json.Manager == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4008, "msg": "负责人不能为空"})
		return
	}
	if json.Roles != "admin" && json.Roles != "editor" {
		c.JSON(http.StatusOK, gin.H{"code": 4009, "msg": "角色必须为admin或editor"})
		return
	}
	res, _ := model.User{}.InsertItem(json)
	if res == true {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "用户添加成功"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4010, "msg": "该用户已存在"})
}

//UCResetPassword  ... 管理员重置用户密码
func UCResetPassword(c *gin.Context) {
	keyid := c.DefaultQuery("keyid", "")
	if keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid 不能为空"})
		return
	}
	res, _ := model.User{}.ResetPas(keyid)
	if res == true {
		//用户密码重置成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "用户密码重置成功"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "该用户不存在"})
}

//UCChangePassword   ... 用户修改密码
func UCChangePassword(c *gin.Context) {
	pasinf := ChangePasInf{} //用户输入的修改密码信息
	err1 := c.BindJSON(&pasinf)
	if err1 != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "解析错误"})
		return
	}
	if pasinf.Keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid 不能为空"})
		return
	}
	if pasinf.Oldpas == "" || pasinf.Newpas == "" || pasinf.Confirmpas == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4011, "msg": "密码不能为空"})
		return
	}
	if pasinf.Oldpas == pasinf.Newpas || pasinf.Oldpas == pasinf.Confirmpas {
		c.JSON(http.StatusOK, gin.H{"code": 4012, "msg": "新密码和旧密码不能相同"})
		return
	}
	if pasinf.Newpas != pasinf.Confirmpas {
		c.JSON(http.StatusOK, gin.H{"code": 4013, "msg": "新密码和确认密码不同"})
		return
	}
	item, err2 := model.User{}.GetItemById(pasinf.Keyid) //通过keyid值查询用户是否在数据库中
	if err2 != nil {
		//未找到该用户
		c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "该用户不存在"})
		return
	}
	if item.Password != pasinf.Oldpas {
		//如果旧密码错误
		c.JSON(http.StatusOK, gin.H{"code": 4014, "msg": "密码错误"})
		return
	}
	res, _ := model.User{}.ChangePas(pasinf.Keyid, pasinf.Newpas)
	if res == true {
		//用户修改密码成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "密码修改成功"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "该用户不存在"})
}

//UCUpdateInformation   ...用户更新个人信息
func UCUpdateInformation(c *gin.Context) {
	json := model.User{}
	err := c.BindJSON(&json)
	log.Printf("%v", &json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "解析错误"})
		return
	}
	if json.Keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid 不能为空"})
		return
	}
	res, _ := model.User{}.EditorUpdateItemById(json)
	if res == true {
		//用户个人信息更新成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "信息修改成功"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "该用户不存在"})
}

//UCAdminUpdateItem   ...管理员更新用户信息
func UCAdminUpdateItem(c *gin.Context) {
	json := model.User{}
	err := c.BindJSON(&json)
	log.Printf("%v", &json)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 4001, "msg": "解析错误"})
		return
	}
	if json.Keyid == "" {
		c.JSON(http.StatusOK, gin.H{"code": 4006, "msg": "Keyid 不能为空"})
		return
	}
	if json.Roles != "admin" && json.Roles != "editor" {
		c.JSON(http.StatusOK, gin.H{"code": 4009, "msg": "角色必须为admin或editor"})
		return
	}
	res, _ := model.User{}.AdminUpdateItem(json)
	if res == true {
		//用户个人信息更新成功
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "信息修改成功"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 4007, "msg": "该用户不存在"})
}

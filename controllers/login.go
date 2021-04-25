package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"mall/models"
	"mall/utils"
	"fmt"
)

type LoginController struct {
	beego.Controller
}

// 用户登入
func (c *LoginController) Login() {
	username := c.GetString("username")
	password := c.GetString("password")
	fmt.Printf("用户登录, username#%s, password#%s\n", username, password)

	result := utils.Ok()
	user := models.User{}
	err := models.O.QueryTable("user").Filter("Username", username).Filter("Password", password).One(&user)
	if err == nil {
		fmt.Println(user)
		var data = map[string]interface{}{"uid": user.Id, "username": user.Username, "role": user.Role, "phone": user.PhoneNumber}
		result.Data = data
		c.SetSession("uid", user.Id)
	} else {
		fmt.Println(err.Error())
		result = utils.Error("账号或密码错误")
	}
	c.Data["json"] = &result
	c.ServeJSON()
}

// 用户登出
func (c *LoginController) Logout() {
	c.DelSession("uid")
	result := utils.Ok()
	c.Data["json"] = &result
	c.ServeJSON()
}

package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"mall/models"
	"mall/utils"
)

type UserController struct {
	beego.Controller
}

// 用户注册
func (c *UserController) Register() {
	result := utils.Ok()
	var user models.User
	requestBody := c.Ctx.Input.RequestBody
	fmt.Printf("用户注册, params#%s\n", string(requestBody))
	err := json.Unmarshal(requestBody, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	u := models.User{}
	err = models.O.QueryTable("user").Filter("Username", user.Username).One(&u)
	if err == nil {
		result = utils.Error("账号已存在!")
	} else {
		models.O.Insert(&user)
	}
	c.Data["json"] = &result
	c.ServeJSON()
}

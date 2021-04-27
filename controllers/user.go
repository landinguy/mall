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

// 查询客户信息
func (c *UserController) GetById() {
	result := utils.Ok()
	id := c.GetString(":id")

	fmt.Printf("id#%s\n", id)
	user := models.User{}
	models.O.QueryTable("user").Filter("Id", id).One(&user)
	result.Data = user
	c.Data["json"] = &result
	c.ServeJSON()
}

// 保存用户
func (c *UserController) SaveUser() {
	result := utils.Ok()
	var user models.User
	requestBody := c.Ctx.Input.RequestBody
	fmt.Printf("保存用户, params#%s\n", string(requestBody))
	json.Unmarshal(requestBody, &user)
	if user.Id != 0 {
		models.O.Update(&user, "Username", "PhoneNumber", "AgriculturalCooperative", "Address", "WorkNumber")
	} else {
		exist := models.User{}
		err := models.O.QueryTable("user").Filter("Username", user.Username).One(&exist)
		if err == nil {
			result = utils.Error("账号已存在!")
		} else {
			models.O.Insert(&user)
		}
	}
	c.Data["json"] = &result
	c.ServeJSON()
}

// 查询用户
func (c *UserController) GetUser() {
	result := utils.Ok()
	var reqDto models.UserQueryReqDto
	requestBody := c.Ctx.Input.RequestBody
	fmt.Printf("查询用户, params#%s\n", string(requestBody))
	json.Unmarshal(requestBody, &reqDto)
	queryWrapper := models.O.QueryTable("user")
	total, _ := queryWrapper.Count()
	var list []*models.User
	queryWrapper.Offset((reqDto.PageNo - 1) * reqDto.PageSize).Limit(reqDto.PageSize).OrderBy("-id").All(&list)
	result.Data = map[string]interface{}{"total": total, "list": list}
	c.Data["json"] = &result
	c.ServeJSON()
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

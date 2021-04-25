package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"mall/models"
	"mall/utils"
)

type CustomerController struct {
	beego.Controller
}

// 添加客户信息
func (c *CustomerController) Add() {
	result := utils.Ok()
	var customer models.Customer
	requestBody := c.Ctx.Input.RequestBody
	fmt.Printf("添加客户信息, params#%s\n", string(requestBody))
	err := json.Unmarshal(requestBody, &customer)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}

	uid := c.GetSession("uid")
	exist := models.Customer{}
	err = models.O.QueryTable("customer").Filter("UserId", uid).One(&exist)
	if err == nil {
		result = utils.Error("客户信息已存在!")
	} else {
		customer.UserId = uid.(int)
		models.O.Insert(&customer)
	}
	c.Data["json"] = &result
	c.ServeJSON()
}

// 查询客户信息
func (c *CustomerController) QueryById() {
	result := utils.Ok()
	userId := c.GetString(":uid")

	fmt.Printf("userId#%s\n", userId)

	customer := models.Customer{}
	err := models.O.QueryTable("customer").Filter("UserId", userId).One(&customer)
	if err == nil {
		result.Data = customer
	}
	c.Data["json"] = &result
	c.ServeJSON()
}

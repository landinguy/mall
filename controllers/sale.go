package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"mall/models"
	"mall/utils"
)

type SaleController struct {
	beego.Controller
}

// 查询商品销售信息
func (c *SaleController) GetDetail() {
	result := utils.Ok()
	var params models.PurchaseDetail
	requestBody := c.Ctx.Input.RequestBody
	fmt.Printf("登记采购信息, params#%s\n", string(requestBody))
	json.Unmarshal(requestBody, &params)
	models.O.Insert(&params)
	c.Data["json"] = &result
	c.ServeJSON()
}

// 登记商品销售信息
func (c *SaleController) AddDetail() {
	result := utils.Ok()
	var params models.SaleDetail
	requestBody := c.Ctx.Input.RequestBody
	fmt.Printf("登记销售信息, params#%s\n", string(requestBody))
	json.Unmarshal(requestBody, &params)
	models.O.Insert(&params)
	c.Data["json"] = &result
	c.ServeJSON()
}

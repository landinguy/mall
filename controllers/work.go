package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/golang-module/carbon"
	"mall/models"
	"mall/utils"
)

type WorkController struct {
	beego.Controller
}

// 上工登记
func (c *WorkController) AddWorkRecord() {
	result := utils.Ok()
	var params models.WorkRecord
	requestBody := c.Ctx.Input.RequestBody
	fmt.Printf("上工登记, params#%s\n", string(requestBody))
	json.Unmarshal(requestBody, &params)
	today := carbon.Now().ToDateString()

	exist := models.WorkRecord{}
	err := models.O.QueryTable("WorkRecord").Filter("UserId", params.UserId).Filter("Date__gt", today).One(&exist)
	if err == nil {
		result = utils.Error("今天已上工登记")
	} else {
		models.O.Insert(&params)
	}
	c.Data["json"] = &result
	c.ServeJSON()
}

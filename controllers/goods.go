package controllers

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"mall/models"
	"mall/utils"
)

type GoodsController struct {
	beego.Controller
}

// 查询商品
func (c *GoodsController) GetGoods() {
	result := utils.Ok()
	var reqDto models.GoodsQueryReqDto
	requestBody := c.Ctx.Input.RequestBody
	fmt.Printf("查询商品, params#%s\n", string(requestBody))
	json.Unmarshal(requestBody, &reqDto)
	queryWrapper := models.O.QueryTable("goods")
	if reqDto.GoodsName != "" {
		queryWrapper = queryWrapper.Filter("goods_name__contains", reqDto.GoodsName)
	}

	total, _ := queryWrapper.Count()
	var list []*models.Goods
	queryWrapper.Offset((reqDto.PageNo - 1) * reqDto.PageSize).Limit(reqDto.PageSize).OrderBy("-id").All(&list)
	result.Data = map[string]interface{}{"total": total, "list": list}
	c.Data["json"] = &result
	c.ServeJSON()
}

// 添加商品
func (c *GoodsController) AddGoods() {
	result := utils.Ok()
	var goods models.Goods
	requestBody := c.Ctx.Input.RequestBody
	fmt.Printf("添加商品, params#%s\n", string(requestBody))
	err := json.Unmarshal(requestBody, &goods)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	models.O.Insert(&goods)
	c.Data["json"] = &result
	c.ServeJSON()
}

// 添加商品分类
func (c *GoodsController) AddGoodsType() {
	result := utils.Ok()
	var goodsType models.GoodsType
	requestBody := c.Ctx.Input.RequestBody
	fmt.Printf("添加商品分类, params#%s\n", string(requestBody))
	err := json.Unmarshal(requestBody, &goodsType)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	exist := models.GoodsType{}
	err = models.O.QueryTable("goods_type").Filter("Name", goodsType.Name).One(&exist)
	if err == nil {
		result = utils.Error("商品分类已存在!")
	} else {
		models.O.Insert(&goodsType)
	}
	c.Data["json"] = &result
	c.ServeJSON()
}

// 查询商品分类
func (c *GoodsController) GetGoodsType() {
	result := utils.Ok()
	var list []*models.GoodsType
	models.O.QueryTable("goods_type").All(&list)
	result.Data = list
	c.Data["json"] = &result
	c.ServeJSON()
}

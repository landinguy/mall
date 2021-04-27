package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

var O orm.Ormer

// 用户实体
type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phoneNumber"`
	//农业合作社
	AgriculturalCooperative string `json:"agriculturalCooperative"`
	//地址
	Address string `json:"address"`
	//工号
	WorkNumber string `json:"workNumber"`
}

// 上工记录实体
type WorkRecord struct {
	Id int `json:"id"`
	//工作人员id
	UserId int `json:"userId"`
	//工作人员名称
	Username string `json:"username"`
	//工号
	WorkNumber string `json:"workNumber"`
	//联系方式
	PhoneNumber string `json:"phoneNumber"`
	//上工日期
	Date time.Time `orm:"auto_now_add;type(datetime)" json:"date"`
}

// 商品实体
type Goods struct {
	Id int `json:"id"`
	//商品名称
	GoodsName string `json:"goodsName"`
	//商品类型
	GoodsType string `json:"goodsType"`
	//商品价格（单位/元）
	Price float32 `json:"price"`
	//商品添加时间
	Date time.Time `orm:"auto_now_add;type(datetime)" json:"date"`
}

// 商品分类实体
type GoodsType struct {
	Id int `json:"id"`
	//类型名称
	Name string `json:"name"`
}

// 订单实体
type Order struct {
	Id int `json:"id"`
	//用户id
	UserId int `json:"userId"`
	//订单状态（1.待发货 2.已发货 3.确认收货）
	Status int `json:"status"`
}

// 订单明细实体
type OrderDetail struct {
	Id int `json:"id"`
	//订单id
	OrderId int `json:"orderId"`
	//商品id
	GoodsId int `json:"goodsId"`
	//商品数量
	Number int `json:"number"`
}

// 采购明细
type PurchaseDetail struct {
	Id int `json:"id"`
	//用户id
	UserId int `json:"userId"`
	//商品名
	GoodsName string `json:"goodsName"`
	//商品类别
	GoodsType string `json:"goodsType"`
	//生产日期
	ProductionDate string `json:"productionDate"`
	//采购日期
	PurchaseDate string `orm:"json:"purchaseDate`
	//采购数量
	PurchaseNumber int `json:"purchaseNumber"`
	//厂家地址
	ManufacturerAddress string `json:"manufacturerAddress"`
	//厂家联系方式
	ManufacturerPhone string `json:"manufacturerPhone"`
}

// 仓储明细
type StorageDetail struct {
	Id int `json:"id"`
	//用户id
	UserId int `json:"userId"`
	//商品id
	GoodsId int `json:"goodsId"`
	//入仓日期
	Date time.Time `orm:"auto_now_add;type(datetime)" json:"date"`
	//存放仓号
	WarehouseNumber int `json:"warehouseNumber"`
	//入仓数量
	Number int `json:"number"`
	//总存储量
	Total int `json:"total"`
}

// 销售明细
type SaleDetail struct {
	Id int `json:"id"`
	//用户id
	UserId int `json:"userId"`
	//商品id
	GoodsId int `json:"goodsId"`
	//售出日期
	Date time.Time `orm:"auto_now_add;type(datetime)" json:"date"`
	//售出数量
	Number int `json:"number"`
	//买家姓名
	CustomerName string `json:"customerName"`
	//买家联系方式
	CustomerPhone string `json:"customerPhone"`
}

func init() {
	orm.RegisterModel(
		new(User), new(WorkRecord), new(Goods), new(GoodsType), new(Order), new(OrderDetail), new(PurchaseDetail), new(StorageDetail), new(SaleDetail),
	)
}

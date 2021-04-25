package models

import "github.com/beego/beego/v2/client/orm"

var O orm.Ormer

// 用户实体
type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phone_number"`
}

// 工作人员信息实体
type Worker struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	//工号
	WorkNumber string `json:"work_number"`
}

// 上工记录实体
type WorkRecord struct {
	Id int `json:"id"`
	//工作人员id
	WorkerId int `json:"worker_id"`
	//工作人员名称
	WorkerName string `json:"worker_name"`
	//工号
	WorkNumber string `json:"work_number"`
	//联系方式
	PhoneNumber string `json:"phone_number"`
	//上工日期
	Date string `json:"date"`
}

// 顾客实体
type Customer struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	//农业合作社
	AgriculturalCooperative string `json:"agricultural_cooperative"`
	//地址
	Address string `json:"address"`
}

// 商品实体
type Goods struct {
	Id int `json:"id"`
	//商品名称
	GoodsName string `json:"goods_name"`
	//商品类型
	GoodsType string `json:"goods_type"`
	//商品价格（单位/元）
	Price float32 `json:"price"`
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
	UserId int `json:"user_id"`
	//顾客id
	CustomerId int `json:"customer_id"`
	//订单状态（1.待发货 2.已发货 3.确认收货）
	Status int `json:"status"`
}

// 订单明细实体
type OrderDetail struct {
	Id int `json:"id"`
	//订单id
	OrderId int `json:"order_id"`
	//商品id
	GoodsId int `json:"goods_id"`
	//商品数量
	Number int `json:"number"`
}

// 采购明细
type PurchaseDetail struct {
	Id int `json:"id"`
	//工人id
	WorkerId int `json:"worker_id"`
	//商品id
	GoodsId int `json:"goods_id"`
	//生产日期
	ProductionDate string `json:"production_date"`
	//采购日期
	Date string `json:"date"`
	//采购数量
	PurchaseNumber int `json:"purchase_number"`
	//厂家地址
	ManufacturerAddress string `json:"manufacturer_address"`
	//厂家联系方式
	ManufacturerPhone string `json:"manufacturer_phone"`
}

// 仓储明细
type StorageDetail struct {
	Id int `json:"id"`
	//工人id
	WorkerId int `json:"worker_id"`
	//商品id
	GoodsId int `json:"goods_id"`
	//入仓日期
	Date string `json:"date"`
	//存放仓号
	WarehouseNumber int `json:"warehouse_number"`
	//入仓数量
	Number int `json:"number"`
	//总存储量
	Total int `json:"total"`
}

// 销售明细
type SaleDetail struct {
	Id int `json:"id"`
	//工人id
	WorkerId int `json:"worker_id"`
	//商品id
	GoodsId int `json:"goods_id"`
	//售出日期
	Date string `json:"date"`
	//售出数量
	Number int `json:"number"`
	//买家姓名
	CustomerName string `json:"customer_name"`
	//买家联系方式
	CustomerPhone string `json:"customer_phone"`
}

func init() {
	orm.RegisterModel(
		new(User), new(Worker), new(WorkRecord), new(Customer), new(Goods), new(GoodsType), new(Order), new(OrderDetail),
		new(PurchaseDetail), new(StorageDetail), new(SaleDetail),
	)
}

package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"mall/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")

	//user
	beego.Router("/user/register", &controllers.UserController{}, "post:Register")
	beego.Router("/user", &controllers.UserController{}, "post:GetUser")
	beego.Router("/user/save", &controllers.UserController{}, "post:SaveUser")
	beego.Router("/user/:id", &controllers.UserController{}, "get:GetById")

	//goods
	beego.Router("/goods/type/add", &controllers.GoodsController{}, "post:AddGoodsType")
	beego.Router("/goods/type", &controllers.GoodsController{}, "get:GetGoodsType")
	beego.Router("/goods/save", &controllers.GoodsController{}, "post:SaveGoods")
	beego.Router("/goods", &controllers.GoodsController{}, "post:GetGoods")
	beego.Router("/goods/:id", &controllers.GoodsController{}, "delete:DeleteGoods")

	//work
	beego.Router("/workRecord/save", &controllers.WorkController{}, "post:AddWorkRecord")

	//purchase
	beego.Router("/purchaseDetail/save", &controllers.PurchaseController{}, "post:AddDetail")
	beego.Router("/purchaseDetail", &controllers.PurchaseController{}, "post:GetDetail")

	//storage
	beego.Router("/storageDetail/save", &controllers.StorageController{}, "post:AddDetail")
	beego.Router("/storageDetail", &controllers.StorageController{}, "post:GetDetail")

	//sale
	beego.Router("/saleDetail/save", &controllers.SaleController{}, "post:AddDetail")
	beego.Router("/saleDetail", &controllers.SaleController{}, "post:GetDetail")
}

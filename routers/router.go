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

	//customer
	beego.Router("/customer/add", &controllers.CustomerController{}, "post:Add")
	beego.Router("/customer/:uid", &controllers.CustomerController{}, "get:QueryById")

	//goods
	beego.Router("/goods/type/add", &controllers.GoodsController{}, "post:AddGoodsType")
	beego.Router("/goods/type", &controllers.GoodsController{}, "get:GetGoodsType")
	beego.Router("/goods/add", &controllers.GoodsController{}, "post:AddGoods")
	beego.Router("/goods", &controllers.GoodsController{}, "post:GetGoods")
}

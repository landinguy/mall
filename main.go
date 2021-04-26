package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	_ "github.com/go-sql-driver/mysql"
	"mall/models"
	_ "mall/routers"
)

func init() {
	//注册数据驱动
	orm.RegisterDriver("mysql", orm.DRMySQL) // mysql / sqlite3 / postgres 这三种是beego默认已经注册过的，所以可以无需设置
	//注册数据库 ORM 必须注册一个别名为 default 的数据库，作为默认使用
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/mall?charset=utf8&loc=Local")
	//自动创建表 参数二为是否开启创建表   参数三是否更新表
	orm.RunSyncdb("default", false, true)
	models.O = orm.NewOrm()
}

func main() {
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{"http://10.198.70.1:8088"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	web.Run()
}

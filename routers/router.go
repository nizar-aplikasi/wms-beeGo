package routers

import (
	Login 		"wms/controllers/login"
	Incoming 		"wms/controllers/incoming"
	Outcoming 		"wms/controllers/outcoming"
	Dashboard 	"wms/controllers/dashboard"
	Users 		"wms/controllers/users"
	Auth 		"wms/controllers/auth"
	Unit 		"wms/controllers/unit"
	Report 		"wms/controllers/reporting"
	"github.com/astaxie/beego"
)

func init() {

	//GOODS ROUTER
	beego.Router("/", &Login.LoginController{}, "get:Get")
	beego.Router("/login", &Login.LoginController{},"post:Login")
	beego.Router("/logout", &Login.LoginController{},"get:Logout")
	beego.Router("/dashboard", &Dashboard.MainController{})
	beego.Router("/incoming-goods",&Incoming.MainController{}, "get:All")
	beego.Router("/incoming-goods/add",&Incoming.MainController{}, "post:Add")
	beego.Router("/incoming-goods/detail/:id([0-9]+)", &Incoming.MainController{}, "get:Detail")
	beego.Router("/incoming-goods/detail-proccess/:id([0-9]+)", &Incoming.MainController{}, "get:DetailProccess")
	beego.Router("/incoming-goods/update/:id([0-9]+)", &Incoming.MainController{}, "post:Update")
	beego.Router("/incoming-goods/proccess/:id([0-9]+)", &Outcoming.OutController{}, "post:OutProccess")
	beego.Router("/incoming-goods/delete/:id([0-9]+)", &Incoming.MainController{}, "*:Delete")
	
	beego.Router("/outcoming-goods", &Outcoming.OutController{}, "get:All")
	beego.Router("/outcoming-goods/detail/:id([0-9]+)", &Outcoming.OutController{},"get:Detail")
	beego.Router("/outcoming-goods/update/:id([0-9]+)", &Outcoming.OutController{},"post:Update")
	// beego.Router("/outcoming-goods/reflect", &Outcoming.OutController{},"get:Reflect")
	// beego.Router("/outcoming-goods/print", &Outcoming.OutController{},"get:Print")

	//UNITS ROUTER
	beego.Router("/units", &Unit.UnitController{}, "get:All")
	beego.Router("/units/add", &Unit.UnitController{}, "post:Add")
	beego.Router("/units/detail/:id([0-9]+)", &Unit.UnitController{}, "get:Detail")
	beego.Router("/units/update/:id([0-9]+)", &Unit.UnitController{}, "post:Update")
	beego.Router("/units/delete/:id([0-9]+)", &Unit.UnitController{}, "*:Delete")

	//USERS ROUTER
	beego.Router("/users", &Users.UserController{}, "get:All")
	beego.Router("/users/delete/:id([0-9]+)", &Users.UserController{}, "*:Delete")
	beego.Router("/users/add", &Users.UserController{}, "post:Add")
	
	//AUTH ROUTER
	beego.Router("/auth", &Auth.MainController{}, "post:Login")
	
	beego.Router("/reporting", &Report.ReportingController{}, "get:All")
	beego.Router("/reporting/export", &Report.ReportingController{}, "get:Export")

}


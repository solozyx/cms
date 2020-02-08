package routers

import (
	"github.com/astaxie/beego"

	"github.com/solozyx/cms/controllers"
)

func init() {
	//login
	beego.Router("/login", &controllers.LoginController{}, "*:Index")

	//home
	beego.Router("/", &controllers.HomeController{}, "Get:Index")

	//menu
	beego.Router("/menu", &controllers.MenuController{}, "Get:Index")
	beego.Router("/menu/list", &controllers.MenuController{}, "*:List")
	beego.Router("/menu/edit", &controllers.MenuController{}, "Get:Edit")
	beego.Router("/menu/editdo", &controllers.MenuController{}, "*:EditDo")
	beego.Router("/menu/add", &controllers.MenuController{}, "Get:Add")
	beego.Router("/menu/adddo", &controllers.MenuController{}, "*:AddDo")
	beego.Router("/menu/deletedo", &controllers.MenuController{}, "Get:DeleteDo")

	//user
	beego.Router("/user", &controllers.UserController{}, "Get:Index")
	beego.Router("/user/add", &controllers.UserController{}, "Get:Add")
	beego.Router("/user/adddo", &controllers.UserController{}, "*:AddDo")
	beego.Router("/user/edit", &controllers.UserController{}, "Get:Edit")
	beego.Router("/user/editdo", &controllers.UserController{}, "*:EditDo")
	beego.Router("/user/deletedo", &controllers.UserController{}, "Get:DeleteDo")
	beego.Router("/user/list", &controllers.UserController{}, "*:List")

	//format
	beego.Router("/format/edit", &controllers.FormatController{}, "Get:Edit")
	beego.Router("/format/examples", &controllers.FormatController{}, "Get:Examples")
	beego.Router("/format/editdo", &controllers.FormatController{}, "*:EditDo")

	//list
	beego.Router("/data/?:mid", &controllers.DataController{}, "Get:Index")
	beego.Router("/data/list/?:mid", &controllers.DataController{}, "*:List")
	beego.Router("/data/add/?:mid", &controllers.DataController{}, "Get:Add")
	beego.Router("/data/adddo/?:mid", &controllers.DataController{}, "*:AddDo")
	beego.Router("/data/edit/?:mid", &controllers.DataController{}, "Get:Edit")
	beego.Router("/data/editdo/?:mid", &controllers.DataController{}, "*:EditDo")
	beego.Router("/data/deletedo/?:mid", &controllers.DataController{}, "*:DeleteDo")
}

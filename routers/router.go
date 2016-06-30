package routers

import (
	"github.com/astaxie/beego"
	"ssologin/controllers"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/cookie", &controllers.CookieController{})
	ns := beego.NewNamespace(
		"v1",
		beego.NSRouter(
			"/login",
			&controllers.LoginController{}),
	)
	beego.AddNamespace(ns)
}

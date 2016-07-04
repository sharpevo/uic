package routers

import (
	"github.com/astaxie/beego"
	"ssologin/controllers"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/cookie", &controllers.CookieController{})
	beego.Router("/register", &controllers.RegisterController{})
	ns := beego.NewNamespace(
		"v1",
		beego.NSRouter(
			"/login",
			&controllers.LoginController{}),
		beego.NSRouter(
			"/logout",
			&controllers.LogoutController{}),
	)
	beego.AddNamespace(ns)
}

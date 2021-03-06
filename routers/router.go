package routers

import (
	"github.com/astaxie/beego"
	"uic/controllers"
)

func init() {
	beego.Router("/", &controllers.ProfileController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/forgot", &controllers.ForgotPasswdController{})
	beego.Router("/reset", &controllers.ResetPasswdController{})

	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/roles", &controllers.RoleController{})
	beego.Router("/apps", &controllers.AppController{})
	beego.Router("/userapp", &controllers.UserAppController{})
	beego.Router("/profile", &controllers.ProfileController{})
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

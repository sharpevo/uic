package main

import (
	"github.com/astaxie/beego"
	"sso-client/utils/rbac"
	"strings"
	"time"
	"uic/models"
	"uic/mongo"
	_ "uic/routers"
)

func FormatDate(date time.Time) (result string) {
	if date.IsZero() {
		return "-"
	}
	layout := "15:04 on Jan 02, 2006"
	return date.Format(layout)
}

func SpaceToDot(str string) (result string) {
	return strings.Replace(str, " ", ".", -1)
}

func HasApp(user models.User, appId string) bool {
	return user.Apps[appId]
}

func main() {
	err := mongo.Startup()
	if err != nil {
		beego.Error(err)
	}

	beego.AddFuncMap("Format", FormatDate)
	beego.AddFuncMap("SpaceToDot", SpaceToDot)
	beego.AddFuncMap("HasApp", HasApp)

	beego.AddFuncMap("hasRole", rbac.HasRole)
	rbac.AdminCheck(
		"/roles",
		"/apps",
		"/userapp",
	)
	if !beego.AppConfig.DefaultBool("signupenabled", false) {
		beego.Debug("Registration:", "disabled")
		rbac.AdminCheck(
			"/register",
		)
	}

	beego.Run()

	err = mongo.Shutdown()
	if err != nil {
		beego.Error(err)
	}
}

package main

import (
	"github.com/astaxie/beego"
	"sso-client/utils/rbac"
	"time"
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

func main() {
	err := mongo.Startup()
	if err != nil {
		beego.Error(err)
	}

	beego.AddFuncMap("Format", FormatDate)

	beego.AddFuncMap("hasRole", rbac.HasRole)
	rbac.AdminCheck(
		"/roles",
	)

	beego.Run()

	err = mongo.Shutdown()
	if err != nil {
		beego.Error(err)
	}
}

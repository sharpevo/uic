package main

import (
	"github.com/astaxie/beego"
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

	beego.Run()

	err = mongo.Shutdown()
	if err != nil {
		beego.Error(err)
	}
}

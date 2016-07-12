package main

import (
	"github.com/astaxie/beego"
	"uic/mongo"
	_ "uic/routers"
)

func main() {
	err := mongo.Startup()
	if err != nil {
		beego.Error(err)
	}

	beego.Run()

	err = mongo.Shutdown()
	if err != nil {
		beego.Error(err)
	}
}

package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"uic/models"
)

type AppController struct {
	BaseController
}

func (c *AppController) Post() {
	flash := beego.NewFlash()
	userId := c.GetString("userId")
	appName := strings.Replace(
		strings.ToLower(c.GetString("appName")),
		" ",
		"",
		-1)
	appName = strings.Replace(appName, ".", " ", -1)

	user := models.User{}
	if code, err := user.FindById(userId); err != nil {
		beego.Error("FindUserById:", err)
		if code == models.ERROR_NOT_FOUND {
			flash.Error("No such user.")
		} else {
			flash.Error("Database Error.")
		}
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(".Get"),
			302)
		return
	}

	if !user.Apps[appName] {
		beego.Debug("AddApp:", appName)
		err := user.AddApp(appName)
		if err != nil {
			beego.Error("AddApp:", err)
			flash.Error("Fail to add app.")
			flash.Store(&c.Controller)
			c.Redirect(
				c.URLFor("RoleController.Get"),
				302)
			return
		}

	} else {
		beego.Debug("RemoveApp:", appName)
		err := user.RemoveApp(appName)
		if err != nil {
			beego.Error("RemoveApp:", err)
			flash.Error("Fail to add app.")
			flash.Store(&c.Controller)
			c.Redirect(
				c.URLFor("RoleController.Get"),
				302)
			return
		}
	}

	flash.Notice("Success!")
	flash.Store(&c.Controller)
	c.Redirect(
		c.URLFor(
			"RoleController.Get"),
		302)
}

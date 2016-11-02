package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"uic/models"
)

type UserAppController struct {
	BaseController
}

func (c *UserAppController) Post() {
	userId := c.GetString("userId")
	appId := c.GetString("appId")
	if userId == "" || appId == "" {
		c.Data["json"] = "Invalid input"
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}

	user := models.User{}
	if code, err := user.FindById(userId); err != nil {
		beego.Error("FindUserById:", err)
		if code == models.ERROR_NOT_FOUND {
			c.Data["json"] = "No such user."
		} else {
			c.Data["json"] = "Database Error."
		}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	app := models.App{}
	if err := app.FindById(appId); err != nil {
		beego.Error("FindAppById:", err)
		c.Data["json"] = "No such app."
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}
	beego.Debug("Before change app:", user.Apps)
	if user.Apps[appId] {
		user.RemoveApp(appId)
		beego.Debug("After Remove App:", user.Apps)
	} else {
		user.AddApp(appId)
		beego.Debug("After Add App:", user.Apps)
	}
	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = "Success."
	c.ServeJSON()
}

func (c *UserAppController) Post2() {
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
			c.URLFor("RoleController.Get"),
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

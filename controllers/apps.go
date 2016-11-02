package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"uic/models"
)

type AppController struct {
	BaseController
}

func (c *AppController) Get() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["Apps"] = models.GetAllApps()
	c.Layout = "layout.tpl"
	c.TplName = "app.tpl"
}

func (c *AppController) Post() {
	flash := beego.NewFlash()
	c.Layout = "layout.tpl"
	c.TplName = "app.tpl"

	appName := c.GetString("appName")
	appDomain := c.GetString("appDomain")
	appRemark := c.GetString("appRemark")

	if appName == "" || appDomain == "" {
		flash.Error("Invaild input.")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"appName", appName,
				"appDomain", appDomain),
			302)
		return
	}

	app := models.App{
		Name:    appName,
		Domain:  appDomain,
		Enabled: true,
		Remark:  appRemark,
	}

	if code, err := app.Create(); err != nil || code != 0 {
		if code == models.ERROR_DUPLICATE {
			beego.Error("Domain has been registered.")
			flash.Error("Domain has been registered.")
		} else {
			beego.Error("Unknown Error.")
			flash.Error("Fail to create domain.")
		}
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"appName", appName,
				"appDomain", appDomain),
			302)
		return
	}

	flash.Notice("Success!")
	flash.Store(&c.Controller)
	c.Redirect(
		c.URLFor(".Get"),
		302)
	return
}

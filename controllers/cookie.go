package controllers

import (
	"github.com/astaxie/beego"
)

type CookieController struct {
	BaseController
}

func (c *CookieController) Get() {
	jwt := c.GetString("jwt")
	returnTo := c.GetString("return_to")
	if returnTo == "" {
		returnTo = "www.igenetech.com"
	}
	c.Data["Token"] = jwt
	c.Data["ReturnTo"] = returnTo
	beego.Debug("jwt:", jwt)
	beego.Debug("return_to", returnTo)
	c.TplName = "cookie.tpl"
}

package controllers

import (
//"github.com/astaxie/beego"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "home.tpl"
}

package controllers

import (
	//"github.com/astaxie/beego"
	"uic/models"
)

type ProfileController struct {
	BaseController
}

func (c *ProfileController) Get() {
	user := models.User{}
	user.FindById(c.UserInfo.Id)
	c.Data["User"] = user
	c.TplName = "profile.tpl"
	c.Layout = "layout.tpl"
}

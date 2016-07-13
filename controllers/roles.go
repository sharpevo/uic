package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"uic/models"
)

type RoleController struct {
	BaseController
}

func (c *RoleController) Get() {
	//userInfo := c.GetUserInfo()

	if c.UserInfo.Id == "" {
		beego.Error("Wierd error.")
		c.Data["ReturnTo"] = "www.igenetech.com"
		c.TplName = "logout.tpl"
		c.Layout = "layout.tpl"
		return
	}
	if c.UserInfo.Email != "quwubin@gmail.com" {
		beego.Debug("IllegalUser:", c.UserInfo.Email)
		c.Data["ReturnTo"] = "www.igenetech.com"
		c.TplName = "logout.tpl"
		c.Layout = "layout.tpl"
		return
	}
	beego.ReadFromRequest(&c.Controller)
	c.Data["UserList"], _ = models.GetUsersSortByEmail()
	c.TplName = "roles.tpl"
	c.Layout = "layout.tpl"
}

func (c *RoleController) Post() {
	//userInfo := c.GetUserInfo()

	if c.UserInfo.Id == "" {
		beego.Error("Wierd error.")
		c.Data["ReturnTo"] = "www.igenetech.com"
		c.TplName = "logout.tpl"
		c.Layout = "layout.tpl"
		return
	}
	if c.UserInfo.Email != "quwubin@gmail.com" {
		beego.Debug("IllegalUser:", c.UserInfo.Email)
		c.Data["ReturnTo"] = "www.igenetech.com"
		c.TplName = "logout.tpl"
		c.Layout = "layout.tpl"
		return
	}

	flash := beego.NewFlash()
	userId := c.GetString("userId")
	roleName := strings.ToLower(
		c.GetString("roleName"))

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

	if !user.Roles[roleName] {
		beego.Debug("AddRole:", user.Roles)
		err := user.AddRole(roleName)
		if err != nil {
			beego.Error("AddRole:", err)
			flash.Error("Fail to add role.")
			flash.Store(&c.Controller)
			c.Redirect(
				c.URLFor(".Get"),
				302)
			return
		}
	} else {
		beego.Debug("RemoveRole:", user.Roles)
		err := user.RemoveRole(roleName)
		if err != nil {
			beego.Error("RemoveRole:", err)
			flash.Error("Fail to add role.")
			flash.Store(&c.Controller)
			c.Redirect(
				c.URLFor(".Get"),
				302)
			return
		}
	}

	flash.Notice("Success!")
	flash.Store(&c.Controller)
	c.Redirect(
		c.URLFor(
			".Get"),
		302)
}

func (c *RoleController) Delete() {
	flash := beego.NewFlash()
	userId := c.GetString("userId")
	roleName := c.GetString("roleName")

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
	beego.Debug("RemoveRole:", user.Roles)
	err := user.RemoveRole(roleName)
	if err != nil {
		beego.Error("RemoveRole:", err)
		flash.Error("Fail to add role.")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(".Get"),
			302)
		return
	}
	beego.Debug("RemoveRole:", user.Roles)

	flash.Notice("Success!")
	flash.Store(&c.Controller)
	c.Redirect(
		c.URLFor(
			".Get"),
		302)

}

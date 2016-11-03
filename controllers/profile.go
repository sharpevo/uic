package controllers

import (
	"github.com/astaxie/beego"
	"uic/models"
)

type ProfileController struct {
	BaseController
}

func (c *ProfileController) Get() {
	beego.ReadFromRequest(&c.Controller)
	user := models.User{}
	user.FindById(c.UserInfo.Id)
	c.Data["User"] = user
	c.Data["AppList"], _ = models.GetEnabledApps()
	c.TplName = "profile.tpl"
	c.Layout = "layout.tpl"
}

func (c *ProfileController) Post() {
	user := models.User{}
	user.FindById(c.UserInfo.Id)
	userName := c.GetString("name")
	userEmail := c.GetString("email")
	userPasswd := c.GetString("password")
	userPasswdConfirm := c.GetString("password_confirmation")
	if user.Name != userName ||
		user.Email != userEmail ||
		user.Password != "" {

		flash := beego.NewFlash()

		if user.Email != userEmail {
			existed, err := models.CheckEmail(userEmail)
			beego.Debug("CheckEmail:", userEmail, existed)
			if existed || err != nil {
				flash.Error("Failed to update user info.")
				flash.Store(&c.Controller)
				c.Redirect(
					c.URLFor(".Get"),
					302)
				return
			}
		}

		user.Name = userName
		user.Email = userEmail

		if userPasswd != "" &&
			userPasswd == userPasswdConfirm {
			user.Password = user.EncryptPassword(userPasswd)
		}
		_, err := user.Update()
		if err != nil {
			flash.Error("Fail to update user info.")
		} else {
			flash.Notice("Success.")
			beego.Debug("UpdateUser:", user)
		}
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(".Get"),
			302)
		return
	}
}

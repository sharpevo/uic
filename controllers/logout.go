package controllers

import (
	"github.com/astaxie/beego"
	"uic/models"
)

type LogoutController struct {
	BaseController
}

func (c *LogoutController) Get() {

	user := models.User{}
	if code, err := user.FindById(c.UserInfo.Id); err != nil {
		beego.Error("FindUserById:", err)
		if code == models.ERROR_NOT_FOUND {
			beego.Error("No such user.")
		} else {
			beego.Error("Database Error.")
		}
		return
	}

	userAgent := c.ParseUserAgent()

	beego.Debug("UserTokens:", user.Tokens)
	ok := user.RemoveToken(userAgent)
	if !ok {
		beego.Error("Logout Fail")
		return
	}
	beego.Debug("UserLogout:", user.Email)
	beego.Debug("UserTokens:", user.Tokens)
	returnTo := "www.igenetech.com"
	c.Data["ReturnTo"] = returnTo
	c.TplName = "logout.tpl"
	return
}

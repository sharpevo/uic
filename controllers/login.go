package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"net/http"
	"uic/models"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["Email"] = c.GetString("email")
	c.Layout = "layout.tpl"
	c.TplName = "login.tpl"
}

func (c *LoginController) Post() {
	form := models.LoginForm{}
	flash := beego.NewFlash()
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseLoginForm:", err)
		flash.Error("Invalid Input")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"email",
				form.Email),
			302)
		return
	}
	beego.Debug("ParseLoginForm:", &form)

	if err := c.VerifyForm(&form); err != nil {
		beego.Debug("ValidLoginForm:", err)
		flash.Error("Invalid Input")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"email",
				form.Email),
			302)
		return
	}

	user := models.User{}
	if code, err := user.FindByEmail(form.Email); err != nil {
		beego.Error("FindUserByEmail:", err)
		if code == models.ERROR_NOT_FOUND {
			flash.Error("Invalid Input")
		} else {
			flash.Error("Fail to login")
		}
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"email",
				form.Email),
			302)
		return
	}
	beego.Debug("UserInfo:", &user)

	if ok, err := user.CheckPass(form.Password); err != nil {
		beego.Error("CheckUserPassword:", err)
		flash.Error("Fail to login")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"email",
				form.Email),
			302)
		return
	} else if !ok {
		flash.Error("Invalid Input")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"email",
				form.Email),
			302)
		return
	}

	// Parse duration
	remember := c.GetString("remember")
	duration := int64(30)
	if remember == "on" {
		duration = int64(7200)
	}

	tokenString, err := c.GenerateToken(user.Id.Hex(), duration)
	if err != nil {
		beego.Error("GenerateToken:", err)
		flash.Error("Fail to login")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"email",
				form.Email),
			302)
		return
	}

	beego.Debug("SignedToken:", tokenString)

	if err != nil {
		beego.Error("SignedString:", err)
		flash.Error("Invalid Input")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"email",
				form.Email),
			302)
		return
	}

	// TODO: Parse useragent to the browser/app/wechat
	userAgent := c.ParseUserAgent()

	beego.Debug("UserAgent:", userAgent)
	err = user.AddToken(userAgent, tokenString)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("AddToken:", user)

	// Process cookies
	returnTo := c.GetString("return_to")
	if returnTo == "" {
		returnTo = "accounts.igenetech.com"
	}
	c.Data["Token"] = tokenString
	c.Data["ReturnTo"] = returnTo
	beego.Debug("return_to", returnTo)

	c.SetCookie(".igenetech.com", tokenString, int64(duration))

	c.TplName = "cookie.tpl"

}

func (c *LoginController) PostAPI() {

	form := models.LoginForm{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &form)
	beego.Debug("ParseLoginRequest:", form)
	if err != nil {
		beego.Error("ParseLoginRequest:", err)
		c.AuthFail()
		return
	}

	user := models.User{}
	if code, err := user.FindByEmail(form.Email); err != nil {
		beego.Error("FindUserByEmail:", err)
		if code == models.ERROR_NOT_FOUND {
			c.Data["json"] = models.ErrorJSON("No such user")
		} else {
			c.Data["json"] = models.ErrorJSON("Database Error")
		}
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
		c.ServeJSON()
		return
	}
	beego.Debug("UserInfo:", &user)

	if ok, err := user.CheckPass(form.Password); err != nil {
		beego.Error("CheckUserPassword:", err)
		c.Data["json"] = models.ErrorJSON("Database Error")
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
		c.ServeJSON()
		return
	} else if !ok {
		c.Data["json"] = models.ErrorJSON("Wrong password")
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
		c.ServeJSON()
		return
	}

	tokenString, err := c.GenerateToken(user.Id.Hex(), 30)
	if err != nil {
		beego.Error("GenerateToken:", err)
		c.Data["json"] = ErrToken
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
		c.ServeJSON()
		return
	}

	beego.Debug("SignedToken:", tokenString)

	if err != nil {
		beego.Error("SignedString:", err)
		c.Data["json"] = ErrInput
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
		c.ServeJSON()
		return
	}

	// TODO: Parse useragent to the browser/app/wechat
	userAgent := c.ParseUserAgent()

	beego.Debug("UserAgent:", userAgent)
	err = user.AddToken(userAgent, tokenString)
	if err != nil {
		beego.Error(err)
	}
	beego.Debug("AddToken:", user)

	c.Data["json"] = models.NormalJSON(tokenString)
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.ServeJSON()
}

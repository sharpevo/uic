package controllers

import (
	"github.com/astaxie/beego"
	"uic/models"
)

type RegisterController struct {
	BaseController
}

func (c *RegisterController) Get() {
	c.Layout = "layout.tpl"
	c.TplName = "register.tpl"
	c.Data["Name"] = c.GetString("name")
	c.Data["Email"] = c.GetString("email")
	beego.ReadFromRequest(&c.Controller)
}

func (c *RegisterController) Post() {

	c.Layout = "layout.tpl"
	c.TplName = "register.tpl"

	flash := beego.NewFlash()

	form := models.RegisterForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("ParseRegsiterForm:", err)
		flash.Error("Invalid Input.")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"name", form.Name,
				"email", form.Email),
			302)
		return
	}

	if ok := cpt.VerifyReq(c.Ctx.Request); !ok {
		flash.Error("Invalid Captcha.")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"name", form.Name,
				"email", form.Email),
			302)
		return
	}

	if form.Password != form.ConfirmPassword {
		beego.Debug("PasswordNotMatch:", form.Password, form.ConfirmPassword)
		flash.Error("Passwords don't match.")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"name", form.Name,
				"email", form.Email),
			302)
		return
	}
	beego.Debug("ParseRegsiterForm:", &form)

	if err := c.VerifyForm(&form); err != nil {
		beego.Debug("ValidRegsiterForm:", err)
		flash.Error("Invalid Input.")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"name", form.Name,
				"email", form.Email),
			302)
		return
	}

	user, err := models.ParseUser(&form)

	if err != nil {
		beego.Error("ParseUser:", err)
		flash.Error("Fail to create account.")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"name", form.Name,
				"email", form.Email),
			302)
		return
	}
	beego.Debug("ParseUser:", user)

	if code, err := user.Create(); err != nil || code != 0 {
		if code == models.ERROR_DUPLICATE {
			beego.Error("Email has been registered.")
			flash.Error("Email has been registered.")
		} else {
			beego.Error("Unknown Error.")
			flash.Error("Fail to create account.")
		}
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"name", form.Name,
				"email", form.Email),
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

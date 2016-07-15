package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/dgrijalva/jwt-go"
	"html/template"
	"io/ioutil"
	"strings"
	"time"
	"uic/models"
	"uic/utils/mail"
)

const EXP = 1 // 30 min

type ForgotPasswdController struct {
	BaseController
}

func (c *ForgotPasswdController) Get() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.Data["Email"] = c.GetString("email")
	c.Layout = "layout.tpl"
	c.TplName = "forgotPasswd.tpl"
}

func (c *ForgotPasswdController) Post() {
	flash := beego.NewFlash()
	email := strings.ToLower(c.GetString("email"))

	if ok := cpt.VerifyReq(c.Ctx.Request); !ok {
		flash.Error("Invalid Captcha.")
		flash.Store(&c.Controller)
		c.Redirect(
			c.URLFor(
				".Get",
				"email", email),
			302)
		return
	}

	valid := validation.Validation{}
	valid.Required(email, "email")
	if valid.HasErrors() {
		beego.Error("CheckEmail:", email)
		flash.Error("Invalid email.")
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor(".Get"), 302)
		return
	}
	beego.Debug("ValidEmail:", email)

	user := models.User{}
	_, err := user.FindByEmail(email)
	if err != nil {
		beego.Error("GetUser:", err)
		flash.Error("Fail to send request")
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor(".Get"), 302)
		return
	}
	beego.Debug("GetUser:", user)

	if time.Now().Before(
		user.DateResetPasswd.Add(
			time.Duration(24) * time.Hour)) {
		beego.Error("CheckFreq:", user.DateResetPasswd)
		flash.Error("Password Resetting is allowed once per day.")
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor(".Get"), 302)
		return
	}

	token, err := c.GenerateToken(user.Id.Hex(), int64(EXP))
	if err != nil {
		beego.Error("GenerateResetToken:", err)
		flash.Error("Fail to send request.")
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor(".Get"), 302)
		return
	}
	beego.Debug("GenerateResetToken:", token)

	resetLink := c.BuildLink(
		c.URLFor(
			"ResetPasswdController.Get",
			"token",
			token))

	err = mail.SendResetMail(
		email,
		resetLink)
	if err != nil {
		beego.Error("SendMail:", err)
		flash.Error("Fail to send request.")
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor(".Get"), 302)
		return
	}
	beego.Debug("SendMailWithLink:", email, resetLink)

	user.DateResetPasswd = time.Now()
	user.Update()

	flash.Notice("Success to send reset mail")
	flash.Store(&c.Controller)
	c.Redirect(c.URLFor("LoginController.Get"), 302)
	return
}

type ResetPasswdController struct {
	BaseController
}

func (c *ResetPasswdController) Get() {
	beego.ReadFromRequest(&c.Controller)
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())

	token := c.GetString("token")
	c.Data["token"] = token

	c.Layout = "layout.tpl"
	c.TplName = "resetPasswd.tpl"
}

func (c *ResetPasswdController) Post() {
	flash := beego.NewFlash()
	tokenString := c.GetString("token")
	password := c.GetString("password")
	passwordConfirmation := c.GetString("password_confirmation")

	pubBytes, err := ioutil.ReadFile("keys/ip.rsa.pub")
	if err != nil {
		beego.Error("ReadPubKey:", err)
		flash.Error("Fail to update password")
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor(".Get"), 302)
		return
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubBytes)
	beego.Debug("ReadPubKey:", pubKey)

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return pubKey, nil
		},
	)
	if err != nil {
		beego.Error("ParseToken:", err)
		flash.Error("Fail to update password")
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor(".Get"), 302)
		return
	}

	if !token.Valid {
		beego.Error("InvalidToken:", token)
		flash.Error("Fail to update password")
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor(".Get"), 302)
		return
	}

	valid := validation.Validation{}
	valid.Required(password, "MinSize(6);MaxSize(15)")
	if valid.HasErrors() ||
		password != passwordConfirmation {
		beego.Error("CheckPassword:", password)
		flash.Error("Invalid password.")
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor(".Get"), 302)
		return
	}
	beego.Debug("ValidPassword:", password)

	claims := token.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	beego.Debug("ParseUserId:", userId)

	user := models.User{}
	if _, err := user.FindById(userId); err != nil {
		beego.Error("FindUser:", err)
		flash.Error("Fail to update password")
		flash.Store(&c.Controller)
		c.Redirect(c.URLFor(".Get"), 302)
		return

	}
	user.Password = user.EncryptPassword(password)
	user.Update()

	beego.Debug("UpdatePassword:", user.Id.Hex())
	flash.Notice("Password is updated successfully.")
	flash.Store(&c.Controller)
	c.Redirect(c.URLFor("LoginController.Get"), 302)
	return
}

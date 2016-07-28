package controllers

import (
	"github.com/astaxie/beego"
	"net/http"
	"strings"
	"time"
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

	c.SetCookie(jwt)

	domainList := strings.Split(
		beego.AppConfig.DefaultString("appdomains", ""),
		",")
	beego.Error("Domains:", beego.AppConfig.DefaultString("appdomains", ""))
	beego.Debug("ParseDomains:", domainList)
	c.Data["Domains"] = domainList
	c.TplName = "cookie.tpl"
}

func (c *CookieController) SetCookie(value string) {
	if value == "" {
		beego.Debug("SetCookie:", "Nothing to set.")
		return
	}
	expiration := time.Now().Add(30 * time.Minute)
	cookie := http.Cookie{Name: "token", Value: value, HttpOnly: true, Domain: ".igenetech.com", Expires: expiration}
	beego.Debug("SetCookie:", cookie)
	http.SetCookie(c.Ctx.ResponseWriter, &cookie)
}

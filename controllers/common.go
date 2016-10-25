package controllers

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
	"github.com/astaxie/beego/validation"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"net/http"
	"reflect"
	"sso-client/controllers/common"
	"time"
)

type ControllerError struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	ErrInput   = &ControllerError{400, 10001, "Invalid Inputs"}
	ErrExpired = &ControllerError{400, 10012, "Token is expired"}
	ErrToken   = &ControllerError{400, 10012, "Fail to get valid token"}
)

type BaseController struct {
	common.BaseController
}

type NestPreparer interface {
	NestPrepare()
}

var PriKey *rsa.PrivateKey

var cpt *captcha.Captcha

func init() {
	// use beego cache system store the captcha data
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
}

func (c *BaseController) Prepare() {
	c.GetUserInfo()
	c.Data["UICDomain"] = fmt.Sprintf(
		"%s:%s",
		beego.AppConfig.DefaultString(
			"uicdomain",
			"accounts.igenetech.com"),
		beego.AppConfig.DefaultString(
			"httpport",
			"8080"))

	c.Data["SignUpEnabled"] = beego.AppConfig.DefaultBool("signupenabled", false)

	priBytes, err := ioutil.ReadFile("keys/ip.rsa")
	if err != nil {
		beego.Error("ReadPrivateBytes:", err)
		return
	}
	PriKey, err = jwt.ParseRSAPrivateKeyFromPEM(priBytes)
	beego.Debug("ReadPrivateKey")

	if app, ok := c.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (c *BaseController) ValidateForm(form interface{}) bool {
	c.ParseForm(form)
	formName := reflect.ValueOf(form).Elem().Type().Name()
	c.Data[formName] = form
	vld := validation.Validation{}
	pass, _ := vld.Valid(form)
	if !pass {
		c.Data[formName+"Error"] = &vld.ErrorsMap
	}
	return vld.HasErrors()
}

func (c *BaseController) VerifyForm(form interface{}) (err error) {
	valid := validation.Validation{}
	ok, err := valid.Valid(form)
	if err != nil {
		return err
	}
	if !ok {
		str := ""
		for _, err := range valid.Errors {
			str += err.Key + ":" + err.Message + ";"
		}
		return errors.New(str)
	}
	return nil
}

func (c *BaseController) ParseUserAgent() (userAgent string) {
	userAgent = c.Ctx.Input.UserAgent()
	if userAgent == "wechat" {
		userAgent = "wechat"
	} else {
		userAgent = "browser"
	}

	return userAgent
}

func (c *BaseController) GenerateToken(userId string, exp int64) (tokenString string, err error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		jwt.MapClaims{
			"sub": userId,
			"exp": time.Now().Add(time.Minute * time.Duration(exp)).Unix(),
		})
	tokenString, err = token.SignedString(PriKey)
	return tokenString, err
}

func (c *BaseController) AuthFail() {
	http.Error(c.Ctx.ResponseWriter, "Not logged in", http.StatusUnauthorized)
}

func (c *BaseController) BuildLink(path string) string {
	return fmt.Sprintf("%v:%v%v", c.Ctx.Input.Site(), c.Ctx.Input.Port(), path)
}

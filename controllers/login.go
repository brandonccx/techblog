package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	isExit := c.Input().Get("exit") == "1"
	if isExit {
		beego.Emergency("======================== user is exit")
		c.Ctx.SetCookie("account", "", -1, "/")
		c.Ctx.SetCookie("password", "", -1, "/")
		c.Redirect("/", 302)
		return
	}
	c.TplName = "login.html"
	c.Data["Title"] = "Login"
}

func (c *LoginController) Post() {
	inputs := c.Input()
	account := inputs.Get("Account")
	password := inputs.Get("Password")
	autoLogin := inputs.Get("AutoLogin") == "on"

	if beego.AppConfig.String("account") == account &&
		beego.AppConfig.String("password") == password {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		c.Ctx.SetCookie("account", account, maxAge, "/")
		c.Ctx.SetCookie("password", password, maxAge, "/")
		c.Redirect("/", 301)
		return
	} else {
		c.Redirect("/login", 301)
		return
	}
}

func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("account")
	if err != nil {
		return false
	}
	account := ck.Value
	ck, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := ck.Value

	return beego.AppConfig.String("account") == account &&
		beego.AppConfig.String("password") == password
}

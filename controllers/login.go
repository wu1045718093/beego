package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "index.html"
}

func (c *LoginController) Post() {
	email := c.GetString("email")
	password := c.GetString("password")
	c.Ctx.WriteString("登陆成功" + email + password)
	c.Ctx.Redirect(302, "\\")
}

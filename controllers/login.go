package controllers

import (
	"login/model"

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

	ok := model.PrepareQuerry(email, password)
	if ok == true {
		c.Ctx.Redirect(302, "/")
	} else {
		c.Ctx.WriteString("密码错误")
	}
}

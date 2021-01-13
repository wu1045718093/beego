package controllers

import (
	"login/model"

	beego "github.com/beego/beego/v2/server/web"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	c.TplName = "register.html"
}

func (c *RegisterController) Post() {
	email := c.GetString("email")
	password := c.GetString("password")

	result := model.PrepareInsert(email, password)
	if result == 0 {
		c.Ctx.Redirect(302, "/login.html")
	} else if result == -1 {
		c.Ctx.WriteString("用户名已存在")
	} else {
		c.Ctx.WriteString("未知错误")
	}
}

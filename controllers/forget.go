package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ForgotController struct {
	beego.Controller
}

func (c *ForgotController) Get() {
	c.TplName = "forgot.html"
}

func (c *ForgotController) Post() {
	c.Ctx.WriteString("重置密码成功")
}

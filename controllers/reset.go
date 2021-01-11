package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ResetController struct {
	beego.Controller
}

func (c *ResetController) Get() {
	c.TplName = "reset.html"
}

func (c *ResetController) Post() {
	c.Ctx.WriteString("重置成功")
}

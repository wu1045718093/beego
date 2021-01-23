package controllers

import (
	"fmt"
	"login/model"

	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	email := c.GetString("email")
	password := c.GetString("password")

	collect, ok := model.RedisGet(email, password)
	if ok == true {
		fmt.Println("缓存命中")
		if collect == true {
			c.Ctx.Redirect(302, "/")
		} else {
			c.Ctx.WriteString("登陆失败")
		}
	} else {
		//redis缓存未命中
		fmt.Println("缓存未命中")
		ok = model.PrepareQuerryLogin(email, password)
		if ok == true {
			model.RedisSet(email, password)
			c.Ctx.Redirect(302, "/")
		} else {
			c.Ctx.WriteString("登陆失败")
		}
	}

}

package main

import (
	_ "login/model"
	_ "login/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

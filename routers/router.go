package routers

import (
	"login/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/login", &controllers.LoginController{}, "post:Post")

	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/register", &controllers.RegisterController{}, "post:Post")
}

package routers

import (
	"login/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/index.html", &controllers.LoginController{})
	beego.Router("/index.html", &controllers.LoginController{}, "post:Post")

	beego.Router("/register.html", &controllers.RegisterController{})
	beego.Router("/register.html", &controllers.RegisterController{}, "post:Post")

	beego.Router("/reset.html", &controllers.ResetController{})
	beego.Router("/reset.html", &controllers.ResetController{}, "post:Post")

	beego.Router("/forgot.html", &controllers.ForgotController{})
	beego.Router("/forgot.html", &controllers.ForgotController{}, "post:Post")
}

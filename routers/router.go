package routers

import (
	"github.com/pravj/beehub/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/logout", &controllers.MainController{}, "get:Logout")
    beego.Router("/callback", &controllers.OauthController{}, "get:ParseCode")
}

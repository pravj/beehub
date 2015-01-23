package routers

import (
	"github.com/pravj/beehub/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.MainController{}, "get:Login")
    beego.Router("/callback", &controllers.OauthController{}, "get:ParseCode")
}

package routers

import (
	"github.com/pravj/BeeHub/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}

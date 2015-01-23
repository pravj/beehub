package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "pravj.github.io"
	this.Data["Email"] = "hackpravj@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *MainController) Login() {
	this.Data["Website"] = "pravj.github.io"
	this.Data["Email"] = "hackpravj@gmail.com"
	this.Data["ClientID"] = beego.AppConfig.String("client_id")
	this.TplNames = "index.tpl"
}

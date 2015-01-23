package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["ClientID"] = beego.AppConfig.String("client_id")

        v := this.GetSession("BeeHub")
        if v == nil {
          this.Data["status"] = "logged-out"
        } else {
          this.Data["status"] = "logged-in"
        }

	this.TplNames = "index.tpl"
}

func (this *MainController) Logout() {
  this.DelSession("BeeHub")
  this.Redirect("/", 302)
  return
}

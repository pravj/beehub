package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["ClientID"] = beego.AppConfig.String("client_id")

        v := this.GetSession("beehub")
        if v == nil {
          this.TplNames = "index.tpl"
        } else {
          w, _ := v.(map[string]string)
          this.Data["Name"] = w["name"]
          this.Data["Avatar"] = w["avatar"]
          this.TplNames = "user.tpl"
        }
}

func (this *MainController) Logout() {
  this.DelSession("beehub")
  this.Redirect("/", 302)
  return
}

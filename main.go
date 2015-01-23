package main

import (
	"github.com/astaxie/beego"
	_ "github.com/pravj/beehub/routers"
	_ "github.com/pravj/beehub/models"
	"github.com/astaxie/beego/orm"
        _ "github.com/go-sql-driver/mysql"
)

func main() {
        beego.SessionOn = true
	beego.Run()
}

func init() {
    dbPath := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@/" + beego.AppConfig.String("mysqldb")

    orm.RegisterDriver("mysql", orm.DR_MySQL)
    orm.RegisterDataBase("default", "mysql", dbPath)

    err := orm.RunSyncdb("default", false, true)
    if err != nil {
        beego.Info(err)
    }
}

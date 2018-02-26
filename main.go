package main

import (
	"script/ipsecMonitor/server/service"
	_ "script/ipsecMonitor/server/web/models"
	_ "script/ipsecMonitor/server/web/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "ipsecMonitor.db")
	orm.RunSyncdb("default", false, true)
	beego.SetLogFuncCall(true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	o := orm.NewOrm()
	o.Using("default")
	go beego.Run()

	monitor := service.Monitor{Dir: "/var/run/pluto"}
	go monitor.Run()

	select {}
}

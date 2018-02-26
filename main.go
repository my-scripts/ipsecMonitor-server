package main

import (
	"script/ipsecMonitor/server/service"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "ipsecMonitor.db")
	orm.RunSyncdb("default", false, true)
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

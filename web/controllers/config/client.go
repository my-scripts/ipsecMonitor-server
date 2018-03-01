package config

import (
	"script/ipsecMonitor/server/web/controllers/base"
	"script/ipsecMonitor/server/web/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type IpsecClientController struct {
	base.BaseController
}

func (this *IpsecClientController) Get() {
	this.Layout = "layout.html"
	this.TplName = "config/client/clients.html"

	var clients []models.ClientConf

	o := orm.NewOrm()
	o.QueryTable("ClientConf").All(&clients)

	this.Data["Objects"] = clients
	this.Data["ConfigPage"] = true
}

func (this *IpsecClientController) Add() {
	this.Layout = "layout.html"
	this.TplName = "config/client/client_form.html"
	client := models.ClientConf{}
	this.Data["ConfigPage"] = true
	this.Data["Object"] = client
}

func (this *IpsecClientController) Edit() {
	this.Layout = "layout.html"
	this.TplName = "config/client/client_form.html"

	id, _ := this.GetInt("id")
	client := models.ClientConf{Id: id}
	succ := client.GetClient()
	if !succ {
		this.SetError(base.ERROR_DB_READ)
		return
	}
	this.Data["Object"] = client
}

func (this *IpsecClientController) Post() {
	this.Layout = "layout.html"
	this.TplName = "config/client/clients.html"

	id, _ := this.GetInt("id")
	client := models.ClientConf{Id: id}
	if err := this.ParseForm(&client); err != nil {
		beego.Warning("failed parse form , %s", err)
	}

	succ := client.SetClient()
	if !succ {
		this.SetError(base.ERROR_DB_UPDATE)
		return
	}

	this.Ctx.Redirect(302, "/config/client")
}

func (this *IpsecClientController) Delete() {
	param := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(param)
	if err != nil {
		beego.Warn(err)
		return
	}

	o := orm.NewOrm()
	o.Delete(&models.ClientConf{Id: id})

	this.Redirect("/config/client/", 302)
}

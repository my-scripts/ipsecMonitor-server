package server

import (
	"os/exec"
	sysbase "script/ipsecMonitor/server/base"
	"script/ipsecMonitor/server/web/controllers/base"

	"github.com/astaxie/beego"
)

type IpsecServerController struct {
	base.BaseController
}

type jsonresult struct {
	Result bool
}

func (this *IpsecServerController) Get() {
	this.Layout = "layout.html"
	this.TplName = "status/server/server.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Script"] = "status/server/server_js.html"
	var services []*sysbase.IpsecService
	services = append(services, sysbase.GetIpsecService())
	this.Data["Services"] = services
}

func (this *IpsecServerController) Start() {
	cmd := exec.Command("/etc/init.d/ipsec", "start")
	err := cmd.Run()
	if err != nil {
		beego.Warning("start ipsec faild :", err)
	}

	this.Data["json"] = jsonresult{Result: err == nil}
	this.ServeJSON()
}

func (this *IpsecServerController) Stop() {
	cmd := exec.Command("/etc/init.d/ipsec", "stop")
	err := cmd.Run()
	if err != nil {
		beego.Warning("stop ipsec faild :", err)
	}

	this.Data["json"] = jsonresult{Result: err == nil}
	this.ServeJSON()
}

func (this *IpsecServerController) Restart() {
	cmd := exec.Command("/etc/init.d/ipsec", "restart")
	err := cmd.Run()
	if err != nil {
		beego.Warning("restart ipsec faild :", err)
	}

	this.Data["json"] = jsonresult{Result: err == nil}
	this.ServeJSON()
}

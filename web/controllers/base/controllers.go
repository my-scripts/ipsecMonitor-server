package base

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (this *BaseController) SetError(err string) {
	this.Data["HasError"] = true
	this.Data["Error"] = err
}

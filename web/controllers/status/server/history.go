package server

import (
	"script/ipsecMonitor/server/web/controllers/base"
	"script/ipsecMonitor/server/web/models"
	"strconv"
	"time"
)

type IpsecServerHisController struct {
	base.BaseController
}

func (this *IpsecServerHisController) Get() {
	this.Layout = "layout.html"
	this.TplName = "status/server/history.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Script"] = "section/list_js.html"

	param := this.Ctx.Input.Param(":page")
	page, err := strconv.Atoi(param)
	if err != nil {
		return
	}
	this.Data["CurrentPage"] = page

	this.Data["Url"] = "/status/server/history/"
	this.Data["DataCount"] = models.GetHistoryDataCount()
	data := models.GetHistoryData(page)
	for index, item := range data {
		data[index].Time = time.Unix(item.Stamp, 0)
	}
	this.Data["Objects"] = data
}

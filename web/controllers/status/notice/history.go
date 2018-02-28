package notice

import (
	"script/ipsecMonitor/server/web/controllers/base"
	"script/ipsecMonitor/server/web/models"
	"strconv"
	"time"
)

type NoticeHisController struct {
	base.BaseController
}

func (this *NoticeHisController) Get() {
	this.Layout = "layout.html"
	this.TplName = "status/notice/history.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Script"] = "section/list_js.html"

	param := this.Ctx.Input.Param(":page")
	page, err := strconv.Atoi(param)
	if err != nil {
		return
	}
	this.Data["CurrentPage"] = page

	this.Data["Url"] = "/status/notice/"
	this.Data["DataCount"] = models.GetNoticeHistoryDataCount()
	data := models.GetNoticeHistoryData(page)
	for index, item := range data {
		data[index].Time = time.Unix(item.Stamp, 0)
	}
	this.Data["Objects"] = data
}

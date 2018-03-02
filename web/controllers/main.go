package controllers

import (
	"fmt"
	"script/ipsecMonitor/server/web/controllers/base"
	"script/ipsecMonitor/server/web/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
)

type MainController struct {
	base.BaseController
}

func (this *MainController) convertBoottime(boottime uint64) string {
	total := uint64(time.Now().Unix()) - boottime
	days := total / (3600 * 24)
	left := total - days*3600*24
	hours := left / 3600
	left = left - hours*3600
	minutes := left / 60
	seconds := left - minutes*60

	sb := ""
	if days != 0 {
		sb += fmt.Sprintf("%d天 ", days)
	}

	if hours != 0 {
		sb += fmt.Sprintf("%d小时 ", hours)
	}

	if minutes != 0 {
		sb += fmt.Sprintf("%d分钟 ", minutes)
	}

	if seconds != 0 {
		sb += fmt.Sprintf("%d秒", seconds)
	}
	return sb
}

func (this *MainController) getState(state int) string {
	switch state {
	case models.IPSEC_SERVER_START:
		return "启动"
	case models.IPSEC_SERVER_ONLINE:
		return "工作"
	case models.IPSEC_SERVER_STOP:
		return "停止"
	}
	return ""
}

func (this *MainController) Get() {
	this.Layout = "layout.html"
	this.TplName = "index.html"
	// this.LayoutSections = make(map[string]string)
	// this.LayoutSections["HtmlHead"] = "index_head.html"
	// this.LayoutSections["Script"] = "index_js.html"

	this.Data["MainPage"] = true

	info, err := host.Info()
	if err == nil {
		this.Data["Hostname"] = info.Hostname
		this.Data["Boottime"] = this.convertBoottime(info.BootTime)
	} else {
		beego.Warn("failed to get host info", err)
	}

	this.Data["Now"] = time.Now().Format("2006-01-02 15:04:05")

	ds, err := disk.Usage("/")
	if err == nil {
		this.Data["DiskPercent"] = ds.UsedPercent
	}

	o := orm.NewOrm()
	var records []models.IpsecServerHistory

	_, err = o.QueryTable("IpsecServerHistory").OrderBy("-stamp").All(&records)
	if err != nil {
		beego.Warning(err)
	}

	type History struct {
		Time         string
		DisplayState string
		State        int
	}

	histories := make([]History, 0, len(records))
	for _, record := range records {
		h := History{}
		h.Time = time.Unix(record.Stamp, 0).Format("01-02 15:04:05")
		h.State = record.State
		h.DisplayState = this.getState(h.State)
		histories = append(histories, h)
	}
	if len(histories) > 15 {
		this.Data["Objects"] = histories[:15]
	} else {
		this.Data["Objects"] = histories
	}
}

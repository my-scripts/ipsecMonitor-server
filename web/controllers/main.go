package controllers

import (
	"fmt"
	"math"
	"time"
	"vpnportal/controllers/base"
	"vpnportal/models"
	"vpnportal/tasks"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

var flows []string

func init() {
	flows = make([]string, 0, 4)
	flows = append(flows, "Kb")
	flows = append(flows, "Mb")
	flows = append(flows, "Gb")
	flows = append(flows, "B")
}

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

func (this *MainController) BuildLoadData() {
	loadmon := tasks.GetLoadMonitor()
	values := loadmon.Data.Values()

	loadlabels := make([]string, len(values))
	loaddata := make([]float64, len(values))

	for i, v := range values {
		item := v.(tasks.LoadInfo)
		loadlabels[i] = fmt.Sprintf("%d:%d", item.Time.Hour(), item.Time.Minute())
		loaddata[i] = item.Load * 100
	}

	this.Data["LoadLabels"] = loadlabels
	this.Data["LoadData"] = loaddata
}

func (this *MainController) BuildIoData() {
	iomon := tasks.GetIoMonitor()
	values := iomon.Data.Values()

	iolables := make([]string, len(values))
	iorecv := make([]uint64, len(values))
	iosend := make([]uint64, len(values))

	for i := 1; i < len(values); i++ {
		pre := values[i-1].(tasks.IoInfo)
		cur := values[i].(tasks.IoInfo)

		iolables[i-1] = fmt.Sprintf("%d:%d", cur.Time.Hour(), cur.Time.Minute())

		rebit := cur.BytesRecv - pre.BytesRecv
		rekb := rebit / 1024
		remb := rekb / 1024
		regb := remb / 1024

		sebit := cur.BytesSend - pre.BytesSend
		sekb := sebit / 1024
		semb := sekb / 1024
		segb := semb / 1024

		if rebit > 1024 {
			iorecv[i-1] = rekb
			this.Data["IOFlow"] = flows[0]
			if rekb > 1024 {
				iorecv[i-1] = remb
				this.Data["IOFlow"] = flows[1]
				if remb > 1024 {
					iorecv[i-1] = regb
					this.Data["IOFlow"] = flows[2]
				}
			}

		}
		if rebit < 1024 {
			iorecv[i-1] = rebit
			this.Data["IOFlow"] = flows[3]
		}

		if sebit > 1024 {
			iosend[i-1] = sekb
			this.Data["IOFlow"] = flows[0]
			if sekb > 1024 {
				iosend[i-1] = semb
				this.Data["IOFlow"] = flows[1]
				if semb > 1024 {
					iosend[i-1] = segb
					this.Data["IOFlow"] = flows[2]
				}
			}

		}
		if sebit < 1024 {
			iosend[i-1] = sebit
			this.Data["IOFlow"] = flows[3]
		}
	}

	this.Data["IOLabels"] = iolables
	this.Data["IORecv"] = iorecv
	this.Data["IOSend"] = iosend
}

func (this *MainController) BuildMemoryData() {
	intermon := tasks.GetInterMoniter()
	values := intermon.Data.Values()

	interlables := make([]string, len(values))
	intermemory := make([]float64, len(values))

	for i, v := range values {
		item := v.(tasks.MemoryInfo)
		interlables[i] = fmt.Sprintf("%d:%d", item.Time.Hour(), item.Time.Minute())
		intermemory[i] = item.Memory
	}

	this.Data["InterLabels"] = interlables
	this.Data["InterMemory"] = intermemory
}

func (this *MainController) getState(state int) string {
	switch state {
	case models.OPENVPN_CLIENT_INIT:
		return "初始化"
	case models.OPENVPN_CLIENT_OFFLINE:
		return "离线"
	case models.OPENVPN_CLIENT_ONLINE:
		return "上线"
	case models.OPENVPN_CLIENT_RECONNECT:
		return "重连"
	}
	return ""
}

func (this *MainController) Get() {
	this.Layout = "layout.html"
	this.TplName = "index.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "index_head.html"
	this.LayoutSections["Script"] = "index_js.html"

	this.Data["MainPage"] = true

	info, err := host.Info()
	if err == nil {
		this.Data["Hostname"] = info.Hostname
		this.Data["Boottime"] = this.convertBoottime(info.BootTime)
	} else {
		beego.Warn("failed to get host info", err)
	}

	this.Data["Now"] = time.Now().Format("2006-01-02 15:04:05")

	ls, err := load.Avg()
	if err == nil {
		this.Data["LoadPercent"] = math.Floor(ls.Load1 * 100)
	}

	vm, err := mem.VirtualMemory()
	if err == nil {
		this.Data["MemPercent"] = float64(vm.Used) / float64(vm.Total)
	}

	ds, err := disk.Usage("/")
	if err == nil {
		this.Data["DiskPercent"] = ds.UsedPercent
	}

	this.BuildLoadData()
	this.BuildIoData()
	this.BuildMemoryData()

	o := orm.NewOrm()
	var records []models.OpenVpnHistory

	datas, err := o.QueryTable("OpenVpnHistory").OrderBy("-time").All(&records)
	if err != nil {
		beego.Warning(err)
	}
	this.Data["datas"] = datas

	type History struct {
		Time          string
		CN            string
		ConnectedTime string
		DisplayState  string
		State         int
	}

	histories := make([]History, 0, len(records))
	for _, record := range records {
		h := History{}
		h.Time = time.Unix(record.Time, 0).Format("01-02 15:04:05")
		h.CN = record.CN
		h.ConnectedTime = record.ConnectTime
		h.State = record.State
		h.DisplayState = this.getState(h.State)
		histories = append(histories, h)
	}
	if len(histories) > 15 {
		this.Data["Objects"] = histories[:15]
	}
}

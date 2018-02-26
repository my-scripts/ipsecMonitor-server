package tasks

import (
	"time"

	"github.com/aholic/ggtimer"
	"github.com/astaxie/beego"
	"github.com/shirou/gopsutil/load"
	"github.com/zfjagann/golang-ring"
)

type LoadInfo struct {
	Time time.Time
	Load float64
}

var task LoadMonitor

func GetLoadMonitor() *LoadMonitor {
	return &task
}

type LoadMonitor struct {
	Data ring.Ring
}

func (this *LoadMonitor) Start() {

	this.Data.SetCapacity(20)

	ggtimer.NewTicker(time.Duration(15)*time.Minute, func(time time.Time) {
		ls, err := load.Avg()
		if err != nil {
			beego.Warning("failed to get load,", err)
		}

		info := LoadInfo{Time: time, Load: ls.Load15}
		this.Data.Enqueue(info)
	})
}

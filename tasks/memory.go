package tasks

import (
	"time"

	"github.com/aholic/ggtimer"
	"github.com/astaxie/beego"
	"github.com/shirou/gopsutil/mem"
	"github.com/zfjagann/golang-ring"
)

type MemoryInfo struct {
	Time   time.Time
	Memory float64
}

var intermoniter InterMoniter

func GetInterMoniter() *InterMoniter {
	return &intermoniter
}

type InterMoniter struct {
	Data ring.Ring
}

func (this *InterMoniter) Start() {
	this.Data.SetCapacity(20)

	ggtimer.NewTicker(time.Duration(15)*time.Minute, func(time time.Time) {
		im, err := mem.VirtualMemory()
		if err != nil {
			beego.Warning("failed to get mem:", err)
		}

		minfo := MemoryInfo{Time: time, Memory: float64(im.Used) / float64(im.Total)}
		this.Data.Enqueue(minfo)
	})
}

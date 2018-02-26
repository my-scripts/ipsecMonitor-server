package tasks

import (
	"time"

	"github.com/aholic/ggtimer"
	"github.com/shirou/gopsutil/net"
	"github.com/zfjagann/golang-ring"
)

var iomonitor IoMonitor

func GetIoMonitor() *IoMonitor {
	return &iomonitor
}

type IoInfo struct {
	Time      time.Time
	BytesRecv uint64
	BytesSend uint64
}

type IoMonitor struct {
	Nic  string
	Data ring.Ring
}

func (this *IoMonitor) Start() {
	this.Data.SetCapacity(20)

	if this.Nic == "" {
		return
	}

	ggtimer.NewTicker(time.Duration(15)*time.Minute, func(time time.Time) {

		stats, err := net.IOCounters(true)
		if err != nil {
			return
		}

		for _, v := range stats {
			if v.Name == this.Nic {

				info := IoInfo{Time: time, BytesRecv: v.BytesRecv, BytesSend: v.BytesSent}
				this.Data.Enqueue(info)
			}
		}
	})
}

package service

import (
	"log"
	"path"
	"script/ipsecMonitor/server/base"
	"script/ipsecMonitor/server/web/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/howeyc/fsnotify"
)

type Monitor struct {
	Dir string
}

func (this *Monitor) Run() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		beego.Warning(err)
		return
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				pidfile := path.Base(ev.Name)
				if pidfile == "pluto.pid" && ev.IsModify() {
					beego.Warning("ipsec start")
					his := models.IpsecServerHistory{Stamp: time.Now().Unix(), State: 0}
					his.AddHistory()
					time.Sleep(time.Second * 1)
					go monitorPort()

				}
				if pidfile == "pluto.pid" && ev.IsDelete() {
					his := models.IpsecServerHistory{Stamp: time.Now().Unix(), State: 2}
					his.AddHistory()
					beego.Warning("ipsec is close")
				}
			case err := <-watcher.Error:
				beego.Warning("error:", err)
			}
		}
	}()

	err = watcher.Watch(this.Dir)
	if err != nil {
		log.Fatal(err)
	}

	<-done
	defer beego.Warning("watcher is close")
}

func notice() {
	beego.Warning("start notice client")
}

func monitorPort() {
	for i := 0; i < 3; i++ {
		if base.IpsecIsWork() {
			beego.Warning("ipsec working")
			his := models.IpsecServerHistory{Stamp: time.Now().Unix(), State: 1}
			his.AddHistory()
			go notice()
			break
		}
		time.Sleep(time.Second * 1)
	}

}

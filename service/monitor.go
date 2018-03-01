package service

import (
	"log"
	"path"
	"script/ipsecMonitor/server/base"
	serverrpc "script/ipsecMonitor/server/rpc"
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
	stamp := time.Now().Unix()
	clientport, _ := beego.AppConfig.Int("clientport")
	clients := models.GetClients()
	for k, _ := range clients {
		go func(index int) {
			vv := clients[index]
			client := serverrpc.IpsecRpcClient{}
			if !client.Connect(vv.Addr, clientport) {
				his := models.NoticeHistory{Stamp: time.Now().Unix(), Alias: vv.Alias, Success: false}
				his.AddHistory()
				return
			}
			defer client.Close()

			reply := client.RestartIpsec(stamp)

			his := models.NoticeHistory{Stamp: time.Now().Unix(), Alias: vv.Alias, Success: reply.Succ}
			his.AddHistory()
		}(k)
	}
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

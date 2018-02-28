package service

import (
	"fmt"
	"log"
	"path"
	"script/ipsecMonitor/server/base"

	"github.com/howeyc/fsnotify"
)

type Monitor struct {
	Dir string
}

func (this *Monitor) Run() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Println(err)
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
					if base.IpsecIsWork() {
						go notice()
					}
				}
				if pidfile == "pluto.pid" && ev.IsDelete() {
					fmt.Println("ipsec is close")
				}
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch(this.Dir)
	if err != nil {
		log.Fatal(err)
	}

	<-done
	defer log.Println("watcher is close")
}

func notice() {
	fmt.Println("start notice client")
}

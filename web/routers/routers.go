package routers

import (
	"script/ipsecMonitor/server/web/controllers"
	"script/ipsecMonitor/server/web/controllers/config"
	"script/ipsecMonitor/server/web/controllers/status/server"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/config/client/", &config.IpsecClientController{})
	beego.Router("/config/client/add", &config.IpsecClientController{}, "get:Add")
	beego.Router("/config/client/edit", &config.IpsecClientController{}, "get:Edit")
	beego.Router("/config/client/:id/del", &config.IpsecClientController{}, "get:Delete")

	beego.Router("/status/server/", &server.IpsecServerController{})
	beego.Router("/status/server/history/:page/", &server.IpsecServerHisController{})

	beego.Router("/server/start/", &server.IpsecServerController{}, "post:Start")
	beego.Router("/server/stop/", &server.IpsecServerController{}, "post:Stop")
	beego.Router("/server/restart/", &server.IpsecServerController{}, "post:Restart")

}

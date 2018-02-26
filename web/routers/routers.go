package routers

import (
	"script/ipsecMonitor/server/web/controllers"
	"script/ipsecMonitor/server/web/controllers/config"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/config/client/", &config.IpsecClientController{})
	beego.Router("/config/client/add", &config.IpsecClientController{}, "get:Add")
	beego.Router("/config/client/edit", &config.IpsecClientController{}, "get:Edit")
	beego.Router("/config/client/:id/del", &config.IpsecClientController{}, "get:Delete")
}

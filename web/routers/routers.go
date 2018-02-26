package routers

import (
	"script/ipsecMonitor/server/web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

}

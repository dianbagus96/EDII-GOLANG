package routers

import (
	"ClientOSS/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/oss/tdp", &controllers.TdpController{})
}

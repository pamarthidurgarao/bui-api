package routers

import (
	"bui-api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
    beego.Router("/customer/:id", &controllers.CustomerController{})
}

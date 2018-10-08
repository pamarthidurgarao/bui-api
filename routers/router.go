// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"bui-api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//main routing
	beego.Router("/", &controllers.MainController{})

	//Customer routing
	beego.Router("/customer/:id", &controllers.CustomerController{}, "get:GetById")
	beego.Router("/customer", &controllers.CustomerController{}, "post:AddCustomer")
	beego.Router("/customer", &controllers.CustomerController{}, "get:AllCustomers")
	
	//Service routing
	beego.Router("/service/:id", &controllers.ServiceController{}, "get:GetById")
	beego.Router("/service", &controllers.ServiceController{}, "post:AddService")
	beego.Router("/service", &controllers.ServiceController{}, "get:AllServices")
	
}

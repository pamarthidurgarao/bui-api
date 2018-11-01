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
	beego.Router("/api/customer/:id", &controllers.CustomerController{}, "get:GetById")
	beego.Router("/api/customer", &controllers.CustomerController{}, "post:AddCustomer")
	beego.Router("/api/customer", &controllers.CustomerController{}, "get:AllCustomers")
	beego.Router("/api/customer/:id", &controllers.CustomerController{}, "delete:DeleteCustomer")
	beego.Router("/api/customer/mobile/:q", &controllers.CustomerController{}, "get:SearchByMobile")

	//Service routing
	beego.Router("/api/service/:id", &controllers.ServiceController{}, "get:GetById")
	beego.Router("/api/service", &controllers.ServiceController{}, "post:AddService")
	beego.Router("/api/service", &controllers.ServiceController{}, "get:AllServices")
	beego.Router("/api/service", &controllers.ServiceController{}, "delete:DeleteService")

	// Staff Routing
	beego.Router("/api/staff/:id", &controllers.StaffController{}, "get:GetById")
	beego.Router("/api/staff", &controllers.StaffController{}, "post:AddStaff")
	beego.Router("/api/staff", &controllers.StaffController{}, "get:AllStaff")

	// Bill Routing
	beego.Router("/api/bill/:id", &controllers.BillController{}, "get:GetById")
	beego.Router("/api/bill", &controllers.BillController{}, "post:AddBill")
	beego.Router("/api/bill", &controllers.BillController{}, "get:AllBills")

	// Product Routing
	beego.Router("/api/product/:id", &controllers.ProductController{}, "get:GetById")
	beego.Router("/api/product", &controllers.ProductController{}, "post:AddProduct")
	beego.Router("/api/product", &controllers.ProductController{}, "get:AllProducts")

}

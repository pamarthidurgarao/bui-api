package controllers

import (
	"bui-api/constants"
	"bui-api/models"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type BillController struct {
	MainController
}

func (c *BillController) GetById() {
	id := c.Ctx.Input.Param(":id")
	beego.Info(id)
	query := make(models.Map, 0)
	query["_id"] = bson.ObjectIdHex(id)
	d := models.Find(constants.DATABASE, constants.BILL, query)
	c.Data["json"] = map[string]interface{}{"data": d}
	beego.Info(d)
	c.ServeJSON()
}

func (c *BillController) AllBills() {

	d := models.Find(constants.DATABASE, constants.BILL, nil)
	c.Data["json"] = map[string]interface{}{"data": d}
	beego.Info(d)
	c.ServeJSON()
}

func (c *BillController) AddBill() {
	req := c.GetRequestBody()
	beego.Info(req)
	bill := make(models.Map)
	var rawId = req["_id"]
	var id = ""
	if rawId != nil {
		id = rawId.(string)
	}
	if len(id) == 0 {
		bill["customer"] = req["customer"].(string)
		bill["total"] = req["total"].(string)
		bill["discount"] = req["discount"].(string)

		var services = req["services"]
		var reqBody = ItoMapArray1(services)
		resServices := make([]models.Map, 0)
		for _, b := range reqBody {
			service := make(models.Map)
			var db = b.(map[string]interface{})
			service["_id"] = db["_id"].(string)
			service["name"] = db["name"].(string)
			service["price"] = db["price"].(string)
			service["discount"] = db["discount"].(string)
			service["netPrice"] = db["netPrice"].(string)

			var staf = db["staff"]
			var st = staf.(map[string]interface{})
			staff := make(models.Map)
			staff["_id"] = st["_id"].(string)
			staff["name"] = st["name"].(string)
			service["staff"] = staff
			resServices = append(resServices, service)
		}
		bill["services"] = resServices
		res, _ := models.Create(constants.DATABASE, constants.BILL, bill)
		c.Data["json"] = map[string]interface{}{"response": res}
		c.ServeJSON()
	} else {
		beego.Info("Inside update")
		query := make(models.Map)
		query["_id"] = bson.ObjectIdHex(id)
		bill["_id"] = bson.ObjectIdHex(id)
		bill["customer"] = req["customer"].(string)
		bill["total"] = req["total"].(string)
		bill["discount"] = req["discount"].(string)

		var services = req["services"]
		var reqBody = ItoMapArray1(services)
		resServices := make([]models.Map, 0)
		for _, b := range reqBody {
			service := make(models.Map)
			var db = b.(map[string]interface{})
			service["_id"] = db["_id"].(string)
			service["name"] = db["name"].(string)
			service["price"] = db["price"].(string)
			service["discount"] = db["discount"].(string)
			service["netPrice"] = db["netPrice"].(string)

			var staf = db["staff"]
			var st = staf.(map[string]interface{})
			staff := make(models.Map)
			staff["_id"] = st["_id"].(string)
			staff["name"] = st["name"].(string)
			service["staff"] = staff
			resServices = append(resServices, service)
		}
		bill["services"] = resServices
		res := models.Update(constants.DATABASE, constants.BILL, query, bill)
		c.Data["json"] = map[string]interface{}{"response": res}
		c.ServeJSON()
	}
}

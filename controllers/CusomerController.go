package controllers

import (
	"github.com/astaxie/beego"
	"bui-api/models"
	"bui-api/constants"
	"gopkg.in/mgo.v2/bson"
)

type CustomerController struct {
	MainController
}

func (c *CustomerController) GetById() {
	id := c.Ctx.Input.Param(":id")
	beego.Info(id);
	query := make(models.Map,0)
	query["_id"] = bson.ObjectIdHex(id)
	d := models.Find(constants.DATABASE,constants.CUSTOMER,query)
	c.Data["json"] = map[string]interface{}{"data":d}
	beego.Info(d);
	c.ServeJSON();	
} 

func (c *CustomerController) AllCustomers() {

	d := models.Find(constants.DATABASE,constants.CUSTOMER,nil)
	c.Data["json"] = map[string]interface{}{"data":d}
	beego.Info(d);
	c.ServeJSON();	
}

func (c *CustomerController) UpdateCustomer() {

	d := models.Find(constants.DATABASE,constants.CUSTOMER,nil)
	c.Data["json"] = map[string]interface{}{"data":d}
	beego.Info(d);
	c.ServeJSON();	
} 

func (c *CustomerController) AddCustomer() {
	req := c.GetRequestBody()
	beego.Info(req);
	customer:=make(models.Map)
	customer["fullname"] = req["fullname"].(string)
	customer["phone"] = req["phone"].(string)
	customer["email"] = req["email"].(string)
	customer["gender"] = req["gender"].(string)
	customer["dob"] = req["dob"].(string)
	res,_ := models.Create(constants.DATABASE,constants.CUSTOMER,customer)
	c.Data["json"] = map[string]interface{}{"response":res}
	c.ServeJSON();	
} 


func (c *CustomerController) DeleteCustomer() {

	d := models.Find(constants.DATABASE,constants.CUSTOMER,nil)
	c.Data["json"] = map[string]interface{}{"data":d}
	beego.Info(d);
	c.ServeJSON();	
}
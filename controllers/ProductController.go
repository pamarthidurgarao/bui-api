package controllers

import (
	"bui-api/constants"
	"bui-api/models"

	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type ProductController struct {
	MainController
}

func (c *ProductController) GetById() {
	id := c.Ctx.Input.Param(":id")
	beego.Info(id)
	query := make(models.Map, 0)
	query["_id"] = bson.ObjectIdHex(id)
	d := models.Find(constants.DATABASE, constants.PRODUCT, query)
	c.Data["json"] = map[string]interface{}{"data": d}
	beego.Info(d)
	c.ServeJSON()
}

func (c *ProductController) AllProducts() {

	d := models.Find(constants.DATABASE, constants.PRODUCT, nil)
	c.Data["json"] = map[string]interface{}{"data": d}
	beego.Info(d)
	c.ServeJSON()
}

func (c *ProductController) AddProduct() {
	req := c.GetRequestBody()
	beego.Info(req)
	product := make(models.Map)
	var rawId = req["_id"]
	var id = ""
	if rawId != nil {
		id = rawId.(string)
	}
	if len(id) == 0 {
		product["name"] = req["name"].(string)
		product["price"] = req["price"].(string)
		product["quantity"] = req["quantity"].(string)
		product["totalQuantity"] = req["totalQuantity"].(string)
		res, _ := models.Create(constants.DATABASE, constants.PRODUCT, product)
		c.Data["json"] = map[string]interface{}{"response": res}
		c.ServeJSON()
	} else {
		beego.Info("Inside update")
		query := make(models.Map)
		query["_id"] = bson.ObjectIdHex(id)
		product["_id"] = bson.ObjectIdHex(id)
		product["name"] = req["name"].(string)
		product["price"] = req["price"].(string)
		product["quantity"] = req["quantity"].(string)
		product["totalQuantity"] = req["totalQuantity"].(string)
		res := models.Update(constants.DATABASE, constants.PRODUCT, query, product)
		c.Data["json"] = map[string]interface{}{"response": res}
		c.ServeJSON()
	}
}

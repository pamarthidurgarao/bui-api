package controllers

import (
	"bui-api/constants"
	"bui-api/models"
	"strconv"
	"time"

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
	query := make(models.Map)
	query["category"] = req["category"].(string)
	d := models.FindOne(constants.DATABASE, constants.PRODUCT, query)
	typesResponse := make([]models.Map, 0)
	if d != nil && len(d) > 0 {
		beego.Info("Loop inside")
		beego.Info(d)
		product := make(models.Map)
		var body = req["products"]
		var reqBody = ItoMapArray1(body)
		for _, b := range reqBody {
			beego.Info(b)
			var db = b.(map[string]interface{})
			var ypes = d["products"]
			var response = ItoMapArray(ypes)
			if db["id"] != nil {
				for _, data := range response {
					if data["id"] == db["id"] {
						data["name"] = db["name"].(string)
						data["brand"] = db["brand"].(string)
						data["price"] = db["price"].(float64)
						data["stock"] = db["stock"].(float64)
						data["quantity"] = db["quantity"].(string)
						typesResponse = append(typesResponse, data)
						beego.Info(data)
					} else {
						typesResponse = append(typesResponse, data)
						beego.Info(data)
					}
				}
			} else {
				product["name"] = db["name"].(string)
				product["brand"] = db["brand"].(string)
				product["price"] = db["price"].(float64)
				product["stock"] = db["stock"].(float64)
				product["quantity"] = db["quantity"].(string)
				product["id"] = "service_" + strconv.FormatInt(time.Now().UnixNano(), 10)
				typesResponse = ItoMapArray(ypes)
				typesResponse = append(typesResponse, product)
			}
			model := make(models.Map)
			model["category"] = d["category"].(string)
			model["_id"] = d["_id"].(bson.ObjectId)
			model["products"] = typesResponse
			res := models.Update(constants.DATABASE, constants.PRODUCT, query, model)
			c.Data["json"] = map[string]interface{}{"response": res}
			c.ServeJSON()
		}

	} else {
		model := make(models.Map)
		products := make([]models.Map, 0)
		product := make(models.Map)
		model["category"] = req["category"].(string)
		var body = req["products"]
		var reqBody = ItoMapArray1(body)
		for _, b := range reqBody {
			beego.Info(b)
			var db = b.(map[string]interface{})
			product["name"] = db["name"].(string)
			product["brand"] = db["brand"].(string)
			product["price"] = db["price"].(float64)
			product["stock"] = db["stock"].(float64)
			product["quantity"] = db["quantity"].(string)
		}
		product["id"] = "service_" + strconv.FormatInt(time.Now().UnixNano(), 10)
		products = append(products, product)
		model["products"] = products
		res, _ := models.Create(constants.DATABASE, constants.PRODUCT, model)
		c.Data["json"] = map[string]interface{}{"response": res}
		c.ServeJSON()
	}
}

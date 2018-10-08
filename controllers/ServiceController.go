package controllers

import (
	"strconv"
	"time"
	"reflect"
	"github.com/astaxie/beego"
	"bui-api/models"
	"bui-api/constants"
	"gopkg.in/mgo.v2/bson"
)

type ServiceController struct {
	MainController
}

func (c *ServiceController) GetById() {
	id := c.Ctx.Input.Param(":id")
	beego.Info(id);
	query := make(models.Map,0)
	query["_id"] = bson.ObjectIdHex(id)
	d := models.Find(constants.DATABASE,constants.SERVICE,query)
	c.Data["json"] = map[string]interface{}{"data":d}
	beego.Info(d);
	c.ServeJSON();	
} 

func (c *ServiceController) AllServices() {

	d := models.Find(constants.DATABASE,constants.SERVICE,nil)
	c.Data["json"] = map[string]interface{}{"data":d}
	beego.Info(d);
	c.ServeJSON();	
}

func (c *ServiceController) UpdateService() {

	d := models.Find(constants.DATABASE,constants.SERVICE,nil)
	c.Data["json"] = map[string]interface{}{"data":d}
	beego.Info(d);
	c.ServeJSON();	
} 

func (c *ServiceController) AddService() {
	req := c.GetRequestBody()
	beego.Info(req);
	query:=make(models.Map)
	query["category"] = req["category"].(string)
	d := models.FindOne(constants.DATABASE,constants.SERVICE,query)
	if d != nil && len(d) > 0 {
		beego.Info("Loop inside");
		beego.Info(d);
		sType :=make(models.Map)
		sType["id"]=  "service_"+strconv.FormatInt(time.Now().UnixNano(), 10)
		sType["name"] = req["name"].(string)
		sType["time"] = req["time"].(string)
		sType["price"] = req["price"].(string)
		sType["gender"] = req["gender"].(string)
		// types:=make([]models.Map,1)
		var ypes = d["types"];
		object := reflect.ValueOf(ypes)
		// bson.Unmarshal([]byte(object), &types)
		// beego.Info(ypes.Type());
		var items []interface{}
		for i := 0; i < object.Len(); i++ {
			items = append(items, object.Index(i).Interface())
		}	
		// object = append(object,sType);
	}else {
		model:=make(models.Map)
		types:=make([]models.Map,1)
		sType :=make(models.Map)
		model["category"] = req["category"].(string)
		sType["id"]=  "service_"+strconv.FormatInt(time.Now().UnixNano(), 10)
		sType["name"] = req["name"].(string)
		sType["time"] = req["time"].(string)
		sType["price"] = req["price"].(string)
		sType["gender"] = req["gender"].(string)
		types = append(types,sType);
		model["types"] = types;
		res,_ := models.Create(constants.DATABASE,constants.SERVICE,model)
		c.Data["json"] = map[string]interface{}{"response":res}
		c.ServeJSON();	
	}
} 


func (c *ServiceController) DeleteService() {
	d := models.Find(constants.DATABASE,constants.SERVICE,nil)
	c.Data["json"] = map[string]interface{}{"data":d}
	beego.Info(d);
	c.ServeJSON();	
}
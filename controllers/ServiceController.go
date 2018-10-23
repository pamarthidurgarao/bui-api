package controllers

import (
	"strconv"
	"time"
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
	typesResponse:=make([]models.Map,0)
	if d != nil && len(d) > 0 {
		beego.Info("Loop inside");
		beego.Info(d);
		sType :=make(models.Map)
		var body=req["types"];
		var reqBody=ItoMapArray1(body);
		for _, b := range reqBody {
			beego.Info(b);
			var db = b.(map[string]interface{});
			var ypes = d["types"];
			var response = ItoMapArray(ypes)
			if db["id"] != nil {
				for _, data := range response {
					if data["id"] == db["id"] {
						data["name"] = db["name"].(string)
						data["time"] = db["time"].(float64)
						data["price"] = db["price"].(string)
						data["gender"] = db["gender"].(string)
						typesResponse = append(typesResponse,data);	
						beego.Info(data);
					}else{
						typesResponse = append(typesResponse,data);
						beego.Info(data);
					}
				}
			}else{
				sType["name"] = db["name"].(string)
				sType["time"] = db["time"].(float64)
				sType["price"] = db["price"].(string)
				sType["gender"] = db["gender"].(string)
				sType["id"]=  "service_"+strconv.FormatInt(time.Now().UnixNano(), 10)
				typesResponse = ItoMapArray(ypes)
				typesResponse = append(typesResponse,sType);
			}
			model:=make(models.Map)
			model["category"] = d["category"].(string)
			model["_id"] = d["_id"].(bson.ObjectId)
			model["types"] = typesResponse;
			res:= models.Update(constants.DATABASE,constants.SERVICE,query,model)
			c.Data["json"] = map[string]interface{}{"response":res}
			c.ServeJSON();	
		}
		
	}else {
		model:=make(models.Map)
		types:=make([]models.Map,1)
		sType :=make(models.Map)
		model["category"] = req["category"].(string)
		sType["id"]=  "service_"+strconv.FormatInt(time.Now().UnixNano(), 10)
		var body=req["types"];
		var reqBody=ItoMapArray1(body);
		for _, b := range reqBody {
			beego.Info(b);
			var db = b.(map[string]interface{});
			sType["name"] = db["name"].(string)
			sType["time"] = db["time"].(float64)
			sType["price"] = db["price"].(string)
			sType["gender"] = db["gender"].(string)
		}
		sType["id"]=  "service_"+strconv.FormatInt(time.Now().UnixNano(), 10)
		types = append(types,sType);
		model["types"] = types;
		res,_ := models.Create(constants.DATABASE,constants.SERVICE,model)
		c.Data["json"] = map[string]interface{}{"response":res}
		c.ServeJSON();	
	}
} 


func (c *ServiceController) DeleteService() {
	req := c.GetRequestBody()
	beego.Info(req);
	query:=make(models.Map)
	query["category"] = req["category"].(string)
	d := models.FindOne(constants.DATABASE,constants.SERVICE,query)
	if d != nil && len(d) > 0 {
		beego.Info("Loop inside");
		beego.Info(d);
		var body=req["types"];
		var reqBody=ItoMapArray1(body);
		typesResponse:=make([]models.Map,1)
		for _, b := range reqBody {
			beego.Info(b);
			var db = b.(map[string]interface{});
			var serviceId = db["id"].(string)
			var ypes = d["types"];
			var types = ItoMapArray(ypes)
			for _, b := range types {
				if b["id"] != serviceId {
					typesResponse = append(typesResponse,b);
				}
			}
		}
		model:=make(models.Map)
		model["category"] = d["category"].(string)
		model["_id"] = d["_id"].(bson.ObjectId)
		model["types"] = typesResponse;
		res:= models.Update(constants.DATABASE,constants.SERVICE,query,model)
		c.Data["json"] = map[string]interface{}{"response":res}
		c.ServeJSON();	
	}else {
		c.Data["json"] = map[string]interface{}{"response":"Invalid Data"}
		c.ServeJSON();	
	}	
}

func ItoMapArray(data interface{}) []models.Map {
	retArray := make([]models.Map, 0, 10)
	elements := data.([]interface{})
		for _, v := range elements {
		obj := v.(models.Map)
		retArray = append(retArray, obj)
	}
	return retArray
}

func ItoMapArray1(data interface{}) []interface{} {
	retArray := make([]interface{}, 0, 10)
	elements := data.([]interface{})
		for _, v := range elements {
		retArray = append(retArray, v)
	}
	return retArray
}

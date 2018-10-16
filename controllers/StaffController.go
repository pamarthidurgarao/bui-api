package controllers

import (
	"github.com/astaxie/beego"
	"bui-api/models"
	"bui-api/constants"
	"gopkg.in/mgo.v2/bson"
)

type StaffController struct {
	MainController
}

func (c *StaffController) GetById() {
	id := c.Ctx.Input.Param(":id")
	beego.Info(id);
	query := make(models.Map,0)
	query["_id"] = bson.ObjectIdHex(id)
	d := models.Find(constants.DATABASE,constants.STAFF,query)
	c.Data["json"] = map[string]interface{}{"data":d}
	beego.Info(d);
	c.ServeJSON();	
} 

func (c *StaffController) AllStaff() {

	d := models.Find(constants.DATABASE,constants.STAFF,nil)
	c.Data["json"] = map[string]interface{}{"data":d}
	beego.Info(d);
	c.ServeJSON();	
}

func (c *StaffController) AddStaff() {
	req := c.GetRequestBody()
	beego.Info(req);
	staff:=make(models.Map)
	var rawId = req["_id"]
	var id = "";
	if(rawId!=nil){
		id=rawId.(string)
	}
	if len(id) == 0 {	
		staff["name"] = req["name"].(string)
		staff["mobile"] = req["mobile"].(string)
		staff["preferedGender"] = req["preferedGender"].(string)
		staff["gender"] = req["gender"].(string)
		staff["position"] = req["position"].(string)
		var reqPreferences=ItoMapArray1(req["preferences"]);
		preferences:=make([]models.Map,0)
		for _, b := range reqPreferences {
			beego.Info(b);
			preference:=make(models.Map)
			var db = b.(map[string]interface{});
			preference["id"] = db["id"].(string)
			preference["name"] = db["name"].(string)
			preferences = append(preferences,preference)
		}
		staff["preferences"]=preferences;
		var reqTimings=ItoMapArray1(req["timings"]);
		timings:=make([]models.Map,0)
		for _, b := range reqTimings {
			beego.Info(b);
			time:=make(models.Map)
			var db = b.(map[string]interface{});
			time["day"] = db["day"].(string)
			time["from"] = db["from"].(string)
			time["to"] = db["to"].(string)
			timings = append(timings,time)
		}
		staff["timings"]=timings;
		res,_ := models.Create(constants.DATABASE,constants.STAFF,staff)
		c.Data["json"] = map[string]interface{}{"response":res}
		c.ServeJSON();	
	}else{
		beego.Info("Inside update");
		query:=make(models.Map)
		query["_id"] = bson.ObjectIdHex(id)
		staff["_id"] = bson.ObjectIdHex(id)
		staff["name"] = req["name"].(string)
		staff["mobile"] = req["mobile"].(string)
		staff["preferedGender"] = req["preferedGender"].(string)
		staff["gender"] = req["gender"].(string)
		staff["position"] = req["position"].(string)
		var reqPreferences=ItoMapArray1(req["preferences"]);
		preferences:=make([]models.Map,0)
		for _, b := range reqPreferences {
			beego.Info(b);
			preference:=make(models.Map)
			var db = b.(map[string]interface{});
			preference["id"] = db["id"].(string)
			preference["name"] = db["name"].(int)
			preferences = append(preferences,preference)
		}
		staff["preferences"]=preferences;
		var reqTimings=ItoMapArray1(req["timings"]);
		timings:=make([]models.Map,0)
		for _, b := range reqTimings {
			beego.Info(b);
			time:=make(models.Map)
			var db = b.(map[string]interface{});
			time["day"] = db["day"].(string)
			time["from"] = db["from"].(int)
			time["to"] = db["to"].(string)
			timings = append(timings,time)
		}
		staff["timings"]=timings;
		res:= models.Update(constants.DATABASE,constants.STAFF,query,staff)
		c.Data["json"] = map[string]interface{}{"response":res}
		c.ServeJSON();	
	}
} 

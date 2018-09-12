package controllers

import (
	"github.com/astaxie/beego"
	"github.com/bui-api/models/connect"
)

type CustomerController struct {
	beego.Controller
}

func (c *CustomerController) Get() {
	id := c.GetString("key")
	c.Data["json"] = map[string]interface{}{"ObjectId": id }
	&connect.getConnectionURI("test")
	c.ServeJSON();	
} 
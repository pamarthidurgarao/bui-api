package controllers

import (
	"github.com/astaxie/beego"
)

type CustomerController struct {
	beego.Controller
}

func (c *CustomerController) Get() {
	id := c.GetString("key")
	c.Data["json"] = map[string]interface{}{"ObjectId": id }
	c.ServeJSON();	
} 
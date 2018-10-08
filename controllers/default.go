package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"bui-api/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) GetRequestBody () models.Map {
	req := make(models.Map)
	beego.Info(c.Ctx.Input.RequestBody);
	json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	return req
} 

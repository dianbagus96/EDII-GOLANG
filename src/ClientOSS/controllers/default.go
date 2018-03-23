package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "oss.ekon.go.id"
	c.Data["Email"] = "Tim Teknis INSW @2017"
	c.TplName = "index.tpl"
}

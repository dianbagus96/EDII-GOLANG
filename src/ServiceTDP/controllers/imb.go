package controllers

import (
	"ServiceTDP/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about object
type ImbController struct {
	beego.Controller
}

/*
func (o *ObjectController) Post() {
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	obs := models.AddOne(ob)
	//o.Data["json"] = map[string]string{"ObjectId": objectid}
	res, _ := json.Marshal(obs)
	o.Data["json"] = checksum(res)
	o.ServeJSON()
}
*/

func (o *ImbController) Post() {
	var ob models.Imb
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	obs := models.AddOneImb(ob)

	o.Data["json"] = obs
	o.ServeJSON()
}

func (o *ImbController) Get() {

	o.ServeJSON()
}

func (o *ImbController) GetAll() {
	o.Data["json"] = "OSS Service IMB"
	o.ServeJSON()
}

func (o *ImbController) Put() {

	o.ServeJSON()
}

func (o *ImbController) Delete() {

	o.ServeJSON()
}

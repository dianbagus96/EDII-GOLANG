package controllers

import (
	"ServiceOSS/models"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/astaxie/beego"
)

func checksum(res []byte) map[string]interface{} {
	row := make(map[string]interface{})
	s := string(res[:])
	row["checksum"] = GetMD5Hash(s)
	row["data"] = res
	return row
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	//str := base64.StdEncoding.EncodeToString([]byte(text))
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Operations about object
type ObjectController struct {
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

func (o *ObjectController) Post() {
	//var ob models.Object
	//json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	//obs := models.AddOne(ob)

	o.Data["json"] = "asd"
	o.ServeJSON()
}

func (o *ObjectController) Get() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		ob := models.GetOne(objectId)
		res, _ := json.Marshal(ob)
		o.Data["json"] = checksum(res)
	}
	o.ServeJSON()
}

func (o *ObjectController) GetAll() {
	obs := models.GetAll()
	res, _ := json.Marshal(obs)
	o.Data["json"] = checksum(res)
	o.ServeJSON()
}

func (o *ObjectController) Put() {
	ossId := o.Ctx.Input.Param(":objectId")
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	obs := models.Update(ossId, ob)
	res, _ := json.Marshal(obs)
	o.Data["json"] = checksum(res)
	o.ServeJSON()
}

func (o *ObjectController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	if objectId != "" {
		obs := models.Delete(objectId)
		res, _ := json.Marshal(obs)
		o.Data["json"] = checksum(res)
	} else {
		o.Data["json"] = "No Data Deleted !"
	}
	o.ServeJSON()
}

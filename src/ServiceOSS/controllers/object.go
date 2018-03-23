package controllers

import (
	"ServiceOSS/models"
	"crypto/md5"
	"time"

	"encoding/hex"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/astaxie/beego"
)

var (
	Log        *log.Logger
	logpath    = flag.String("logpath", "log/Request/", "Log Path")
	logpathRes = flag.String("logpaths", "log/Response/", "Log Paths")
)

func NewLog(logpath string) {
	println("LogFile: " + logpath)
	now := time.Now()

	file, err := os.Create(logpath + now.String())
	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}

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
	var ob models.PPM
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	flag.Parse()
	NewLog(*logpath)
	Log.Println(string(o.Ctx.Input.RequestBody))

	obs := models.AddOnes(ob)

	//fmt.Println(obs)
	//o.Data["json"] = map[string]string{"ObjectId": objectid}
	//res, _ := json.Marshal(obs)
	//o.Data["json"] = checksum(res)

	NewLog(*logpathRes)
	res, _ := json.Marshal(obs)
	Log.Println(string(res))

	o.Data["json"] = obs
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

/*
decypt
b := Decrypt(a)
	o.Data["json"] = string(b)
	//o.Data["json"] = obs
	//a := Encrypt(obs)
	//b := Decrypt(a)
*/

func (o *ObjectController) GetAll() {
	/*obs := models.GetAll()
	res, _ := json.Marshal(obs)
	o.Data["json"] = checksum(res)
	*/
	o.Data["json"] = "Welcome to Services Sending IP OSS V.1.0 Dev"
	o.ServeJSON()
}

func (o *ObjectController) Put() {
	/*ossId := o.Ctx.Input.Param(":objectId")
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	obs := models.Update(ossId, ob)
	res, _ := json.Marshal(obs)
	o.Data["json"] = checksum(res)*/
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

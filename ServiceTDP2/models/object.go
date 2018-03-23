package models

import (
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Objects map[string]*Object
)

func init() {
	//orm.RegisterDataBase("default", "mysql", dbUser+":"+dbPass+"@"+dbHost+"/"+dbName+"?charset=utf8", 30)
	//orm.RegisterDataBase("default", "mysql", dbUser+":"+dbPass+"@tcp("+dbHost+")/"+dbName+"?charset=utf8", 30)
	//orm.RegisterDataBase("default", "mysql", "root:@/isbgdb?charset=utf8", 30)
	//orm.RegisterDataBase("default", "mysql", "ossuser:db0ss2017@tcp(10.1.6.87:3306)/oss?charset=utf8", 30)
	dbConfig := beego.AppConfig.String("dbConfig")
	orm.RegisterDataBase("default", "mysql", dbConfig, 30)
	orm.Debug = true
	Objects = make(map[string]*Object)
}

func AddOne(Data Object) []map[string]interface{} {

	list := make([]map[string]interface{}, 1)
	row := make(map[string]interface{})
	/* S: INSERT HEADER */
	err := InsertTDP(Data)
	if err == nil {
		row["ResponseCode"] = "00"
		row["ResponseDescription"] = "Success"
		list[0] = row
		return list
	} else {
		row["ResponseCode"] = "99"
		row["ResponseDescription"] = err.Error()
		list[0] = row
		return list
	}

}

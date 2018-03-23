package models

import (
	_ "github.com/go-sql-driver/mysql"
)

func AddOneImb(Data Imb) []map[string]interface{} {

	list := make([]map[string]interface{}, 1)
	row := make(map[string]interface{})
	/* S: INSERT HEADER */
	err := InsertIMB(Data)
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

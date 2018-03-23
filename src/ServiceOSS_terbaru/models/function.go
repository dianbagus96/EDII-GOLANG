package models

import (
	"errors"
	"math/rand"
	"time"

	"github.com/astaxie/beego/orm"
)

func shuffle(Title string) string {
	src := []string{
		"A", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	}

	final := make([]string, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}

	var char string
	for p, _ := range final {
		char += final[p]
	}

	runes := []rune(char)
	var count int = 15 - len(Title)
	safeSubstring := Title + string(runes[0:count])
	return safeSubstring
}

func generateCode(Title string) string {
	shuffled := shuffle(Title)
	return shuffled
}

func InsertHdr(Object Object) (Res bool, Errors error, OSSID string) {
	o := orm.NewOrm()
	oss_id := generateCode("OSSID")

	sql := "INSERT INTO tblfcosshdr(oss_id, permit_group, oss_date, company_id, applicant_id, status, f_prepare, f_construction, f_production, created_by, created_date) VALUES (?, ?, NOW(), ?, ?, 0, 0, 0, 0, 'ServiceEngine', NOW()) "
	res, err := o.Raw(sql, oss_id, Object.PermitGroup, Object.CompanyID, Object.ApplicantID).Exec()

	if err == nil {
		num, _ := res.RowsAffected()
		if num > 0 {
			return true, nil, oss_id
		} else {
			return false, errors.New("Failed to Insert"), ""
		}
	} else {

		return false, errors.New("Error Insert Header"), ""
	}

	return false, errors.New("Exception Error"), ""
}

func InsertMaster(Object PPM) (bool, []map[string]interface{}) {

	list := make([]map[string]interface{}, 1)
	row := make(map[string]interface{})

	o := orm.NewOrm()
	o.Begin()

	var bol bool = true
	errRes := InsertDmCompany(Object)
	if errRes != nil {
		row["ResponseCode"] = "99"
		row["ResponseDescription"] = errRes.Error()
		list[0] = row
		bol = false
		o.Rollback()
	}

	errRes2 := InsertCompanyDeed(Object)
	if errRes2 != nil {
		row["ResponseCode"] = "99"
		row["ResponseDescription"] = errRes2.Error()
		list[0] = row
		bol = false
		o.Rollback()
	}

	errRes3, shareID := InsertCompanyShare(Object)
	if errRes3 != nil {
		row["ResponseCode"] = "99"
		row["ResponseDescription"] = errRes3.Error()
		list[0] = row
		bol = false
		o.Rollback()
	}

	errRes4 := InsertCompanyApplicant(Object)
	if errRes4 != nil {
		row["ResponseCode"] = "99"
		row["ResponseDescription"] = errRes4.Error()
		list[0] = row
		bol = false
		o.Rollback()
	}

	errRes5 := InsertShareholders(Object, shareID)
	if errRes5 != nil {
		row["ResponseCode"] = "99"
		row["ResponseDescription"] = errRes5.Error()
		list[0] = row
		bol = false
		o.Rollback()
	}

	o.Commit()
	return bol, list
}

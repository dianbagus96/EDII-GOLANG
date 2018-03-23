package models

import (
	"errors"
	"math/rand"
	"strings"
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

func InsertHdr(Object Object3) (Res bool, Errors error, OSSID string) {
	o := orm.NewOrm()
	oss_id := generateCode("OSSID")

	sql := "INSERT INTO tblfcosshdr(oss_id, permit_group, oss_date, company_id, applicant_id, status, f_prepare, f_construction, f_production, created_by, created_date) VALUES (?, ?, NOW(), ?, ?, 0, 0, 0, 0, 'ServiceEngine', NOW()) "
	res, err := o.Raw(sql).Exec()

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

func InsertPPM(Object PPM) (Res bool, Errors error, OSSID string) {
	o := orm.NewOrm()
	var maps []orm.Params

	oss_id := generateCode("OSSID")
	num, _ := o.Raw("SELECT ppm_id FROM tblfcossppm WHERE ppm_id = ? AND oss_id = ?", Object.PPM.IDip, Object.PPM.Oss_id).Values(&maps)

	if num > 0 {
		return true, nil, ""
	} else {
		sql := `
				INSERT INTO tblfcossppm(ppm_id, oss_id, business_id,  company_id, institute_id, reg_status, created_date, created_by, ppm_no, ppm_date, ppm_sign_name, ppm_sign_date, ppm_sign_position, ppm_sign_ga)
				VALUES (?, ?, ?, ?, ?, ?, NOW(), 'ServiceOSS', ?, ?, ?, ?, ?, ?)
			`
		res, err := o.Raw(sql, Object.PPM.IDip, Object.PPM.Oss_id, ``, Object.PPM.IDPerusahaan, `1`, `1`, Object.PPM.No_ppm, Object.PPM.Tgl_ppm, Object.PPM.Nama_ttd_ppm, Object.PPM.Tgl_ttd_ppm, Object.PPM.Jabatan_ttd_ppm, `06`).Exec()

		if err == nil {
			num, _ := res.RowsAffected()
			if num > 0 {
				return true, nil, oss_id
			} else {
				return false, errors.New("Failed to Insert"), ""
			}
		} else {
			panic(err.Error())
			return false, errors.New("Error Insert PPM"), ""
		}
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

	errRes6 := InsertCompanyBusiness(Object)
	if errRes6 != nil {
		row["ResponseCode"] = "99"
		row["ResponseDescription"] = errRes6.Error()
		list[0] = row
		bol = false
		o.Rollback()
	}

	o.Commit()
	return bol, list
}

func getPropID(Name string, types string) string {
	o := orm.NewOrm()
	var maps []MasterCode

	var sql string
	var hasil string

	if types == "prop" {
		sql = "SELECT prop_id 'id' FROM tbldmpropinsi WHERE prop_name LIKE '%" + Name + "%'  LIMIT 1"
	} else if types == "kab" {
		sql = `SELECT kab_id 'id' FROM tbldmkabupaten WHERE kab_name LIKE '%` + Name + `%'  LIMIT 1`
	} else if types == "kec" {
		sql = `SELECT kec_id 'id' FROM tbldmkecamatan WHERE kec_name LIKE '%` + Name + `%'  LIMIT 1`
	} else if types == "kel" {
		sql = `SELECT kel_id 'id' FROM tbldmkelurahan WHERE kel_name LIKE '%` + Name + `%'  LIMIT 1`
	} else if types == "country" {
		sql = `SELECT country_id FROM tbldmcountry WHERE country_name LIKE '%` + Name + `%'  LIMIT 1`
	}

	num, err := o.Raw(sql).QueryRows(&maps)
	if err != nil {
		panic(err.Error() + `|` + types)
	}

	if num == 0 {
		hasil = ""

	} else {
		hasil = maps[0].Id
	}

	return hasil
}

func kbliSlice(kbli string) string {
	s := strings.Split(kbli, "_")
	return s[0]
}

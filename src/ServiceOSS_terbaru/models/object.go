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

	dbConfig := beego.AppConfig.String("dbConfig")
	orm.RegisterDataBase("default", "mysql", dbConfig, 30)
	orm.Debug = true

	Objects = make(map[string]*Object)
}

func AddOne(Data Permit) []map[string]interface{} {

	list := make([]map[string]interface{}, 1)
	row := make(map[string]interface{})

	o := orm.NewOrm()

	/* S: INSERT HEADER */
	bol, er, oss_id := InsertHdr(Data)
	if bol == false {
		row["ResponseCode"] = "99"
		row["ResponseDescription"] = er.Error()
		list[0] = row
		return list
	}
	/* E: INSERT HEADER */

	sql := "INSERT INTO tblfcossppm(ppm_id, oss_id, business_id, company_id, institute_id, reg_status, ppm_no, ppm_date, ppm_sign_name, ppm_sign_nip, ppm_sign_date, ppm_sign_position, ppm_sign_ga, ppm_note, ppm_status, f_checklist, created_by, created_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, 0, ?, 'ServiceEngine', NOW())"
	res, err := o.Raw(sql, generateCode("PPM"), oss_id, Data.BusinessID, Data.CompanyID, Data.InstituteID, Data.RegStatus, Data.PermitNo, Data.PermitDate, Data.PermitSignName, Data.PermitSignNIP, Data.PermitSignDate, Data.PermitSignPosition, Data.PermitSignGA, Data.PermitNote, Data.FlagChecklist).Exec()

	//var w io.Writer
	//orm.DebugLog = orm.NewLog(w)

	if err == nil {
		num, _ := res.RowsAffected()

		if num > 0 {
			row["ResponseCode"] = "00"
			row["ResponseDescription"] = "Success"
			row["OSSID"] = oss_id
			list[0] = row
			return list
		} else {
			row["ResponseCode"] = "27"
			row["ResponseDescription"] = "Failed Insert"
			list[0] = row
			return list
		}

	} else {
		row["ResponseCode"] = "99"
		row["ResponseDescription"] = "Exception Error"
		list[0] = row
		return list
	}
}

func AddOnes(Data PPM) []map[string]interface{} {

	list := make([]map[string]interface{}, 1)
	row := make(map[string]interface{})

	/* S: INSERT PPM */
	if Data.PPM.no_ppm != "" {

		o := orm.NewOrm()

		sql := "UPDATE tblfcossppm SET ppm_no = ?, ppm_date = ?, ppm_sign_name = ?, ppm_sign_date = ?, ppm_sign_position = ?, ppm_sign_ga = ? WHERE ppm_id = ?"
		res, err := o.Raw(sql, Data.PPM.no_ppm, Data.PPM.tgl_ppm, Data.PPM.nama_ttd_ppm, Data.PPM.tgl_ttd_ppm, Data.PPM.jabatan_ttd_ppm, Data.PPM.penerbit, Data.PPM.IDip).Exec()

		if err == nil {
			num, _ := res.RowsAffected()
			if num <= 0 {

				row["ResponseCode"] = "27"
				row["ResponseDescription"] = "Failed Insert"
				list[0] = row
				return list

			} else {

				row["ResponseCode"] = "00"
				row["ResponseDescription"] = "Success"
				list[0] = row
				return list

			}
		}
	} else {
		bol, rest := InsertMaster(Data)

		if bol == false {
			return rest
		} else {
			row["ResponseCode"] = "00"
			row["ResponseDescription"] = "Success"
			list[0] = row
			return list
		}
	}
	/* E: INSERT PPM */

	row["ResponseCode"] = "98"
	row["ResponseDescription"] = "Other Exception"
	list[0] = row
	return list
}

func GetOne(ObjectId string) []map[string]interface{} {

	u := orm.NewOrm()
	var maps []orm.Params

	row := make(map[string]interface{})

	num, err := u.Raw("SELECT A.ppm_id, A.oss_id, A.company_id, B.npwp `comp_npwp`, B.name `comp_name`, C.institute_name `inst_name`, CASE A.ppm_status WHEN 0 THEN 'DRAFT'  WHEN 1 THEN 'PROCESS' WHEN 2 THEN 'APPROVED' WHEN 9 THEN 'REJECTED' END `status`, A.ppm_no, A.ppm_date, A.ppm_sign_name, A.ppm_sign_nip, A.ppm_sign_date, A.ppm_sign_position, D.ga_name, A.ppm_note, CASE A.f_checklist WHEN 0 THEN 'FALSE' WHEN 1 THEN 'TRUE' END `f_checklist`, A.created_date, A.updated_date FROM tblfcossppm A LEFT JOIN tbldmcompany B ON B.company_id = A.company_id LEFT JOIN tbldminstitute C ON C.institute_id = A.institute_id LEFT JOIN tbldmga D ON D.ga_id = A.ppm_sign_ga WHERE A.ppm_id = ?", ObjectId).Values(&maps)
	list := make([]map[string]interface{}, len(maps))

	if err != nil {
		list := make([]map[string]interface{}, 1)

		row["ResponseCode"] = "99"
		row["ResponseDescription"] = "Exception Error"
		list[0] = row
		return list
	}

	if num <= 0 {
		list := make([]map[string]interface{}, 1)
		row["ResponseCode"] = "01"
		row["ResponseDescription"] = "ObjectId Not Exist"

		list[0] = row
		return list

	} else {

		for k, v := range maps {
			row := make(map[string]interface{})

			row["OSSID"] = v["oss_id"]
			row["ID"] = v["ppm_id"]
			row["CompanyID"] = v["company_id"]
			row["CompanyNPWP"] = v["comp_npwp"]
			row["CompanyName"] = v["comp_name"]
			row["InstituteName"] = v["inst_name"]
			row["Status"] = v["status"]
			row["PermitNo"] = v["company_id"]
			row["PermitDate"] = v["ppm_no"]
			row["PermitNote"] = v["ppm_note"]
			row["PermitSignName"] = v["ppm_sign_name"]
			row["PermitSignNIP"] = v["ppm_sign_nip"]
			row["PermitSignGA"] = v["ga_name"]
			row["PermitSignDate"] = v["ppm_sign_date"]
			row["PermitSignPosition"] = v["ppm_sign_position"]
			row["PermitChecklist"] = v["f_checklist"]
			row["DateCreated"] = v["created_date"]
			row["DateUpdated"] = v["updated_date"]

			list[k] = row
		}

	}

	return list
}

func GetAll() []map[string]interface{} {

	u := orm.NewOrm()
	var maps []orm.Params

	row := make(map[string]interface{})

	_, err := u.Raw("SELECT A.ppm_id, A.oss_id, A.company_id, B.npwp `comp_npwp`, B.name `comp_name`, C.institute_name `inst_name`, CASE A.ppm_status WHEN 0 THEN 'DRAFT'  WHEN 1 THEN 'PROCESS' WHEN 2 THEN 'APPROVED' WHEN 9 THEN 'REJECTED' END `status`, A.ppm_no, A.ppm_date, A.ppm_sign_name, A.ppm_sign_nip, A.ppm_sign_date, A.ppm_sign_position, D.ga_name, A.ppm_note, CASE A.f_checklist WHEN 0 THEN 'FALSE' WHEN 1 THEN 'TRUE' END `f_checklist`, A.created_date, A.updated_date FROM tblfcossppm A LEFT JOIN tbldmcompany B ON B.company_id = A.company_id LEFT JOIN tbldminstitute C ON C.institute_id = A.institute_id LEFT JOIN tbldmga D ON D.ga_id = A.ppm_sign_ga").Values(&maps)

	if err != nil {
		list := make([]map[string]interface{}, 1)
		row["ResponseCode"] = "99"
		row["ResponseDescription"] = "Exeption Error"
		list[0] = row

		return list
	}

	list := make([]map[string]interface{}, len(maps))
	for k, v := range maps {
		row := make(map[string]interface{})
		row["OSSID"] = v["oss_id"]
		row["ID"] = v["ppm_id"]
		row["CompanyID"] = v["company_id"]
		row["CompanyNPWP"] = v["comp_npwp"]
		row["CompanyName"] = v["comp_name"]
		row["InstituteName"] = v["inst_name"]
		row["Status"] = v["status"]
		row["PermitNo"] = v["ppm_no"]
		row["PermitDate"] = v["ppm_date"]
		row["PermitNote"] = v["ppm_note"]
		row["PermitSignName"] = v["ppm_sign_name"]
		row["PermitSignNIP"] = v["ppm_sign_nip"]
		row["PermitSignGA"] = v["ga_name"]
		row["PermitSignDate"] = v["ppm_sign_date"]
		row["PermitSignPosition"] = v["ppm_sign_position"]
		row["PermitChecklist"] = v["f_checklist"]
		row["DateCreated"] = v["created_date"]
		row["DateUpdated"] = v["updated_date"]

		list[k] = row
	}

	return list
}

func Update(Oss_ID string, Data Object) []map[string]interface{} {
	list := make([]map[string]interface{}, 1)
	row := make(map[string]interface{})

	o := orm.NewOrm()
	sql := "UPDATE tblfcossppm SET ppm_no = ?, ppm_date = ?, ppm_note = ?, ppm_sign_name = ?, ppm_sign_nip = ?, ppm_sign_ga = ?, ppm_sign_date = ?, ppm_sign_position = ? WHERE ppm_id = ? AND oss_id = ?"

	res, err := o.Raw(sql, Data.PermitNo, Data.PermitDate, Data.PermitNote, Data.PermitSignName, Data.PermitSignNIP, Data.PermitSignGA, Data.PermitSignDate, Data.PermitSignPosition, Data.PermitID, Oss_ID).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num > 0 {
			row["ResponseCode"] = "00"
			row["ResponseDescription"] = "Success"
			list[0] = row
		} else {
			row["ResponseCode"] = "01"
			row["ResponseDescription"] = "Data Not Found"
			list[0] = row
		}

	} else {
		row["ResponseCode"] = "99"
		row["ResponseDescription"] = "Exception Error"
		list[0] = row
	}

	return list
}

func Delete(ObjectId string) []map[string]interface{} {

	o := orm.NewOrm()
	list := make([]map[string]interface{}, 1)
	row := make(map[string]interface{})

	sql := "DELETE FROM tblfcossppm WHERE ppm_id = ?"
	res, err := o.Raw(sql, ObjectId).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		if num > 0 {
			row["ResponseCode"] = "00"
			row["ResponseDescription"] = "Success"
			list[0] = row
		} else {
			row["ResponseCode"] = "01"
			row["ResponseDescription"] = "Data Not Found"
			list[0] = row
		}
	} else {
		row["ResponseCode"] = "27"
		row["ResponseDescription"] = "Failed Delete Data"
		list[0] = row
	}

	return list
}

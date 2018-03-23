package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

func InsertTDP(Object Object) (err error) {

	o := orm.NewOrm()
	var maps []orm.Params
	num, _ := o.Raw("SELECT company_id FROM tblfcosstdp WHERE oss_id = ?", Object.Oss_id).Values(&maps)

	if num > 0 {
		//UPDATE

		sql := " UPDATE institute_id = ?, reg_status = ?, kbli_id = ?, business_type = ?, date_activities = ?, main_activities = ?, other_activities = ?, patent_holder = ?, copyright_holder = ?, main_comodity = ?, agenda_no = ? WHERE company_id = ? AND oss_id = ?"
		res, err := o.Raw(sql, Object.Institute_id, Object.Reg_status, Object.Kbli_id, Object.Business_type, Object.Date_activities, Object.Main_activities, Object.Other_activities,
			Object.Patent_holder, Object.Copyright_holder, Object.Main_comodity, Object.Agenda_no, Object.Company_id, Object.Oss_id).Exec()

		if err == nil {
			num, _ := res.RowsAffected()

			if num > 0 {
				return nil
			} else {
				return errors.New("No Data Will Be Update")
			}
		}

	} else {

		//INSERT
		sql := `INSERT INTO tbldmcompany(tdp_id, oss_id, company_id, institute_id, reg_status, kbli_id, business_type, date_activities, main_activities, other_activities, patent_holder, copyright_holder, main_comodity, agenda_no) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		res, err := o.Raw(sql, generateCode("TDP"), Object.Oss_id, Object.Company_id, Object.Institute_id, Object.Reg_status, Object.Kbli_id, Object.Business_type,
			Object.Date_activities, Object.Main_activities, Object.Other_activities, Object.Patent_holder, Object.Copyright_holder, Object.Main_comodity, Object.Agenda_no).Exec()

		if err == nil {
			num, _ := res.RowsAffected()
			if num > 0 {
				return nil
			} else {
				return errors.New("Failed Insert Data")
			}
		}

	}

	return errors.New(err.Error())
}

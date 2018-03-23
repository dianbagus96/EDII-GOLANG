package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

func InsertTDP(Object Object) (err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	/*
		if Object.Tdp_no != "" {
			/*sql := `UPDATE tblfcosstdp SET tdp_no = ?, tdp_date = ?, tdp_valid_date = ?, tdp_sign_name = ?, tdp_sign_nip = ?, tdp_sign_date = ?, tdp_sign_position = ?, tdp_sign_ga = ?, tdp_note = ? WHERE npwp = ? AND oss_id = ?`
			res, err := o.Raw(sql, Object.Tdp_no, Object.Tdp_date, Object.Tdp_valid_date, Object.Tdp_sign_name, Object.Tdp_sign_nip, Object.Tdp_sign_date, Object.Tdp_sign_position, Object.Tdp_sign_ga, Object.Tdp_sign_note, Object.Npwp, Object.Oss_id).Exec()
			if err == nil {
				num, _ := res.RowsAffected()
				if num > 0 {
					return nil
				} else {
					return errors.New("No Data Will Be Update")
				}
			}ALTER TABLE sicantik.tmpermohonan ADD f_oss2 char(1) DEFAULT 0 NULL ;

		} else {*/
	num, _ := o.Raw("SELECT company_name FROM tblfcosstdp WHERE tdp_no = ?", Object.Tdp_no).Values(&maps)
	if num > 0 {
		//UPDATE
		sql := `UPDATE tblfcosstdp SET institute_id = ?, reg_status = ?, kbli_id = ?, business_type = ?, date_activities = ?, main_activities = ?, main_activities = ?, patent_holder = ?, copyright_holder = ?, main_comodity = ?, agenda_no = ?, company_name = ? ,
				tdp_date = ?, tdp_valid_date = ?, tdp_sign_name = ?, tdp_sign_nip = ?, tdp_sign_date = ?, tdp_sign_position = ?, tdp_sign_ga = ?, tdp_note = ?
				WHERE tdp_no = ?`
		res, err := o.Raw(sql, Object.Institute_id, Object.Reg_status, Object.Kbli_id, Object.Business_type, Object.Date_activities, Object.Main_activities, Object.Other_activities, Object.Patent_holder, Object.Copyright_holder, Object.Main_comodity, Object.Agenda_no, Object.Company_name,
			Object.Tdp_date, Object.Tdp_valid_date, Object.Tdp_sign_name, Object.Tdp_sign_nip, Object.Tdp_sign_date, Object.Tdp_sign_position, Object.Tdp_sign_ga, Object.Tdp_sign_note,
			Object.Tdp_no).Exec()
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
		sql := `INSERT INTO 
		tblfcosstdp(company_type, company_addr, tdp_id, oss_id, company_name, institute_id, reg_status, kbli_id, business_type, date_activities, main_activities, other_activies, patent_holder, copyright_holder, main_comodity, agenda_no, tdp_date, tdp_valid_date, tdp_sign_name, tdp_sign_nip, tdp_sign_date, tdp_sign_position, tdp_sign_ga, tdp_note, tdp_no) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		res, err := o.Raw(sql, `Perseroan Terbatas`, Object.Comapny_Addr, generateCode("TDP"), Object.Oss_id, Object.Company_name, Object.Institute_id, Object.Reg_status, Object.Kbli_id, Object.Business_type, Object.Date_activities, Object.Main_activities, Object.Other_activities, Object.Patent_holder, Object.Copyright_holder, Object.Main_comodity, Object.Agenda_no,
			Object.Tdp_date, Object.Tdp_valid_date, Object.Tdp_sign_name, Object.Tdp_sign_nip, Object.Tdp_sign_date, Object.Tdp_sign_position, Object.Tdp_sign_ga, Object.Tdp_sign_note, Object.Tdp_no).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			if num > 0 {
				return nil
			} else {
				return errors.New("Failed Insert Data")
			}
		}
	}
	//}
	return errors.New(err.Error())
}

func InsertIMB(Object Imb) (err error) {
	o := orm.NewOrm()
	var maps []orm.Params

	num, _ := o.Raw("SELECT company_name FROM tblfcossimb WHERE imb_no = ?", Object.Tdp_no).Values(&maps)
	if num > 0 {
		//UPDATE
		sql := `UPDATE tblfcossimb SET  
			oss_id = ?, 
			company_name = ?, 
			company_addr = ?, 
			imb_date = ?, 
			imb_sign_name = ?, 
			imb_sign_nip = ?, 
			imb_sign_date = ?, 
			imb_sign_position = ?, 
			imb_note = ?, 
		WHERE imb_no = ?`
		res, err := o.Raw(sql, Object.Oss_id, Object.Company_name, Object.Comapny_Addr, Object.Tdp_date, Object.Tdp_sign_name, Object.Tdp_sign_nip, Object.Tdp_sign_date, Object.Tdp_sign_position, Object.Tdp_sign_note).Exec()
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
		sql := `INSERT INTO 
		tblfcossimb(
			imb_id, 
			oss_id, 
			company_name, 
			company_type, 
			company_addr, 
			institute_id, 
			imb_no, 
			imb_date, 
			imb_sign_name, 
			imb_sign_nip, 
			imb_sign_date, 
			imb_sign_position, 
			imb_note, 
			created_date) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NOW())`
		res, err := o.Raw(sql,
			generateCode("IMB"),
			Object.Oss_id,
			Object.Company_name,
			`Perseroan Terbatas`,
			Object.Comapny_Addr,
			Object.Institute_id,
			Object.Tdp_no,
			Object.Tdp_date,
			Object.Tdp_sign_name,
			Object.Tdp_sign_nip,
			Object.Tdp_sign_date,
			Object.Tdp_sign_position,
			Object.Tdp_sign_note).Exec()
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

package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

func InsertDmCompany(Object PPM) (err error) {

	o := orm.NewOrm()
	var maps []orm.Params
	num, _ := o.Raw("SELECT company_id FROM tbldmcompany WHERE company_id = ?", Object.PPM.IDPerusahaan).Values(&maps)

	//return errors.New(strconv.FormatInt(num, 10))

	if num > 0 {
		//UPDATE
		sql := " UPDATE tbldmcompany SET legal_status =  ?, invest_status =  ?, name =  ?, npwp =  ?, prop =  ?, kab =  ?, kec =  ?, kel =  ?, address =  ?, rt =  ?, rw =  ?, zip_id =  ?, telp =  ?, email =  ?, pic_identity_no =  ?, pic_name =  ?, pic_address =  ?, pic_position =  ?, pic_telp =  ?, pic_email =  ? WHERE company_id = ?"
		res, err := o.Raw(sql, Object.PPM.Status_badan, Object.PPM.Status_modal, Object.PPM.Nama_perusahaan, Object.PPM.Npwp_perusahaan, Object.PPM.Prop_perusahaan, Object.PPM.Kab_perusahaan, Object.PPM.Kec_perusahaan, Object.PPM.Kel_perusahaan, Object.PPM.Alamat_perusahaan, Object.PPM.Rt_perusahaan, Object.PPM.Rw_perusahaan, Object.PPM.Kodepos_perusahaan, Object.PPM.Telp_perusahaan, Object.PPM.Email_perusahaan, Object.PPM.Npwp_pic, Object.PPM.Nama_pic, Object.PPM.Alamat_pic, Object.PPM.Jabatan_pic, Object.PPM.Telp_pic, Object.PPM.Email_pic, Object.PPM.IDPerusahaan).Exec()

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
		sql := `INSERT INTO tbldmcompany(company_id, legal_status, invest_status, name, npwp, prop, kab, kec, kel, address, rt, rw, zip_id, telp, email, pic_identity_no, pic_name, pic_address, pic_position, pic_telp, pic_email) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )`
		res, err := o.Raw(sql, Object.PPM.IDPerusahaan, Object.PPM.Status_badan, Object.PPM.Status_modal, Object.PPM.Nama_perusahaan, Object.PPM.Npwp_perusahaan, Object.PPM.Prop_perusahaan, Object.PPM.Kab_perusahaan, Object.PPM.Kec_perusahaan, Object.PPM.Kel_perusahaan, Object.PPM.Alamat_perusahaan, Object.PPM.Rt_perusahaan, Object.PPM.Rw_perusahaan, Object.PPM.Kodepos_perusahaan, Object.PPM.Telp_perusahaan, Object.PPM.Email_perusahaan, Object.PPM.Npwp_pic, Object.PPM.Nama_pic, Object.PPM.Alamat_pic, Object.PPM.Jabatan_pic, Object.PPM.Telp_pic, Object.PPM.Email_pic).Exec()

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

func InsertCompanyDeed(Object PPM) (err error) {

	o := orm.NewOrm()
	var maps []orm.Params
	num, _ := o.Raw("SELECT deed_no FROM tblfccompanydeed WHERE deed_no = ?", Object.PPM.No_akta).Values(&maps)

	//return errors.New(strconv.FormatInt(num, 10))

	if num > 0 {
		//UPDATE
		sql := " UPDATE tblfccompanydeed SET notaris_name =  ?, deed_date =  ? WHERE deed_no = ?"
		res, err := o.Raw(sql, Object.PPM.Nama_notaris, Object.PPM.Tgl_akta, Object.PPM.No_akta).Exec()

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
		sql := `INSERT INTO tblfccompanydeed(notaris_name, deed_no, deed_date) VALUES (?, ?, ?)`
		res, err := o.Raw(sql, Object.PPM.Nama_notaris, Object.PPM.No_akta, Object.PPM.Tgl_akta).Exec()

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

func InsertCompanyBusiness(Object PPM) (err error) {

	o := orm.NewOrm()
	var maps []orm.Params
	num, _ := o.Raw("SELECT deed_no FROM tblfccompanybusiness WHERE deed_no = ?", Object.PPM.No_akta).Values(&maps)

	//return errors.New(strconv.FormatInt(num, 10))

	if num > 0 {
		//UPDATE
		sql := " UPDATE tblfccompanybusiness SET project_sch =  ?, line_business =  ?, kbli_id =  ?, prop =  ?, kab =  ?, address =  ?, land_type =  ?, imp_total =  ?, fmp_total =  ?, purchase_land =  ?, machine_equip =  ?, machinge_equip_us =  ?, working_capital =  ?, total_invest =  ?, total_asset =  ?, total_omset =  ? WHERE deed_no = ?"
		res, err := o.Raw(sql, Object.PPM.Jadwal_proyek, Object.PPM.Bidang_usaha, Object.PPM.Kbli, Object.PPM.Prop_bidang_usaha, Object.PPM.Alamat_bidang_usaha, Object.PPM.Status_tanah_bidang_usaha, Object.PPM.Jumlah_tki, Object.PPM.Jumlah_tka, Object.PPM.Nilai_tanah, Object.PPM.Nilai_mesing_peralatan, Object.PPM.Nilai_mesin_peralatan_us, Object.PPM.Modal_kerja, Object.PPM.Total_investasi, Object.PPM.Total_aset, Object.PPM.Total_omset, Object.PPM.No_akta).Exec()

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
		sql := `INSERT INTO tblfccompanybusiness(project_sch, line_business, kbli_id, prop, kab, address, land_type, imp_total, fmp_total, purchase_land, machine_equip, machinge_equip_us, working_capital, total_invest, total_asset, total_omset) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		res, err := o.Raw(sql, Object.PPM.Jadwal_proyek, Object.PPM.Bidang_usaha, Object.PPM.Kbli, Object.PPM.Prop_bidang_usaha, Object.PPM.Alamat_bidang_usaha, Object.PPM.Status_tanah_bidang_usaha, Object.PPM.Jumlah_tki, Object.PPM.Jumlah_tka, Object.PPM.Nilai_tanah, Object.PPM.Nilai_mesing_peralatan, Object.PPM.Nilai_mesin_peralatan_us, Object.PPM.Modal_kerja, Object.PPM.Total_investasi, Object.PPM.Total_aset, Object.PPM.Total_omset).Exec()

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

func InsertCompanyShare(Object PPM) (err error, shareID string) {

	o := orm.NewOrm()
	var maps []orm.Params
	num, _ := o.Raw("SELECT company_id FROM tblfccompanyshare WHERE company_id = ?", Object.PPM.IDPerusahaan).Values(&maps)

	//return errors.New(strconv.FormatInt(num, 10))

	if num > 0 {
		//UPDATE
		sql := " UPDATE tblfccompanyshare SET number_of_share =  ?, value_of_share =  ?, total_foreign_share =  ?, total_percent_foreign =  ?, total_domestic_share =  ?, total_value_share =  ?, authorized_capital =  ?, issued_capital =  ?, equity =  ?, paid_up_capital =  ?, return_earning =  ?, foreign_loan =  ?, domestic_loan =  ?, total_capital =  ?, total_funds =  ? WHERE company_id = ?"
		res, err := o.Raw(sql, Object.PPM.Jumlah_saham_lembar, Object.PPM.Nilai_saham_lembar, Object.PPM.Total_saham_asing, Object.PPM.Total_persen_saham_asing, Object.PPM.Total_saham_lokal, Object.PPM.Total_saham, Object.PPM.Modal_dasar, Object.PPM.Modal_ditempatkan, Object.PPM.Ekuitas, Object.PPM.Modal_disetor, Object.PPM.Laba, Object.PPM.Pinjaman_asing, Object.PPM.Pinjaman_lokal, Object.PPM.Total_modal, Object.PPM.Total_dana, Object.PPM.IDPerusahaan).Exec()

		if err == nil {
			num, _ := res.RowsAffected()

			if num > 0 {
				return nil, ""
			} else {
				return errors.New("No Data Will Be Update"), ""
			}
		}

	} else {

		//INSERT
		sql := `INSERT	 INTO tblfccompanyshare(share_id, company_id, number_of_share, value_of_share, total_foreign_share, total_percent_foreign, total_domestic_share, total_value_share, authorized_capital, issued_capital, equity, paid_up_capital, return_earning, foreign_loan, domestic_loan, total_capital, total_funds) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		res, err := o.Raw(sql, generateCode("SHARE"), Object.PPM.IDPerusahaan, Object.PPM.Jumlah_saham_lembar, Object.PPM.Nilai_saham_lembar, Object.PPM.Total_saham_asing, Object.PPM.Total_persen_saham_asing, Object.PPM.Total_saham_lokal, Object.PPM.Total_saham, Object.PPM.Modal_dasar, Object.PPM.Modal_ditempatkan, Object.PPM.Ekuitas, Object.PPM.Modal_disetor, Object.PPM.Laba, Object.PPM.Pinjaman_asing, Object.PPM.Pinjaman_lokal, Object.PPM.Total_modal, Object.PPM.Total_dana).Exec()

		if err == nil {
			num, _ := res.RowsAffected()
			if num > 0 {
				return nil, ""
			} else {
				return errors.New("Failed Insert Data"), ""
			}
		}

	}

	return errors.New(err.Error()), ""
}

func InsertCompanyApplicant(Object PPM) (err error) {

	o := orm.NewOrm()
	var maps []orm.Params
	num, _ := o.Raw("SELECT company_id FROM tblfccompanyapplicant WHERE company_id = ? AND identity_no = ?", Object.PPM.IDPerusahaan, Object.PPM.No_id_pemohon).Values(&maps)

	//return errors.New(strconv.FormatInt(num, 10))

	if num > 0 {
		//UPDATE
		sql := " UPDATE tblfccompanyapplicant SET name =  ?, position =  ?, identity_type =  ?, identity_no =  ?, prop =  ?, kab =  ?, kec =  ?, kel =  ?, address =  ?, zip_id =  ?, telp =  ?, hp =  ?, email =  ? WHERE company_id = ? AND identity_no = ?"
		res, err := o.Raw(sql, Object.PPM.Nama_pemohon, Object.PPM.Jabatan_pemohon, Object.PPM.Jenis_id_pemohon, Object.PPM.No_id_pemohon, Object.PPM.Prop_pemohon, Object.PPM.Kab_pemohon, Object.PPM.Kec_pemohon, Object.PPM.Kel_pemohon, Object.PPM.Alamat_pemohon, Object.PPM.Kodepos_pemohon, Object.PPM.Telp_pemohon, Object.PPM.Hp_pemohon, Object.PPM.Email_pemohon, Object.PPM.IDPerusahaan, Object.PPM.No_id_pemohon).Exec()

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
		sql := `INSERT INTO tblfccompanyapplicant(applicant_id, company_id, name, position, identity_type, identity_no, prop, kab, kec, kel, address, zip_id, telp, hp, email) 
				VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
		res, err := o.Raw(sql, generateCode("APL"), Object.PPM.IDPerusahaan, Object.PPM.Nama_pemohon, Object.PPM.Jabatan_pemohon, Object.PPM.Jenis_id_pemohon, Object.PPM.No_id_pemohon, Object.PPM.Prop_pemohon, Object.PPM.Kab_pemohon, Object.PPM.Kec_pemohon, Object.PPM.Kel_pemohon, Object.PPM.Alamat_pemohon, Object.PPM.Kodepos_pemohon, Object.PPM.Telp_pemohon, Object.PPM.Hp_pemohon, Object.PPM.Email_pemohon).Exec()

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

func InsertShareholders(Data PPM, shareID string) (err error) {

	o := orm.NewOrm()
	//var maps []orm.Params

	var count_err int = 0
	for _, v := range Data.PPM.Pemegang_saham {
		var maps []orm.Params
		num, _ := o.Raw("SELECT identity_no FROM tblfccompanyshareholder WHERE identity_no = ? ", v.No_id_pemegang_Saham).Values(&maps)

		if num > 0 {
			sql := `UPDATE tblfccompanyshareholder SET shareholder_name =  ?, identity_type =  ?, identity_no =  ?, identity_valid =  ?, npwp =  ?, address =  ?, country =  ?, telp =  ?, fax =  ?, email =  ?, share_value =  ? WHERE identity_no = ? `
			_, err := o.Raw(sql, v.Nama_pemegang_saham, v.Jenis_id_pemegang_saham, v.No_id_pemegang_Saham, v.Masa_berlaku_id, v.Npwp_pemegang_saham, v.Alamat_pemegang_saham, v.Negara_pemegang_saham, v.Telp_pemegang_saham, v.Fax_pemegang_saham, v.Email_pemegang_saham, v.Nilai_saham, v.No_id_pemegang_Saham).Exec()
			if err != nil {
				count_err++
			}
		} else {
			sql := `INSERT INTO 
			tblfccompanyshareholder(share_id, shareholder_name, identity_type, identity_no, identity_valid, npwp, address, country, telp, fax, email, share_value) 
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
			_, err = o.Raw(sql, shareID, v.Nama_pemegang_saham, v.Jenis_id_pemegang_saham, v.No_id_pemegang_Saham, v.Masa_berlaku_id, v.Npwp_pemegang_saham, v.Alamat_pemegang_saham, v.Negara_pemegang_saham, v.Telp_pemegang_saham, v.Fax_pemegang_saham, v.Email_pemegang_saham, v.Nilai_saham).Exec()
			if err != nil {
				count_err++
			}
		}

	}

	if count_err != 0 {
		return errors.New("Failed Insert Data")
	}

	return nil
}

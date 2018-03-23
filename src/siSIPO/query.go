package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var jsonString string

func buildParamSaham(company_id string) string {

	var shm []orm.Params
	var i int64
	o := orm.NewOrm()
	num, err := o.Raw("SELECT shareholder_name, address, telp, country, npwp, share_value FROM tblfccompanyshareholder A LEFT JOIN tblfccompanyshare B ON B.share_id = A.share_id WHERE B.company_id = '" + company_id + "'").Values(&shm)
	var loop string

	for i = 0; i <= num-1; i++ {
		shareholder_name := shm[i]["shareholder_name"].(string)
		address := shm[i]["address"].(string)
		telp := shm[i]["telp"].(string)
		country := shm[i]["country"].(string)
		npwp := shm[i]["npwp"].(string)
		share_value := shm[i]["share_value"].(string)

		loop = loop + `"loop": {
			"nama": "` + shareholder_name + `",
			"alamat": "` + address + `",
			"kode_pos": "-",
			"telepon": "` + telp + `",
			"negara": "` + country + `",
			"npwp": "` + npwp + `",
			"jumlah_saham": "` + share_value + `",
			"jumlah_modal": "-"
		},`
	}

	if err != nil {
		panic(err.Error())
	}

	jsonString = `"saham": {
                ` + loop + `
			  },`

	return jsonString
}

func buildParamCabang() string {
	var loop string
	var i int64

	for i = 0; i <= 1; i++ {
		loop = loop + `"loop": {
			"nama": "-",
			"no_tdp": "-",
			"alamat": "-",
			"idprop": "-",
			"idkab": "-",
			"kode_pos": "-",
			"telepon": "-",
			"status_perusahaan": "-",
			"jenis_perusahaan": "-"
		  }`
	}
	jsonString = `"cabang": {
		` + loop + `
	  }`

	return jsonString
}

func buildParamIjin() string {
	jsonString = `"ijin_lain": {
		"loop": {
		  "id_ijin": "-",
		  "nomor": "-",
		  "penerbit": "-",
		  "tanggal": "-",
		  "berlaku": "-"
		}
	  },`

	return jsonString
}

func buildParamPimpinan(company_id string) string {

	var shm []orm.Params
	var i int64
	o := orm.NewOrm()
	num, err := o.Raw("SELECT shareholder_name, shareholder_position, address, telp, country, npwp, share_value FROM tblfccompanyshareholder A LEFT JOIN tblfccompanyshare B ON B.share_id = A.share_id WHERE B.company_id = '" + company_id + "'").Values(&shm)
	var loop string

	if err != nil {
		panic(err.Error())
	}

	for i = 0; i <= num-1; i++ {
		shareholder_name := shm[i]["shareholder_name"].(string)
		shareholder_position := shm[i]["shareholder_position"].(string)
		address := shm[i]["address"].(string)
		telp := shm[i]["telp"].(string)
		country := shm[i]["country"].(string)
		share_value := shm[i]["share_value"].(string)

		jabatan := buildParamJabatan("loop")

		loop = loop + `"loop": {
			"nama": "` + shareholder_name + `",
			"jabatan": "` + shareholder_position + `",
			"tempat_lahir": "-",
			"tanggal_lahir": "-",
			"alamat": "` + address + `",
			"kode_pos": "-",
			"telepon": "` + telp + `",
			"negara": "` + country + `",
			"tanggal_jabatan": "-",
			"jumlah_saham": "` + share_value + `",
			"jumlah_modal": "-",
			"jabatan_lain": {
			  ` + jabatan + `
			}
		  },`

	}

	jsonString = loop

	return jsonString
}

func buildParamJabatan(jabatan string) string {
	jsonString = `"loop": {
		"nama": "-",
		"jabatan": "-",
		"alamat": "-",
		"kode_pos": "-",
		"telepon": "-",
		"tanggal_jabatan": "-"
	  }`

	return jsonString
}

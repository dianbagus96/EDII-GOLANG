package main

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/go-resty/resty"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jasonlvhit/gocron"
)

func init() {
	//orm.DRMySQL
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "ossuser:db0ss2017@tcp(10.1.2.87:3306)/oss?charset=utf8")
}

func task() {

	var tdp []orm.Params
	var i int64
	o := orm.NewOrm()
	num, err := o.Raw(`
        SELECT F.*, E.notaris_name,E.deed_type, D.deed_id, D.legal_no, D.legal_date, C.bank_id, C.bank_name, A.*,IFNULL(B.company_type,'') 'comp_type', IFNULL(B.company_status,'') 'comp_status',IFNULL(B.name,'') 'comp_name', IFNULL(B.address,'') 'comp_addr', IFNULL(B.npwp,'') 'comp_npwp', IFNULL(B.telp,'') 'comp_telp', IFNULL(B.fax,'') 'comp_fax', IFNULL(B.kab,'') 'comp_kab', IFNULL(B.prop,'') 'comp_prop', IFNULL(B.kec,'') 'comp_kec',IFNULL(B.kel,'') 'comp_kel',IFNULL(B.zip_id,'') 'comp_zip_id', IFNULL(B.pic_name,'') 'nama_pic', IFNULL(B.pic_address,'') 'alamat_pic', IFNULL(B.pic_identity_no,'') 'pic_id_no', IFNULL(B.pic_position,'') 'posisi_pic', IFNULL(B.pic_email,'') 'email_pic' 
        FROM tblfcosstdp A 
        LEFT JOIN tbldmcompany B 
          ON B.company_id = A.company_id
        LEFT JOIN (SELECT A.company_id, A.bank_id, B.bank_name FROM tblfccompanybank A LEFT JOIN tbldmbank B ON B.bank_id = A.bank_id LIMIT 1) C 
          ON C.company_id = A.company_id
        LEFT JOIN (SELECT company_id, deed_id , legal_no, legal_date FROM tblfccompanylegal ORDER BY created_date DESC LIMIT 1) D 
          ON D.company_id = A.company_id
        LEFT JOIN tblfccompanydeed E 
          ON E.deed_id = D.deed_id
        LEFT JOIN tblfccompanyshare F
          ON F.company_id = A.company_id
        WHERE A.tdp_status= 0
  `).Values(&tdp)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Total Data = ", num)

	resty.SetHeaders(map[string]string{
		"X-App-Type":                   "TDP",
		"x-api-key":                    "SRhz5nFeTCy1YliItAGugjaJbPKosBvrfq860Mkd",
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "x-app-type",
	})
	//.SetHeader(`x-api-key`, `MQbTaGOivtnc1y4ldD2JVf3EZWUhKsBYe6gHw5Lj`).SetHeader(`x-app-type`, `TDP`)
	for i = 0; i <= num-1; i++ {
		//var ossID, comp_status, npwp, comp_name, comp_address, comp_prop, comp_kab, comp_kec, comp_kel, zip_id, comp_telp, comp_fax, pic_name, pic_address, pic_id_no, pic_email string
		ossID := tdp[i]["oss_id"].(string)
		comp_status := tdp[i]["comp_status"].(string)
		npwp := tdp[i]["comp_npwp"].(string)
		comp_name := tdp[i]["comp_name"].(string)
		comp_address := tdp[i]["comp_addr"].(string)
		comp_prop := tdp[i]["comp_prop"].(string)
		comp_kab := tdp[i]["comp_kab"].(string)
		comp_kec := tdp[i]["comp_kec"].(string)
		comp_kel := tdp[i]["comp_kel"].(string)
		zip_id := tdp[i]["comp_zip_id"].(string)
		comp_telp := tdp[i]["comp_telp"].(string)
		comp_fax := tdp[i]["comp_fax"].(string)
		pic_name := tdp[i]["nama_pic"].(string)
		pic_address := tdp[i]["alamat_pic"].(string)
		pic_id_no := tdp[i]["pic_id_no"].(string)
		pic_email := tdp[i]["email_pic"].(string)
		//bank_id := tdp[i]["bank_id"].(string)
		bank_name := tdp[i]["bank_name"].(string)
		notaris_name := tdp[i]["notaris_name"].(string)
		legal_no := tdp[i]["legal_no"].(string)
		legal_date := tdp[i]["legal_date"].(string)
		authorized_capital := tdp[i]["authorized_capital"].(string)
		value_of_share := tdp[i]["value_of_share"].(string)
		total_value_share := tdp[i]["total_value_share"].(string)

		saham := buildParamSaham(tdp[i]["company_id"].(string))
		cabang := buildParamCabang()
		ijin_lain := buildParamIjin()
		pimpinan := buildParamPimpinan(tdp[i]["company_id"].(string))

		paramJson := `{
            "binus_tdp": {
              "permohonan": {
                "oss_id": "` + ossID + `",
                "jenis_permohonan": "1",
                "pembaharuan": "0",
                "no_ijin": "",
                "tgl_ijin": "",
                "tgl_akhir": "",
                "kbli": "0",
                "pokok": "-",
                "nama_ttd": "",
                "jabat_ttd": "",
                "nip_ttd": ""
              },
              "perusahaan": {
                "bentuk_usaha": "5",
                "status_perusahaan": "` + comp_status + `",
                "npwp": "` + npwp + `",
                "nama": "` + comp_name + `",
                "alamat": "` + comp_address + `",
                "idprop": "` + comp_prop + `",
                "idkab": "` + comp_kab + `",
                "idkec": "` + comp_kec + `",
                "idkel": "` + comp_kel + `",
                "kode_pos": "` + zip_id + `",
                "telepon": "` + comp_telp + `",
                "fax": "` + comp_fax + `"
              },
              "kegiatan_perusahaan": {
                "lokasi_produksi": "-",
                "prop_produksi": "-",
                "kab_produksi": "-",
                "id_tanam_modal": "-",
                "tgl_pendirian": "-",
                "tgl_kegiatan": "-",
                "lama_berdiri": "-",
                "id_kerjasama": "-",
                "nama_merek_dagang": "-",
                "no_merek_dagang": "-",
                "pemegang_paten": "-",
                "no_paten": "-",
                "pemegang_cipta": "-",
                "no_cipta": "-",
                "jumlah_wni": "-",
                "jumlah_wna": "-",
                "jenis_perusahaan": "-",
                "lain_1": "-",
                "lain_2": "-",
                "komoditi_pokok": "-",
                "komoditi_lain_1": "-",
                "komoditi_lain_2": "-",
                "omset": "-",
                "total_aset": "-",
                "id_mata_rantai": "-",
                "flag_mesin": "-",
                "kapasitas_terpasang": "-",
                "kapasitas_satuan": "-",
                "kapasitas_produksi": "-",
                "kapasitas_produksi_satuan": "-",
                "kandungan_lokal": "-",
                "kandungan_impor": "-",
                "id_pengecer": "-"
              },
              "koperasi": {
                "jenis_koperasi": "-",
                "jumlah_anggota_koperasi": "-",
                "simpanan_pokok": "-",
                "simpanan_wajib": "-",
                "dana_cadangan": "-",
                "hibah": "-",
                "pinjaman_anggota": "-",
                "pinjaman_koperasi": "-",
                "pinjaman_bank": "-",
                "pinjaman_lain": "-"
              },
              "perusahaan_induk": {
                "nama": "-",
                "no_tdp": "-",
                "alamat": "-",
                "idprop": "-",
                "idkab": "-",
                "idkec": "-",
                "idkel": "-"
              },
              "penanggung_jawab": {
                "nama": "` + pic_name + `",
                "alamat": "` + pic_address + `",
                "tempat_lahir": "-",
                "tanggal_lahir": "-",
                "idprop": "-",
                "idkab": "-",
                "idkec": "-",
                "idkel": "-",
                "telepon": "-",
                "jenis_identitas": "-",
                "no_identitas": "` + pic_id_no + `",
                "negara": "-",
                "email": "` + pic_email + `"
              },
              "bank": {
                "bank_1": "` + bank_name + `",
                "bank_2": "-",
                "jumlah": "-"
              },
              "modal": {
                "modal_dasar": "` + authorized_capital + `",
                "modal_ditempatkan": "-",
                "modal_disetor": "-",
                "jumlah_saham": "` + total_value_share + `",
                "nilai_per_saham": "` + value_of_share + `",
                "modal_sekutu_aktif": "-",
                "modal_sekutu_pasif": "-"
              },
              "legalitas": {
                "no_pendirian": "` + legal_no + `",
                "tgl_pendirian": "` + legal_date + `",
                "alamat_pendirian": "-",
                "notaris_pendirian": "` + notaris_name + `",
                "telepon_pendirian": "-",
                "no_perubahan": "-",
                "tgl_perubahan": "-",
                "notaris_perubahan": "-",
                "tipe_pengesahan": "-",
                "no_pengesahan": "-",
                "tgl_pengesahan": "-",
                "no_perubahan_pengesahan": "-",
                "tgl_perubahan_pengesahan": "-",
                "no_perubahan_anggaran": "-",
                "tgl_perubahan_anggaran": "-",
                "no_perubahan_direksi": "-",
                "tgl_perubahan_direksi": "-"
              },
              ` + ijin_lain + `
              "pimpinan": {
                "jumlah_penanggung_jawab": "-",
                "jumlah_direktur": "-",
                "jumlah_komisaris": "-",
                "jumlah_pengawas": "-",
                "jumlah_sekutu_aktif": "-",
                "jumlah_sekutu_pasif": "-",
                "jumlah_sekutu_aktif_baru": "-",
                "jumlah_sekutu_pasif_baru": "-",
                "jumlah_pemegang_saham": "-",
                ` + pimpinan + `
              },
              ` + saham + cabang + `              
            }
          }`

		//panic(paramJson)

		resp, err := resty.R().SetBody(map[string]interface{}{"data": paramJson}).Post("http://hub.kemendag.go.id/index.php/oss_api/send_req")

		//http://103.29.187.96/android/oss/public/index.php/server/setsipo
		//http: //103.29.187.96/android/oss/public/index.php/server/setapi
		//http://hub.kemendag.go.id/index.php/oss_api/send_req

		// explore response object
		fmt.Printf("\nOSS ID: %v", ossID)
		fmt.Printf("\nError: %v", err)
		fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
		fmt.Printf("\nResponse Status: %v", resp.Status())
		fmt.Printf("\nResponse Time: %v", resp.Time())
		fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
		fmt.Printf("\nResponse Body: %v \n", resp) // or resp.String() or string(resp.Body())

		fmt.Println("-----------------------------------")
	}
	//}

	// if err = rows.Err(); err != nil {
	// 	panic(err.Error()) // proper error handling instead of panic in your app
	// }

	// GET request
	//http://103.29.187.96/android/oss/public/index.php/server/api_maste
	//resp, err := resty.R().Get("http://103.29.187.96/android/oss/public/index.php/server/penerbit?prop=31")

}

// muntuk =? untuk - home

/*
	List Code =>
	00 : Berhasil
	99 : Insert Data Failed
	88 : OSSID Sudah Ada
	27 : Post Data Not Found
*/

func main() {
	gocron.Every(30).Seconds().Do(task)
	_, time := gocron.NextRun()
	fmt.Println(time)
	<-gocron.Start()
}

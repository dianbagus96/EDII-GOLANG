package main

import (
	"database/sql"
	"fmt"

	"github.com/go-resty/resty"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "ossuser:db0ss2017@tcp(10.1.2.87:3306)/oss?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query(`SELECT api_id, oss_id, no_reff_bank, date_reff_bank, name, address, IFNULL('-', fax) fax, IFNULL('-', telp) telp FROM tblfcossapi A LEFT JOIN tbldmcompany B ON B.company_id = A.company_id WHERE api_status = 1`)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Setting a Proxy URL and Port
	resty.SetProxy("http://192.168.5.180:8013")

	// Fetch rows
	//for rows.Next() {
	var api_id, oss_id, no_reff_bank, name, date_reff_bank, address, fax, telp string
	rows.Scan(&api_id, &oss_id, &no_reff_bank, &date_reff_bank, &name, &address, &fax, &telp)
	fmt.Println(api_id, oss_id, no_reff_bank, name, address, fax, telp, date_reff_bank)

	//var api_id, oss_id, no_reff_bank, name, address, fax, telp string
	rows.Scan(&api_id, &oss_id, &no_reff_bank, &name, &address, &fax, &telp)
	fmt.Println(api_id, oss_id, no_reff_bank, name, address, fax, telp)

	resty.SetHeaders(map[string]string{
		`X-App-Type`: `API`,
		"x-api-key":  "MQbTaGOivtnc1y4ldD2JVf3EZWUhKsBYe6gHw5Lj",
	})

	//resty.SetHeaders(map[string]string{
	//		"Content-Type": "application/json",
	//		"Accept":       "application/json",
	//	})

	resp, err := resty.R().SetBody(map[string]interface{}{"data": `{
			"permohonan": {
			  "header": {
				"oss_id": "` + oss_id + `",
				"kantor_dinas": "78",
				"jenis_permohonan": "1",
				"jenis_api": "1",
				"id_bentuk_usaha": "1",
				"npwp_perusahaan": "123456789012345",
				"nama_perusahaan": "` + name + `",
				"alamat_perusahaan": "` + address + `",
				"telp_perusahaan": "` + telp + `",
				"fax_perusahaan": "` + fax + `",
				"id_kab": "0906",
				"id_prop": "09",
				"kd_pos": "12345",
				"no_surat": "No Surat",
				"tgl_surat": "2018-10-01",
				"no_aknot": "No Aknot",
				"no_surat_domisili": "No domisili",
				"no_tdp": "No Tdp",
				"tgl_awal_tdp": "2018-01-01",
				"tgl_akhir_tdp": "2020-01-01",
				"nama_ref_bank": "-",
				"no_ref_bank": "` + no_reff_bank + `",
				"no_izin_usaha": "No IU",
				"jenis_usaha": "Industri test",
				"dagangan_utama": "Penjualan barang test",
				"dagangan_lain": "Penjualan barang apapun"
			  },
			  "direksi_perusahaan": {
				"data": [
				  {
					"nama": "John Doe",
					"alamat": "Wisma Merdeka",
					"jabatan": "CTO",
					"tipe_identitas": "WNI",
					"no_identitas": "KTP",
					"npwp": "NPWP",
					"no_imta": "-"
				  },
				  {
					"nama": "Jane Doe",
					"alamat": "Wisma Merdeka",
					"jabatan": "MVP",
					"tipe_identitas": "WNI",
					"no_identitas": "KTP",
					"npwp": "NPWP",
					"no_imta": "-"
				  }
				]
			  }
			}
			}`}).
		Post("http://hub.kemendag.go.id/index.php/oss_api/send_req")
	//Post("http://103.29.187.96/android/oss/public/index.php/server/setapi")

	//http://103.29.187.96/android/oss/public/index.php/server/setsipo
	//http: //103.29.187.96/android/oss/public/index.php/server/setapi

	// explore response object
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
	fmt.Printf("\nResponse Body: %v \n", resp) // or resp.String() or string(resp.Body())

	fmt.Println("-----------------------------------")
	//}

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// GET request
	//http://103.29.187.96/android/oss/public/index.php/server/api_maste
	//resp, err := resty.R().Get("http://103.29.187.96/android/oss/public/index.php/server/penerbit?prop=31")

}

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

	rows, err := db.Query(`SELECT * FROM tblfcosstdp A LEFT JOIN tbldmcompany B ON B.company_id = A.company_id WHERE tdp_status = 1`)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Setting a Proxy URL and Port
	resty.SetProxy("http://192.168.5.180:8013")

	// Fetch rows
	for rows.Next() {
		var api_id, oss_id, no_reff_bank, name, date_reff_bank, address, fax, telp string
		rows.Scan(&api_id, &oss_id, &no_reff_bank, &date_reff_bank, &name, &address, &fax, &telp)
		fmt.Println(api_id, oss_id, no_reff_bank, name, address, fax, telp, date_reff_bank)

		var p_npwp = `1`
		var p_nama = `1`
		var p_alamat = `1`
		var pj_nama = `1`
		var pj_alamat = `1`

		//var api_id, oss_id, no_reff_bank, name, address, fax, telp string
		rows.Scan(&api_id, &oss_id, &no_reff_bank, &name, &address, &fax, &telp)
		fmt.Println(api_id, oss_id, no_reff_bank, name, address, fax, telp)

		fmt.Println("start app lho")

		resty.SetHeaders(map[string]string{
			"X-App-Type":                   "TDP",
			"x-api-key":                    "SRhz5nFeTCy1YliItAGugjaJbPKosBvrfq860Mkd",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "x-app-type",
		})

		//.SetHeader(`x-api-key`, `MQbTaGOivtnc1y4ldD2JVf3EZWUhKsBYe6gHw5Lj`).SetHeader(`x-app-type`, `TDP`)
		resp, err := resty.R().SetBody(map[string]interface{}{"data": `{
            "binus_tdp": {
              "permohonan": {
                "oss_id": "OSS08D97CJB9FK3",
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
                "status_perusahaan": "1",
                "npwp": "` + p_npwp + `",
                "nama": "` + p_nama + `",
                "alamat": "` + p_alamat + `",
                "idprop": "35",
                "idkab": "3513",
                "idkec": "0",
                "idkel": "0",
                "kode_pos": "0",
                "telepon": "-",
                "fax": "-"
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
                "nama": "` + pj_nama + `",
                "alamat": "` + pj_alamat + `",
                "tempat_lahir": "-",
                "tanggal_lahir": "-",
                "idprop": "35",
                "idkab": "3513",
                "idkec": "-",
                "idkel": "-",
                "telepon": "-",
                "jenis_identitas": "-",
                "no_identitas": "-",
                "negara": "-",
                "email": "-"
              },
              "bank": {
                "bank_1": "-",
                "bank_2": "-",
                "jumlah": "-"
              },
              "modal": {
                "modal_dasar": "-",
                "modal_ditempatkan": "-",
                "modal_disetor": "-",
                "jumlah_saham": "-",
                "nilai_per_saham": "-",
                "modal_sekutu_aktif": "-",
                "modal_sekutu_pasif": "-"
              },
              "legalitas": {
                "no_pendirian": "-",
                "tgl_pendirian": "-",
                "notaris_pendirian": "-",
                "alamat_pendirian": "-",
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
              "ijin_lain": {
                "loop": {
                  "id_ijin": "-",
                  "nomor": "-",
                  "penerbit": "-",
                  "tanggal": "-",
                  "berlaku": "-"
                }
              },
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
                "loop": {
                  "nama": "-",
                  "jabatan": "-",
                  "tempat_lahir": "-",
                  "tanggal_lahir": "-",
                  "alamat": "-",
                  "kode_pos": "-",
                  "telepon": "-",
                  "negara": "-",
                  "tanggal_jabatan": "-",
                  "jumlah_saham": "-",
                  "jumlah_modal": "-",
                  "jabatan_lain": {
                    "loop": {
                      "nama": "-",
                      "jabatan": "-",
                      "alamat": "-",
                      "kode_pos": "-",
                      "telepon": "-",
                      "tanggal_jabatan": "-"
                    }
                  }
                }
              },
              "saham": {
                "loop": {
                  "nama": "-",
                  "alamat": "-",
                  "kode_pos": "-",
                  "telepon": "-",
                  "negara": "-",
                  "npwp": "-",
                  "jumlah_saham": "-",
                  "jumlah_modal": "-"
                }
              },
              "cabang": {
                "loop": {
                  "nama": "-",
                  "no_tdp": "-",
                  "alamat": "-",
                  "idprop": "-",
                  "idkab": "-",
                  "kode_pos": "-",
                  "telepon": "-",
                  "status_perusahaan": "-",
                  "jenis_perusahaan": "-"
                }
              }
            }
          }`}).Post("http://hub.kemendag.go.id/index.php/oss_api/send_req")

		//http://103.29.187.96/android/oss/public/index.php/server/setsipo
		//http: //103.29.187.96/android/oss/public/index.php/server/setapi
		//http://hub.kemendag.go.id/index.php/oss_api/send_req

		// explore response object
		fmt.Printf("\nError: %v", err)
		fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
		fmt.Printf("\nResponse Status: %v", resp.Status())
		fmt.Printf("\nResponse Time: %v", resp.Time())
		fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
		fmt.Printf("\nResponse Body: %v \n", resp) // or resp.String() or string(resp.Body())

		fmt.Println("-----------------------------------")
	}

	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

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

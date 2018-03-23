package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func init() {

}

func main() {
	db, err := sql.Open("mysql", "ossuser:db0ss2017@tcp(10.1.2.87:3306)/oss?charset=utf8")
	errors(err)

	db.SetMaxOpenConns(5)
	defer db.Close()

	var comp []ossHdr

	err = db.QueryRow("SELECT * FROM tblfcosshdr WHERE status = 1").Scan(comp)
	errors(err)

	for _, element := range comp {
		var apli applicant
		var comp dmCompany

		err = db.QueryRow("SELECT * FROM tbldmcompany WHERE company_id = ?", element.Company_id).Scan(comp)
		errors(err)

		err = db.QueryRow("SELECT * FROM tblfccompanyapplicant WHERE applicant_id = ?", element.Applicant_id).Scan(apli)
		errors(err)

		url := "http://10.1.6.87/serveCantik/oss/tdp/"
		var jsonStr = []byte(`{
			"c_paralel" : "1",
			"d_terima_berkas" : "` + time.Now().Format("Y-m-d") + `",
			"d_perubahan" : "0000-00-00",
			"d_perpanjangan" : "0000-00-00",
			"d_daftarulang" : "0000-00-00",
			"no_referensi" : "` + apli.Identity_no + `",
			"n_pemohon" : "` + apli.Name + `",
			"telp_pemohon" : "` + apli.Telp + `",
			"a_pemohon" : "` + apli.Address + `",
			"source" : "` + apli.Identity_type + `",
			"email" : "` + apli.Email + `",
			"n_perusahaan" : "` + comp.Name + `",
			"npwp" : "` + comp.Npwp + `",
			"a_perusahaan" : "` + comp.Address + `",
			"d_tgl_berdiri" : "` + comp.Establish_date + `",
			"i_telp_perusahaan" : "` + comp.Telp + `",
			"no_reg_perusahaan" : "-",
			"rt" : "` + comp.Rt + `",
			"rw" : "` + comp.Rw + `",
			"d_entry" : "` + time.Now().Format("Y-m-d") + `",
			"fax" : "` + comp.Fax + `",
			"no_ip" : "",
			"n_pimpinan" : "",
			"jabatan_pemimpin" : "",
			"id_perusahaan_spipise" : ""
		}`)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("X-Custom-Header", "OSSTDP")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		errors(err)

		defer resp.Body.Close()

		res := new(ResponseTDP)
		json.NewDecoder(resp.Body).Decode(res)

		if res.Status == `00` {
			_, err = db.Exec("UPDATE tbl SET WHERE field = ?")
			errors(err)
		}
	}
}

func errors(err error) {
	if err != nil {
		panic(err)
	}
}

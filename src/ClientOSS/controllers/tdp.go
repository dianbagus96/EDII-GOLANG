package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TdpController struct {
	beego.Controller
}

type ResponseTDP struct {
	Status  string
	Message string
}

type ossHdr struct {
	Status       string
	Oss_id       string
	Applicant_id string
	Company_id   string
}

type ossTdp struct {
}

type dmCompany struct {
	Company_id     string
	Name           string
	Npwp           string
	Address        string
	Prop           string
	Kab            string
	Kec            string
	Kel            string
	Rt             string
	Rw             string
	Telp           string
	Fax            string
	Establish_date string
}

type applicant struct {
	Npwp          string
	Identity_type string
	Identity_no   string
	Name          string
	Address       string
	Telp          string
	Email         string
}

func init() {

	dbConfig := beego.AppConfig.String("dbConfig")
	orm.RegisterDataBase("default", "mysql", dbConfig, 30)
	orm.Debug = true

}

func (c *TdpController) Get() {
	o := orm.NewOrm()
	//var maps []orm.Params
	var message string = ""

	var hdr []ossHdr
	o.Raw("SELECT * FROM tblfcosshdr WHERE status = 1").QueryRows(&hdr)

	for _, element := range hdr {
		var comp dmCompany
		var apli applicant
		if element.Status == "1" {

			//Proses TDP
			o.Raw("SELECT * FROM tbldmcompany WHERE company_id = ?", element.Company_id).QueryRow(&comp)
			o.Raw("SELECT * FROM tblfccompanyapplicant WHERE applicant_id = ?", element.Applicant_id).QueryRow(&apli)

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
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			res := new(ResponseTDP)
			json.NewDecoder(resp.Body).Decode(res)

			//fmt.Println("response Status:", resp.Status)
			//fmt.Println("response Headers:", resp.Header)
			//body, _ := ioutil.ReadAll(resp.Body)
			//fmt.Println("response Body:", string(body))

			message = string(res.Message)

			if res.Status == `00` {
				_, err := o.Raw("UPDATE tblfcosstdp SET  tdp_status = 1 WHERE oss_id = ?", element.Oss_id).Exec()
				if err == nil {
					message = "Berhasil"
				}
			}
		}
	}

	c.Data["json"] = message
	c.ServeJSON()
}

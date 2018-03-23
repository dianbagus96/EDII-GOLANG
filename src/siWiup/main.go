/*package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"




















	_ "github.com/go-sql-driver/mysql"
)

func init() {

}

/*
func main() {
	println(" == Start App ==")

	db, err := sql.Open("mysql", "ossuser:db0ss2017@tcp(10.1.2.87:3306)/oss?charset=utf8")
	errors(err)

	db.SetMaxOpenConns(5)
	defer db.Close()

	url := "http://36.67.90.6:5000/ws/get_wiup_all/data.json?key=TS7W1mKuu1CYrukyZJ5XwAvsFRCEEI"

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Custom-Header", "OSSTDP")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	errors(err)

	defer resp.Body.Close()

	res := []Wiups{}
	//json.NewDecoder(resp.Body).Decode(res)

	body, _ := ioutil.ReadAll(resp.Body)
	json.NewDecoder(resp.Body).Decode(res)

	println(` --- S: RESPONSE --- `)
	println(string(body))
	println(len(res))

	for num, element := range res {
		println(num)
		println(element.A_pemohon)
	}
	println(` --- E: RESPONSE --- `)

	//var comp int
	//err = db.QueryRow("SELECT count(*) FROM tblfcosswiup WHERE id_permohonan = ?", res.Id_permohonan).Scan(&comp)
	//errors(err)

	//if comp == 0 {
	//Insert
	//}

	println(" == End App ==")
}

func errors(err error) {
	if err != nil {
		panic(err)
	}
}
*/

package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	println(` == Start App ==`)
	url := "http://36.67.90.6:5000/ws/get_wiup_all/data.json?key=TS7W1mKuu1CYrukyZJ5XwAvsFRCEEI"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	resp := []people{}

	jsonErr := json.Unmarshal(body, &resp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	db, err := sql.Open("mysql", "ossuser:db0ss2017@tcp(10.1.2.87:3306)/oss?charset=utf8")
	errors(err)

	db.SetMaxOpenConns(5)
	defer db.Close()

	var comp int

	for _, el := range resp {
		errs := db.QueryRow("SELECT count(*) `id` FROM tblfcosswiup WHERE id_permohonan = ?", el.Id_permohonan).Scan(&comp)
		errors(errs)

		panic(string(el.Tgl))
		if comp == 0 {
			var er error
			println(" --- S: Insert ---")
			println(el.Id_permohonan)
			sql := `INSERT INTO tblfcosswiup(
						wiup_id,
						oss_id, 
						a_pemohon, 
						alamat_perusahaan, 
						catatan, 
						i_telp_perusahaan, 
						id_permohonan, 
						n_pemohon, 
						n_perusahaan, 
						nama_pimpinan, 
						nip_ttd, 
						no_pendaftaran, 
						no_reg_perusahaan, 
						no_surat, 
						no_surat_edit, 
						npwp, 
						ref_pendataran, 
						telp_pemohon)
				VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
			_, er = db.Exec(sql, generateCode("WIUP"), generateCode("OSS"), el.A_pemohon, el.Alamat_perusahaan, el.Catatan, el.I_telp_perusahaan, el.Id_permohonan, el.N_pemohon, el.N_perusahaan, el.Nama_pimpinan, el.Nip_ttd, el.No_pendaftaran, el.No_reg_perusahaan, el.No_surat, el.No_surat_edit, el.Npwp, el.Ref_pendataran, el.Telp_pemohon)
			if er != nil {
				log.Println(sql, generateCode("WIUP"), generateCode("OSS"), el.A_pemohon, el.Alamat_perusahaan, el.Catatan, el.I_telp_perusahaan, el.Id_permohonan, el.N_pemohon, el.N_perusahaan, el.Nama_pimpinan, el.Nip_ttd, el.No_pendaftaran, el.No_reg_perusahaan, el.No_surat, el.No_surat_edit, el.Npwp, el.Ref_pendataran, el.Telp_pemohon)
				panic(err.Error)
			}
			println(" --- E: Insert ---")
		}
	}

	println(" == End App ==")
}

func errors(err error) {
	if err != nil {
		panic(err)
	}
}

func shuffle(Title string) string {
	src := []string{
		"A", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	}

	final := make([]string, len(src))
	rand.Seed(time.Now().UTC().UnixNano())
	perm := rand.Perm(len(src))

	for i, v := range perm {
		final[v] = src[i]
	}

	var char string
	for p, _ := range final {
		char += final[p]
	}

	runes := []rune(char)
	var count int = 15 - len(Title)
	safeSubstring := Title + string(runes[0:count])
	return safeSubstring
}

func generateCode(Title string) string {
	shuffled := shuffle(Title)
	return shuffled
}

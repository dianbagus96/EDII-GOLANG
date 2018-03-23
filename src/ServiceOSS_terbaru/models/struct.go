/*
*********************************************************************
	--------------------------------------------------------------
	|															 |
	|  			@2018 - PT EDI INDONESIA - ServiceOSS			 |
	|  			updatedBy :										 |
	|				- Herdian 2018-01-10						 |
	|															 |
	--------------------------------------------------------------
*********************************************************************
*/

package models

type Wrap struct {
	Type string
	Data Permit
}

type Permit struct {
	PPM    IP     `json:"IP"`     //1
	Akta   Akta   `json:"Akta"`   //2
	NPWP   NPWP   `json:"NPWP"`   //3
	TDP    TDP    `json:"TDP"`    //4
	IMTA   IMTA   `json:"IMTA"`   //6
	API    API    `json:"API"`    //7
	NIK    NIK    `json:"NIK"`    //8
	LOKASI LOKASI `json:"LOKASI"` //9
	IMB    IMB    `json:"IMB"`    //10
	MERK   MERK   `json:"MERK"`   //11
	IUI    IUI    `json:"IUI"`    //12
	SNI    SNI    `json:"SNI"`    //13
	//RPTKA  RPTKA  `json:"RPTKA"`  //5
}

type IP struct {
	IDip         string `json:"id_ip"`         //ppm_id
	IDPerusahaan string `json:"id_perusahaan"` //company_id
	/* ----------- S: tbldmcompany ----------- */

	//Hp_perusahaan      string `json:"hp_perusahaan"`      //
	//Prop_pic           string `json:"prop_pic"`           //
	//Kab_pic            string `json:"kab_pic"`            //
	//Kec_pic            string `json:"kec_pic"`            //
	//Kel_pic            string `json:"kel_pic"`            //
	//Hp_pic             string `json:"hp_pic"`             //

	Status_badan       string `json:"status_badan"`       //legal_status
	Status_modal       string `json:"status_modal"`       //invest_status
	Nama_perusahaan    string `json:"nama_perusahaan"`    //name
	Npwp_perusahaan    string `json:"npwp_perusahaan"`    //npwp
	Prop_perusahaan    string `json:"kab_perusahaan"`     //prop
	Kab_perusahaan     string `json:"kec_perusahaan"`     //kab
	Kec_perusahaan     string `json:"kec_perusahaan"`     //kec
	Kel_perusahaan     string `json:"kel_perusahaan"`     //kel
	Alamat_perusahaan  string `json:"alamat_perusahaan"`  //address
	Rt_perusahaan      string `json:"rt_perusahaan"`      //rt
	Rw_perusahaan      string `json:"rw_perusahaan"`      //rw
	Kodepos_perusahaan string `json:"kodepos_perusahaan"` //zip_id
	Telp_perusahaan    string `json:"telp_perusahaan"`    //telp
	Email_perusahaan   string `json:"email_perusahaan"`   //email
	Npwp_pic           string `json:"npwp_pic"`           //pic_identity_no
	Nama_pic           string `json:"nama_pic"`           //pic_name
	Alamat_pic         string `json:"alamat_pic"`         //pic_address
	Jabatan_pic        string `json:"jabatan_pic"`        //pic_position
	Telp_pic           string `json:"telp_pic"`           //pic_telp
	Email_pic          string `json:"email_pic"`          //pic_email

	/*
		----------- E: tbldmcompany -----------

		Query 	 : INSERT INTO tbldmcompany(company_id, legal_status, invest_status, name, npwp, prop, kab, kec, kel, address, rt, rw, zip_id, telp, email, pic_identity_no, pic_name, pic_address, pic_postition, pic_telp, pic_email)
		Variable :

		---------------------------------------
	*/

	/* ----------- S: tblfccompanydeed ----------- */

	Nama_notaris         string `json:"nama_notaris"`         //notaris_name
	No_akta              string `json:"no_akta"`              //deed_no
	Tgl_akta             string `json:"tgl_akta"`             //deed_date
	No_pengesahaan_akta  string `json:"no_pengesahaan_akta"`  //
	Tgl_pengesahaan_akta string `json:"tgl_pengesahaan_akta"` //

	/*
		----------- E: tblfccompanydeed -----------

		Query 	 : INSERT INTO tblfccompanydeed(notaris_name, deed_no, deed_date)
		Variable :

		-------------------------------------------
	*/

	/* ----------- S: tblfccompanybusiness ----------- */

	Jadwal_proyek             string `json:"jadwal_proyek"`             //project_sch
	Bidang_usaha              string `json:"bidang_usaha"`              //line_business
	Kbli                      string `json:"kbli"`                      //kbli_id
	Prop_bidang_usaha         string `json:"prop_bidang_usaha"`         //prop
	Kab_bidang_usaha          string `json:"kab_bidang_usaha"`          //kab
	Alamat_bidang_usaha       string `json:"alamat_bidang_usaha"`       //address
	Status_tanah_bidang_usaha string `json:"status_tanah_bidang_usaha"` //land_type
	Jumlah_tki                string `json:"jumlah_tki"`                //imp_total
	Jumlah_tka                string `json:"jumlah_tka"`                //fmp_total
	Nilai_tanah               string `json:"nilai_tanah"`               //purchase_land
	Nilai_mesing_peralatan    string `json:"nilai_mesing_peralatan"`    //machine_equip
	Nilai_mesin_peralatan_us  string `json:"nilai_mesin_peralatan_us"`  //machinge_equip_us
	Modal_kerja               string `json:"modal_kerja"`               //working_capital
	Total_investasi           string `json:"total_investasi"`           //total_invest
	Total_aset                string `json:"total_aset"`                //total_asset
	Total_omset               string `json:"total_omset"`               //total_omset

	/*
		----------- E: tblfccompanybusiness -----------

		Query 	 : INSERT INTO tblfccompanybusiness(project_sch, line_business, kbli_id, prop, kab, address, land_type, imp_total, fmp_total, purchase_land, machine_equip, machinge_equip_us, working_capital, total_invest, total_asset, total_omset)
		Variable :

		-------------------------------------------
	*/

	Pemegang_saham Shareholder `json:"pemegang_saham"`

	/* ----------- S: tblfccompanyshare ----------- */

	Jumlah_saham_lembar      string //number_of_share
	Nilai_saham_lembar       string //value_of_share
	Total_saham_asing        string //total_foreign_share
	Total_persen_saham_asing string //total_percent_foreign
	Total_saham_lokal        string //total_domestic_share
	Total_saham              string //total_value_share
	Modal_dasar              string //authorized_capital
	Modal_ditempatkan        string //issued_capital
	Ekuitas                  string //equity
	Modal_disetor            string //paid_up_capital
	Laba                     string //return_earning
	Pinjaman_asing           string //foreign_loan
	Pinjaman_lokal           string //domestic_loan
	Total_modal              string //total_capital
	Total_dana               string //total_funds

	/*
		----------- E: tblfccompanyshare -----------

		Query 	 : INSERT INTO tblfccompanyshare(number_of_share, value_of_share, total_foreign_share, total_percent_foreign, total_domestic_share, total_value_share, authorized_capital, issued_capital, equity, paid_up_capital, return_earning, foreign_loan, domestic_loan, total_capital, total_funds)
		Variable :

		-------------------------------------------
	*/

	/* ----------- S: tblfccompanyapplicant ----------- */

	Nama_pemohon     string //name
	Jabatan_pemohon  string //position
	Jenis_id_pemohon string //identiity_type
	No_id_pemohon    string //identity_no
	Prop_pemohon     string //prop
	Kab_pemohon      string //kab
	Kec_pemohon      string //kec
	Kel_pemohon      string //kel
	Alamat_pemohon   string //address
	Kodepos_pemohon  string //zip_id
	Telp_pemohon     string //telp
	Hp_pemohon       string //hp
	Email_pemohon    string //email

	/*
		----------- E: tblfccompanyapplicant -----------

		Query 	 : INSERT INTO tblfccompanyshare(name, position, identity_type, identity_no, prop, kab, kec, kel, address, zip_id, telp, hp, email)
		Variable :

		-------------------------------------------
	*/

	/* ----------- S: tblfcossppm ----------- */

	No_ppm          string `json:no_ppm`          //ppm_no
	Tgl_ppm         string `json:tgl_ppm`         //ppm_date
	Nama_ttd_ppm    string `json:nama_ttd_ppm`    //ppm_sign_name
	Tgl_ttd_ppm     string `json:tgl_ttd_ppm`     //ppm_sign_date
	Jabatan_ttd_ppm string `json:jabatan_ttd_ppm` //ppm_sign_position
	Penerbit        string `penerbit:no_ppm`      //ppm_sign_ga

	/*
		----------- E: tblfcossppm -----------

		Query 	 : UPDATE tblfcossppm SET ppm_no = ?, ppm_date = ?, ppm_sign_name = ?, ppm_sign_date = ?, ppm_sign_position = ?, ppm_sign_ga = ? WHERE ppm_id = ?
		Variable :

		-------------------------------------------
	*/
}

type Shareholder []struct {

	/* ----------- S: tblfccompanyshareholder ----------- */

	Nama_pemegang_saham     string `json:"nama_pemegang_saham"`     //shareholder_name
	Jenis_id_pemegang_saham string `json:"jenis_id_pemegang_saham"` //identity_type
	No_id_pemegang_Saham    string `json:"no_id_pemegang_Saham"`    //identity_no
	Masa_berlaku_id         string `json:"masa_berlaku_id"`         //identity_valid
	Npwp_pemegang_saham     string `json:"npwp_pemegang_saham"`     //npwp
	Alamat_pemegang_saham   string `json:"alamat_pemegang_saham"`   //address
	Negara_pemegang_saham   string `json:"negara_pemegang_saham"`   //country
	Telp_pemegang_saham     string `json:"telp_pemegang_saham"`     //telp
	Fax_pemegang_saham      string `json:"fax_pemegang_saham"`      //fax
	Email_pemegang_saham    string `json:"email_pemegang_saham"`    //email
	Nilai_saham             string `json:"nilai_saham"`             //share_value

	/*
		----------- E: tblfccompanyshareholder -----------

		Query 	 : INSERT INTO tblfccompanyshareholder(shareholder_name, identity_type, identity_no, identity_valid, npwp, address, country, telp, fax, email, share_value)
		Variable :

		--------------------------------------------------
	*/
}

type Akta struct {
	//Request
	Nama_notaris string `json:"nama_notaris"`

	//Terbit
	No_akta          string `json:"no_akta"`
	Tgl_akta         string `json:"tgl_akta"`
	Nama_ttd_akta    string `json:"nama_ttd_akta"`
	Nip_ttd_akta     string `json:"nip_ttd_akta"`
	Tgl_ttd_akta     string `json:"tgl_ttd_akta"`
	Jabatan_ttd_akta string `json:"jabatan_ttd_akta"`
	Penerbit         string `json:"penerbit"`
	Catatan          string `json:"catatan"`
}

type NPWP struct {
	//Request
	No_lhp            string `json:"no_lhp"`
	Tahun_pajak       string `json:"tahun_pajak"`
	Tahun_buku        string `json:"tahun_buku"`
	Nama_bendahara    string `json:"nama_bendahara"`
	Divisi_bendahara  string `json:"divisi_bendahara"`
	Alamat_bendahara  string `json:"alamat_bendahara"`
	Blok_bendahara    string `json:"blok_bendahara"`
	Rt_bendahara      string `json:"rt_bendahara"`
	Rw_bendahara      string `json:"rw_bendahara"`
	Prop_bendahara    string `json:"prop_bendahara"`
	Kab_bendahara     string `json:"kab_bendahara"`
	Kec_bendahara     string `json:"kec_bendahara"`
	Kel_bendahara     string `json:"kel_bendahara"`
	Telp_bendahara    string `json:"telp_bendahara"`
	Fax_bendahara     string `json:"fax_bendahara"`
	Kodepos_bendahara string `json:"kodepos_bendahara"`
	No_referensi      string `json:"no_referensi"`
	Nama_pegawai      string `json:"nama_pegawai"`
	Nik_pegawai       string `json:"nik_pegawai"`
	Npwp_pegawai      string `json:"npwp_pegawai"`
	Alamat_pegawai    string `json:"alamat_pegawai"`
	Blok_pegawai      string `json:"blok_pegawai"`
	Rt_pegawai        string `json:"rt_pegawai"`
	Rw_pegawai        string `json:"rw_pegawai"`
	Prop_pegawai      string `json:"prop_pegawai"`
	Kab_pegawai       string `json:"kab_pegawai"`
	Kec_pegawai       string `json:"kec_pegawai"`
	Kel_pegawai       string `json:"kel_pegawai"`
	Kodepos_pegawai   string `json:"kodepos_pegawai"`
	Telp_pegawai      string `json:"telp_pegawai"`
	Fax_pegawai       string `json:"fax_pegawai"`
	Hp_pegawai        string `json:"hp_pegawai"`
	Email_pegawai     string `json:"email_pegawai"`

	//Terbit
	No_npwp          string `json:"no_npwp"`
	Tgl_npwp         string `json:"tgl_npwp"`
	Nama_ttd_npwp    string `json:"nama_ttd_npwp"`
	Nip_ttd_npwp     string `json:"nip_ttd_npwp"`
	Tgl_ttd_npwp     string `json:"tgl_ttd_npwp"`
	Jabatan_ttd_npwp string `json:"jabatan_ttd_npwp"`
	Penerbit         string `json:"penerbit`
	Catatan          string `json:"catatan"`
}

type TDP struct {
	//Request
	Kbli           string `json:"kbli"`
	Bentuk_usaha   string `json:"bentuk_usaha"`
	Tgl_kegiatan   string `json:"tgl_kegiatan"`
	Kegiatan_pokok string `json:"kegiatan_pokok"`
	Kegiatan_lain  string `json:"kegiatan_lain"`
	Hak_cipta      string `json:"hak_cipta"`
	Hak_paten      string `json:"hak_paten"`
	Komoditi_utama string `json:"komoditi_utama"`
	Komoditi_lain  string `json:"komoditi_lain"`
	No_agenda      string `json:"no_agenda"`

	//Terbit
	No_tdp          string `json:"no_tdp"`
	Tgl_tdp         string `json:"tgl_tdp"`
	Nama_ttd_tdp    string `json:"nama_ttd_tdp"`
	Nip_ttd_tdp     string `json:"nip_ttd_tdp"`
	Tgl_ttd_tdp     string `json:"tgl_ttd_tdp"`
	Jabatan_ttd_tdp string `json:"jabatan_ttd_tdp"`
	Penerbit        string `json:"penerbit"`
	Catatan         string `json:"catatan"`
}

/*
type RPTKA struct {
	No_siup            string `json:"no_siup"`
	Tgl_siup           string `json:"tgl_siup"`
	Tgl_berlaku_siup   string `json:"tgl_berlaku_siup"`
	Jabatan            string `json:"jabatan"`
	Total_tkwnap       string `json:"total_tkwnap"`
	Total_tki          string `json:"total_tki"`
	Tgl_berlaku_tkwnap string `json:"tgl_berlaku_tkwnap"`
	Pendidikan_tki     string `json:"pendidikan_tki"`
	Pengalaman_tki     string `json:"pengalaman_tki"`
	Catatan            string `json:"catatan"`

	Posisi             string `json:"posisi"`
	Uraian_pekerjaan   string `json:"uraian_pekerjaan"`
	Pendidikan         string `json:"pendidikan"`
	Pengalaman_kerja   string `json:"pengalaman_kerja"`
	Catatan            string `json:"catatan"`
	Posisi             string `json:"posisi"`
	Total_tkwnap       string `json:"total_tkwnap"`
	Tgl_berlaku_tkwnap string `json:"tgl_berlaku_tkwnap"`
	Mulai_kerja        string `json:"mulai_kerja"`
	Catatan            string `json:"catatan"`
	No_rptka           string `json:"no_rptka"`
}
*/

type IMTA struct {
	//Request
	Nama_lengkap     string `json:"nama_lengkap"`
	Jenis_kelamin    string `json:"jenis_kelamin"`
	Tempat_lahir     string `json:"tempat_lahir"`
	Tgl_lahir        string `json:"tgl_lahir"`
	Alamat           string `json:"alamat"`
	Negara           string `json:"negara"`
	Alamat_tinggal   string `json:"alamat_tinggal"`
	Blok             string `json:"blok"`
	Rt               string `json:"rt"`
	Rw               string `json:"rw"`
	Prop             string `json:"prop"`
	Kab              string `json:"kab"`
	Kec              string `json:"kec"`
	Kel              string `json:"kel"`
	Tipe_visa        string `json:"tipe_visa"`
	No_visa          string `json:"no_visa"`
	Tgl_visa         string `json:"tgl_visa"`
	Tgl_berlaku_visa string `json:"tgl_berlaku_visa"`
	No_kim           string `json:"no_kim"`
	Tgl_kim          string `json:"tgl_kim"`
	Tgl_berlaku_kim  string `json:"tgl_berlaku_kim"`
	No_stmd          string `json:"no_stmd"`
	Tgl_stmd         string `json:"tgl_stmd"`
	Tgl_berlaku_stmd string `json:"tgl_berlaku_stmd"`
	No_skk           string `json:"no_skk"`
	Tgl_skk          string `json:"tgl_skk"`
	Tgl_berlaku_skk  string `json:"tgl_berlaku_skk"`
	Jabatan          string `json:"jabatan"`
	Level_jabatan    string `json:"level_jabatan"`
	Pendidikan       string `json:"pendidikan"`
	Pengalaman_kerCa string `json:"pengalaman_kerja"`
	Tempat_kerja     string `json:"tempat_kerja"`
	Prop_kerja       string `json:"prop_kerja"`
	Kab_kerja        string `json:"kab_kerja"`
	Kec_kerja        string `json:"kec_kerja"`
	Kel_kerja        string `json:"kel_kerja"`
	Mulai_kontrak    string `json:"mulai_kontrak"`
	Selesai_kontrak  string `json:"selesai_kontrak"`
	Fas_rumah        string `json:"fas_rumah"`
	Fas_kendaraan    string `json:"fas_kendaraan"`
	Gaji             string `json:"gaji"`

	//Terbit
	No_imta          string `json:"no_imta"`
	Tgl_imta         string `json:"tgl_imta"`
	Nama_ttd_imta    string `json:"nama_ttd_imta"`
	Nip_ttd_imta     string `json:"nip_ttd_imta"`
	Tgl_ttd_imta     string `json:"tgl_ttd_imta"`
	Jabatan_ttd_imta string `json:"jabatan_ttd_imta"`
	Penerbit         string `json:"penerbit"`
	Catatan          string `json:"catatan"`
}

type API struct {
	//Request
	No_permohonan      string `json:"no_permohonan"`
	Tgl_permohonan     string `json:"tgl_permohonan"`
	No_referensi_bank  string `json:"no_referensi_bank"`
	Tgl_referensi_bank string `json:"tgl_referensi_bank"`

	//Terbit
	No_api          string `json:"no_api"`
	Tgl_api         string `json:"tgl_api"`
	Nama_ttd_api    string `json:"nama_ttd_api"`
	Nip_ttd_api     string `json:"nip_ttd_api"`
	Tgl_ttd_api     string `json:"tgl_ttd_api"`
	Jabatan_ttd_api string `json:"jabatan_ttd_api"`
	Penerbit        string `json:"penerbit"`
	Catatan         string `json:"catatan"`
}

type NIK struct {
	//Request
	Pendapatan                 string `json:"pendapatan"`
	Pendapatan_kotor           string `json:"pendapatan_kotor"`
	Aset                       string `json:"aset"`
	Aset_tetap                 string `json:"aset_tetap"`
	Aset_lain                  string `json:"aset_lain"`
	Total_aset                 string `json:"total_aset"`
	Hutang_jangka_panjang      string `json:"hutang_jangka_panjang"`
	Hutang_jangka_pendek       string `json:"hutang_jangka_pendek"`
	Total_hutang               string `json:"total_hutang"`
	Piutang                    string `json:"piutang"`
	Total_modal                string `json:"total_modal"`
	Akuntan_publik             string `json:"akuntan_publik"`
	Laporan_publik             string `json:"laporan_publik"`
	Pendapat_publik            string `json:"pendapat_publik"`
	No_lha_bc                  string `json:"no_lha_bc"`
	Laporan_bc                 string `json:"laporan_bc"`
	Pendapat_bc                string `json:"pendapat_bc"`
	Sistem_akuntan             string `json:"sistem_akuntan"`
	It_inventori               string `json:"it_inventori"`
	Ahli_pabean                string `json:"ahli_pabean"`
	Noseri_sertifikat          string `json:"noseri_sertifikat"`
	No_sertifikat              string `json:"no_sertifikat"`
	Nama_kapal                 string `json:"nama_kapal"`
	Noreg_kapal                string `json:"noreg_kapal"`
	Kapasitas_kapal            string `json:"kapasitas_kapal"`
	Tipe_rute_kapal            string `json:"tipe_rute_kapal"`
	Rute_asal_kapal            string `json:"rute_asal_kapal"`
	Rute_tujuan_kapal          string `json:"rute_tujuan_kapal"`
	Nama_pesawat               string `json:"nama_pesawat"`
	Noreg_pesawat              string `json:"noreg_pesawat"`
	Kapasitas_pesawat          string `json:"kapasitas_pesawat"`
	Tipe_rute_pesawat          string `json:"tipe_rute_pesawat"`
	Rute_asal_pesawat          string `json:"rute_asal_pesawat"`
	Rute_tujuan_pesawat        string `json:"rute_tujuan_pesawat"`
	No_skep_tps                string `json:"no_skep_tps"`
	Tipe_tps                   string `json:"tipe_tps"`
	Izin_pjt                   string `json:"izin_pjt"`
	Nama_ahli_pjt              string `json:"nama_ahli_pjt"`
	Jabatan_ahli_pjt           string `json:"jabatan_ahli_pjt"`
	Noseri_sertifikat_pjt      string `json:"noseri_sertifikat_pjt"`
	No_sertifikat_pjt          string `json:"no_sertifikat_pjt"`
	No_sertifkat_iso           string `json:"no_sertifkat_iso"`
	Tgl_sertifikat_iso         string `json:"tgl_sertifikat_iso"`
	Tgl_berlaku_sertifikat_iso string `json:"tgl_berlaku_sertifikat_iso"`
	Penerbit_iso               string `json:"penerbit_iso"`
	Nama_ahli_lain             string `json:"nama_ahli_lain"`
	Jabatan_ahli_lain          string `json:"jabatan_ahli_lain"`
	Noseri_sertifikat_lain     string `json:"noseri_sertifikat_lain"`
	No_sertifikat_lain         string `json:"no_sertifikat_lain"`
	Anggota_asosiasi           string `json:"anggota_asosiasi"`
	Hasil_survey               string `json:"hasil_survey"`
	Catatan_survey             string `json:"catatan_survey"`

	//Terbit
	No_nik          string `json:"no_nik"`
	Tgl_nik         string `json:"tgl_nik"`
	Nama_ttd_nik    string `json:"nama_ttd_nik"`
	Nip_ttd_nik     string `json:"nip_ttd_nik"`
	Tgl_ttd_nik     string `json:"tgl_ttd_nik"`
	Jabatan_ttd_nik string `json:"jabatan_ttd_nik"`
	Penerbit        string `json:"penerbit"`
	Catatan         string `json:"catatan"`
}

type LOKASI struct {
	//Request
	Luas_wilayah          string `json:"luas_wilayah"`
	Prop_wilayah          string `json:"prop_wilayah"`
	Kab_wilayah           string `json:"kab_wilayah"`
	Kec_wilayah           string `json:"kec_wilayah"`
	Kel_wilayah           string `json:"kel_wilayah"`
	Status_tanah          string `json:"status_tanah"`
	Fungsi_tanah_skrng    string `json:"fungsi_tanah_skrng"`
	Fungsi_tanah_nanti    string `json:"fungsi_tanah_nanti"`
	Batas_wilayah_utara   string `json:"batas_wilayah_utara"`
	Batas_wilayah_timur   string `json:"batas_wilayah_timur"`
	Batas_wilayah_selatan string `json:"batas_wilayah_selatan"`
	Batas_wilayah_barat   string `json:"batas_wilayah_barat"`

	//Terbit
	No_lokasi          string `json:"no_lokasi"`
	Tgl_lokasi         string `json:"tgl_lokasi"`
	Nama_ttd_lokasi    string `json:"nama_ttd_lokasi"`
	Nip_ttd_lokasi     string `json:"nip_ttd_lokasi"`
	Tgl_ttd_lokasi     string `json:"tgl_ttd_lokasi"`
	Jabatan_ttd_lokasi string `json:"jabatan_ttd_lokasi"`
	Penerbit           string `json:"penerbit"`
	Catatan            string `json:"catatan"`
}

type IMB struct {
	//Request
	Konstruksi    string `json:"konstruksi"`
	Tipe_bangunan string `json:"tipe_bangunan"`

	//Terbit
	No_imb          string `json:"no_imb"`
	Tgl_imb         string `json:"tgl_imb"`
	Nama_ttd_imb    string `json:"nama_ttd_imb"`
	Nip_ttd_imb     string `json:"nip_ttd_imb"`
	Tgl_ttd_imb     string `json:"tgl_ttd_imb"`
	Jabatan_ttd_imb string `json:"jabatan_ttd_imb"`
	Penerbit        string `json:"penerbit"`
	Catatan         string `json:"catatan"`
}

type MERK struct {
	//Request
	Lokasi_pabrik    string `json:"lokasi_pabrik"`
	No_permohonan    string `json:"no_permohonan"`
	Tgl_permohonan   string `json:"tgl_permohonan"`
	No_penerimaan    string `json:"no_penerimaan"`
	Tgl_penerimaan   string `json:"tgl_penerimaan"`
	Nama_konsultan   string `json:"nama_konsultan"`
	Kantor_konsultan string `json:"kantor_konsultan"`
	Alamat_konsultan string `json:"alamat_konsultan"`
	Telp_konsultan   string `json:"telp_konsultan"`
	Fax_konsultan    string `json:"fax_konsultan"`
	Email_konsultan  string `json:"email_konsultan"`
	No_prioritas     string `json:"no_prioritas"`
	Tgl_prioritas    string `json:"tgl_prioritas"`
	Kelas_merk       string `json:"kelas_merk"`
	Tipe_merk        string `json:"tipe_merk"`
	Terjemahan_merk  string `json:"terjemahan_merk"`
	Warna_merk       string `json:"warna_merk"`
	Label_merk       string `json:"label_merk"`
	Nama_merk        string `json:"nama_merk"`
	Deskripsi_merk   string `json:"deskripsi_merk"`

	//Terbit
	No_merk          string `json:"no_merk"`
	Tgl_merk         string `json:"tgl_merk"`
	Nama_ttd_merk    string `json:"nama_ttd_merk"`
	Nip_ttd_merk     string `json:"nip_ttd_merk"`
	Tgl_ttd_merk     string `json:"tgl_ttd_merk"`
	Jabatan_ttd_merk string `json:"jabatan_ttd_merk"`
	Penerbit         string `json:"penerbit"`
	Catatan          string `json:"catatan"`
}

type IUI struct {
	//Request
	Luas_gudang         string `json:"luas_gudang"`
	Bangunan_persen     string `json:"bangunan_persen"`
	Peralatan_impor     string `json:"peralatan_impor"`
	Peralatan_lokal     string `json:"peralatan_lokal"`
	Mesin               string `json:"mesin"`
	Pabrik_mulai        string `json:"pabrik_mulai"`
	Produksi_mulai      string `json:"produksi_mulai"`
	Pasar_lokal_persen  string `json:"pasar_lokal_persen"`
	Pasar_ekspor_persen string `json:"pasar_ekspor_persen"`
	Merk_persen         string `json:"merk_persen"`

	//Terbit
	No_iui          string `json:"no_iui"`
	Tgl_iui         string `json:"tgl_iui"`
	Nama_ttd_iui    string `json:"nama_ttd_iui"`
	Nip_ttd_iui     string `json:"nip_ttd_iui"`
	Tgl_ttd_iui     string `json:"tgl_ttd_iui"`
	Jabatan_ttd_iui string `json:"jabatan_ttd_iui"`
	Penerbit        string `json:"penerbit"`
	Catatan         string `json:"catatan"`
}

type SNI struct {
	//Request
	Judul_sni string `json:"judul_sni"`
	Qc_sni    string `json:"qc_sni"`

	//Terbit
	No_sni          string `json:"no_sni"`
	Tgl_sni         string `json:"tgl_sni"`
	Nama_ttd_sni    string `json:"nama_ttd_sni"`
	Nip_ttd_sni     string `json:"nip_ttd_sni"`
	Tgl_ttd_sni     string `json:"tgl_ttd_sni"`
	Jabatan_ttd_sni string `json:"jabatan_ttd_sni"`
	Penerbit        string `json:"penerbit"`
	Catatan         string `json:"catatan"`
}

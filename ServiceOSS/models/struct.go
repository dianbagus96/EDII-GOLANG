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

type Object struct {
	//TBLFCOSSPPM
	OSSID              string
	ID                 string
	CompanyID          string
	CompanyNPWP        string
	CompanyName        string
	InstituteName      string
	RegStatus          string
	Status             string
	PermitNo           string
	PermitDate         string
	PermitNote         string
	PermitSignName     string
	PermitSignNIP      string
	PermitSignGA       string
	PermitSignDate     string
	PermitSignPosition string
	PermitChecklist    string
	DateCreated        string
	DateUpdated        string
	BusinessID         string
	InstituteID        string
	FlagChecklist      string
	PermitID           string

	//TBLFCOSSHDR
	PermitGroup      string
	OssDate          string
	ApplicantID      string
	FlagPrepare      string
	FlagConstruction string
	FlagProduction   string
	CreatedBy        string
	CreatedDate      string
}

type Object2 struct {
	StatusBadan    string
	StatusModal    string
	CompanyName    string
	CompanyNPWP    string
	CompanyProv    string
	CompanyKab     string
	CompanyKec     string
	CompanyKel     string
	ComapnyAddress string
	CompanyRT      string
	CompanyRW      string
	CompanyZipCode string
	CompanyPhone   string
	CompanyHP      string
	CompanyMail    string

	PicNPWP     string
	PicName     string
	PicProv     string
	PicKab      string
	PicKec      string
	PicKel      string
	PicAddress  string
	PicPosition string
	PicTelp     string
	PicHP       string
	PicMail     string

	NotarisName          string
	DeedNo               string
	DeedDate             string
	DeedRatificationNo   string
	DeedRatificationDate string

	ProjectDate          string
	BusinessField        string
	KBLINo               string
	BusinessFieldProp    string
	BusinessFieldKab     string
	BusinessFieldAddress string
	StatusLandField      string
	LargeField           string
	TKITotal             string
	TKATotal             string
	ValueLand            string
	ValueMachineTool     string
	ValueMachineToolUS   string
	WorkingAsset         string
	TotalInvestment      string
	TotalAsset           string
	TotalOmset           string

	ShareholdersName         string
	ShareholdersType         string
	ShareholdersNumber       string
	ShareholdersExpiredID    string
	ShareholdersNPWP         string
	ShareholdersAddress      string
	ShareholdersCountry      string
	ShareholderPhone         string
	ShareholderFax           string
	HoldingValue             string
	HoldingTotalSheet        string
	HoldingTotalForeign      string
	HoldingForeignPercentage string
	HoldingTotal             string
	HoldingPercentage        string
	//ShareholdersName         string

	AuthorizeHolding string
	AuthorizePlaced  string
	AuthorizeStored  string
	Equity           string
	Profit           string
	ForeignLoan      string
	ForeignLocal     string
	TotalModal       string
	TotalFund        string

	ApplicantName     string
	ApplicantPosition string
	ApplicantType     string
	ApplicantNo       string
	ApplicantProp     string
	ApplicantKab      string
	ApplicantKec      string
	ApplicantKel      string
	ApplicantAddress  string
	ApplicantZipCode  string
	ApplicantPhone    string
	ApplicantHP       string
	ApplicantMail     string
}

type PPM struct {
	PPM Object3 `json:"IP"`
}

type Object3 struct {
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
	Prop_perusahaan    string `json:"prop_perusahaan"`    //prop
	Kab_perusahaan     string `json:"kab_perusahaan"`     //kab
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

	no_ppm          string `json:no_ppm`          //ppm_no
	tgl_ppm         string `json:tgl_ppm`         //ppm_date
	nama_ttd_ppm    string `json:nama_ttd_ppm`    //ppm_sign_name
	nip_ttd_ppm     string `json:nip_ttd_ppm`     //ppm_sign_name
	tgl_ttd_ppm     string `json:tgl_ttd_ppm`     //ppm_sign_date
	jabatan_ttd_ppm string `json:jabatan_ttd_ppm` //ppm_sign_position
	penerbit        string `json:penerbit`        //ppm_sign_ga
	catatan         string `json:catatan`         //ppm_sign_position

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

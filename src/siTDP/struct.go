package main

import (
	"encoding/xml"
)

type Result struct {
	XMLName xml.Name `xml:"Binus_tdp"`

	//TAG PERMOHONAN
	JenisPermohonan string `xml:"permohonan>jenis_permohonan"`
	Pembaharuan     string `xml:"permohonan>pembaharuan"`
	NoIjin          string `xml:"permohonan>no_ijin"`
	TglIjin         string `xml:"permohonan>tgl_ijin"`
	TglAkhir        string `xml:"permohonan>tgl_akhir"`
	Kbli            string `xml:"permohonan>kbli"`
	Pokok           string `xml:"permohonan>pokok"`
	NamaTtd         string `xml:"permohonan>nama_ttd"`
	JabatTtd        string `xml:"permohonan>jabat_ttd"`
	NipTtd          string `xml:"permohonan>nip_ttd"`

	//TAG PERUSAHAAN
	BentukUsaha      string `xml:"perusahaan>bentuk_usaha"`
	StatusPerusahaan string `xml:"perusahaan>status_perusahaan"`
	Npwp             string `xml:"perusahaan>npwp"`
	Nama             string `xml:"perusahaan>nama"`
	Alamat           string `xml:"perusahaan>alamat"`
	Idprop           string `xml:"perusahaan>idprop"`
	Idkab            string `xml:"perusahaan>idkab"`
	Idkec            string `xml:"perusahaan>idkec"`
	Idkel            string `xml:"perusahaan>idkel"`
	KodePos          string `xml:"perusahaan>kode_pos"`
	Telepon          string `xml:"perusahaan>telepon"`
	Fax              string `xml:"perusahaan>fax"`
	Idgroup          string `xml:"perusahaan>idgroup"`

	//TAG KEGIATAN_PERUSAHAAN
	LokasiProduksi          string `xml:"kegiatan_perusahaan>lokasi_produksi"`
	PropProduksi            string `xml:"kegiatan_perusahaan>prop_produksi"`
	KabProduksi             string `xml:"kegiatan_perusahaan>kab_produksi"`
	IDTanamModal            string `xml:"kegiatan_perusahaan>id_tanam_modal"`
	TglPendirian            string `xml:"kegiatan_perusahaan>tgl_pendirian"`
	TglKegiatan             string `xml:"kegiatan_perusahaan>tgl_kegiatan"`
	LamaBerdiri             string `xml:"kegiatan_perusahaan>lama_berdiri"`
	IDKerjasama             string `xml:"kegiatan_perusahaan>id_kerjasama"`
	NamaMerekDagang         string `xml:"kegiatan_perusahaan>nama_merek_dagang"`
	NoMerekDagang           string `xml:"kegiatan_perusahaan>no_merek_dagang"`
	PemegangPaten           string `xml:"kegiatan_perusahaan>pemegang_paten"`
	NoPaten                 string `xml:"kegiatan_perusahaan>no_paten"`
	PemegangCipta           string `xml:"kegiatan_perusahaan>pemegang_cipta"`
	NoCipta                 string `xml:"kegiatan_perusahaan>no_cipta"`
	JumlahWni               string `xml:"kegiatan_perusahaan>jumlah_wni"`
	JumlahWna               string `xml:"kegiatan_perusahaan>jumlah_wna"`
	JenisPerusahaan         string `xml:"kegiatan_perusahaan>jenis_perusahaan"`
	Lain1                   string `xml:"kegiatan_perusahaan>lain_1"`
	Lain2                   string `xml:"kegiatan_perusahaan>lain_2"`
	KomoditiPokok           string `xml:"kegiatan_perusahaan>komoditi_pokok"`
	KomoditiLain1           string `xml:"kegiatan_perusahaan>komoditi_lain_1"`
	KomoditiLain2           string `xml:"kegiatan_perusahaan>komoditi_lain_2"`
	Omset                   string `xml:"kegiatan_perusahaan>omset"`
	TotalAset               string `xml:"kegiatan_perusahaan>total_aset"`
	IDMataRantai            string `xml:"kegiatan_perusahaan>id_mata_rantai"`
	FlagMesin               string `xml:"kegiatan_perusahaan>flag_mesin"`
	KapasitasTerpasang      string `xml:"kegiatan_perusahaan>kapasitas_terpasang"`
	KapasitasSatuan         string `xml:"kegiatan_perusahaan>kapasitas_satuan"`
	KapasitasProduksi       string `xml:"kegiatan_perusahaan>kapasitas_produksi"`
	KapasitasProduksiSatuan string `xml:"kegiatan_perusahaan>kapasitas_produksi_satuan"`
	KandunganLokal          string `xml:"kegiatan_perusahaan>kandungan_lokal"`
	KandunganImpor          string `xml:"kegiatan_perusahaan>kandungan_impor"`
	IDPengecer              string `xml:"kegiatan_perusahaan>id_pengecer"`

	//TAG KOPERASI
	JenisKoperasi         string `xml:"koperasi>jenis_koperasi"`
	JumlahAnggotaKoperasi string `xml:"koperasi>jumlah_anggota_koperasi"`
	SimpananPokok         string `xml:"koperasi>simpanan_pokok"`
	SimpananWajib         string `xml:"koperasi>simpanan_wajib"`
	DanaCadangan          string `xml:"koperasi>dana_cadangan"`
	Hibah                 string `xml:"koperasi>hibah"`
	PinjamanAnggota       string `xml:"koperasi>pinjaman_anggota"`
	PinjamanKoperasi      string `xml:"koperasi>pinjaman_koperasi"`
	PinjamanBank          string `xml:"koperasi>pinjaman_bank"`
	PinjamanLain          string `xml:"koperasi>pinjaman_lain"`

	//TAG PERUSAHAAN_INDUK
	PRNama   string `xml:"perusahaan_induk>nama"`
	PRNoTdp  string `xml:"perusahaan_induk>no_tdp"`
	PRAlamat string `xml:"perusahaan_induk>alamat"`
	PRIDProp string `xml:"perusahaan_induk>idprop"`
	PRIDKab  string `xml:"perusahaan_induk>idkab"`
	PRIDKec  string `xml:"perusahaan_induk>idkec"`
	PRIDKel  string `xml:"perusahaan_induk>idkel"`

	//TAG PENANGGUNG_JAWAB
	PENama           string `xml:"nama"`
	PEAlamat         string `xml:"alamat"`
	PETempatLahir    string `xml:"tempat_lahir"`
	PETanggalLahir   string `xml:"tanggal_lahir"`
	PEIDProp         string `xml:"idprop"`
	PEIDKab          string `xml:"idkab"`
	PEIDKec          string `xml:"idkec"`
	PEIDKel          string `xml:"idkel"`
	PETelepon        string `xml:"telepon"`
	PEJenisIdentitas string `xml:"jenis_identitas"`
	PENoIdentitas    string `xml:"no_identitas"`
	PENegara         string `xml:"negara"`
	PEEmail          string `xml:"email"`

	/*
		//TAG BANK
		bank_1
		bank_2
		jumlah

		//TAG MODAL
		modal_dasar
		modal_ditempatkan
		modal_disetor
		jumlah_saham
		nilai_per_saham
		modal_sekutu_aktif
		modal_sekutu_pasif

		//TAG LEGALITAS
	*/
}

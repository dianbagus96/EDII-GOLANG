package main

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

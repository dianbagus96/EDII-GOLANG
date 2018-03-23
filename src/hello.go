package serverbinus

import (
	"github.com/fiorix/wsdl2go/soap"
)

// Namespace was auto-generated from WSDL.
var Namespace = "urn:ServerBinus"

// NewServerBinusPortType creates an initializes a ServerBinusPortType.
func NewServerBinusPortType(cli *soap.Client) ServerBinusPortType {
	return &serverBinusPortType{cli}
}

// ServerBinusPortType was auto-generated from WSDL
// and defines interface for the remote service. Useful for testing.
type ServerBinusPortType interface {
	// Check License
	CheckIzin(username string, password string, npwp string, kdizin string) (string, error)

	// Send IUTM From Client
	SendIUTM(username string, password string, xmlIUTM string) (string, error)

	// Send SIUP From Client
	SendSIUP(username string, password string, xmlSIUP string) (string, error)

	// Send STPW From Client
	SendSTPW(username string, password string, xmlSTPW string) (string, error)

	// Send TDG From Client
	SendTDG(username string, password string, xmlTDG string) (string, error)

	// Send TDP From Client
	SendTDP(username string, password string, xmlTDP string) (string, error)

	// Send Test Parameter From Client
	SendTEST(string string) (string, error)
}

// Operation wrapper for CheckIzin.
// OperationCheckIzinRequest was auto-generated from WSDL.
type OperationCheckIzinRequest struct {
	Username *string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password *string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	Npwp     *string `xml:"npwp,omitempty" json:"npwp,omitempty" yaml:"npwp,omitempty"`
	Kdizin   *string `xml:"kdizin,omitempty" json:"kdizin,omitempty" yaml:"kdizin,omitempty"`
}

// Operation wrapper for CheckIzin.
// OperationCheckIzinResponse was auto-generated from WSDL.
type OperationCheckIzinResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SendIUTM.
// OperationSendIUTMRequest was auto-generated from WSDL.
type OperationSendIUTMRequest struct {
	Username *string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password *string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	XmlIUTM  *string `xml:"xmlIUTM,omitempty" json:"xmlIUTM,omitempty" yaml:"xmlIUTM,omitempty"`
}

// Operation wrapper for SendIUTM.
// OperationSendIUTMResponse was auto-generated from WSDL.
type OperationSendIUTMResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SendSIUP.
// OperationSendSIUPRequest was auto-generated from WSDL.
type OperationSendSIUPRequest struct {
	Username *string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password *string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	XmlSIUP  *string `xml:"xmlSIUP,omitempty" json:"xmlSIUP,omitempty" yaml:"xmlSIUP,omitempty"`
}

// Operation wrapper for SendSIUP.
// OperationSendSIUPResponse was auto-generated from WSDL.
type OperationSendSIUPResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SendSTPW.
// OperationSendSTPWRequest was auto-generated from WSDL.
type OperationSendSTPWRequest struct {
	Username *string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password *string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	XmlSTPW  *string `xml:"xmlSTPW,omitempty" json:"xmlSTPW,omitempty" yaml:"xmlSTPW,omitempty"`
}

// Operation wrapper for SendSTPW.
// OperationSendSTPWResponse was auto-generated from WSDL.
type OperationSendSTPWResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SendTDG.
// OperationSendTDGRequest was auto-generated from WSDL.
type OperationSendTDGRequest struct {
	Username *string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password *string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	XmlTDG   *string `xml:"xmlTDG,omitempty" json:"xmlTDG,omitempty" yaml:"xmlTDG,omitempty"`
}

// Operation wrapper for SendTDG.
// OperationSendTDGResponse was auto-generated from WSDL.
type OperationSendTDGResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SendTDP.
// OperationSendTDPRequest was auto-generated from WSDL.
type OperationSendTDPRequest struct {
	Username *string `xml:"username,omitempty" json:"username,omitempty" yaml:"username,omitempty"`
	Password *string `xml:"password,omitempty" json:"password,omitempty" yaml:"password,omitempty"`
	XmlTDP   *string `xml:"xmlTDP,omitempty" json:"xmlTDP,omitempty" yaml:"xmlTDP,omitempty"`
}

// Operation wrapper for SendTDP.
// OperationSendTDPResponse was auto-generated from WSDL.
type OperationSendTDPResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// Operation wrapper for SendTEST.
// OperationSendTESTRequest was auto-generated from WSDL.
type OperationSendTESTRequest struct {
	String *string `xml:"string,omitempty" json:"string,omitempty" yaml:"string,omitempty"`
}

// Operation wrapper for SendTEST.
// OperationSendTESTResponse was auto-generated from WSDL.
type OperationSendTESTResponse struct {
	Return *string `xml:"return,omitempty" json:"return,omitempty" yaml:"return,omitempty"`
}

// serverBinusPortType implements the ServerBinusPortType interface.
type serverBinusPortType struct {
	cli *soap.Client
}

// Check License
func (p *serverBinusPortType) CheckIzin(username string, password string, npwp string, kdizin string) (string, error) {
	α := struct {
		M OperationCheckIzinRequest `xml:"tns:CheckIzin"`
	}{
		OperationCheckIzinRequest{
			&username,
			&password,
			&npwp,
			&kdizin,
		},
	}

	γ := struct {
		M OperationCheckIzinResponse `xml:"CheckIzinResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:CheckIzin#CheckIzin", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Return, nil
}

// Send IUTM From Client
func (p *serverBinusPortType) SendIUTM(username string, password string, xmlIUTM string) (string, error) {
	α := struct {
		M OperationSendIUTMRequest `xml:"tns:sendIUTM"`
	}{
		OperationSendIUTMRequest{
			&username,
			&password,
			&xmlIUTM,
		},
	}

	γ := struct {
		M OperationSendIUTMResponse `xml:"sendIUTMResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:sendIUTM#sendIUTM", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Return, nil
}

// Send SIUP From Client
func (p *serverBinusPortType) SendSIUP(username string, password string, xmlSIUP string) (string, error) {
	α := struct {
		M OperationSendSIUPRequest `xml:"tns:sendSIUP"`
	}{
		OperationSendSIUPRequest{
			&username,
			&password,
			&xmlSIUP,
		},
	}

	γ := struct {
		M OperationSendSIUPResponse `xml:"sendSIUPResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:sendSIUP#sendSIUP", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Return, nil
}

// Send STPW From Client
func (p *serverBinusPortType) SendSTPW(username string, password string, xmlSTPW string) (string, error) {
	α := struct {
		M OperationSendSTPWRequest `xml:"tns:sendSTPW"`
	}{
		OperationSendSTPWRequest{
			&username,
			&password,
			&xmlSTPW,
		},
	}

	γ := struct {
		M OperationSendSTPWResponse `xml:"sendSTPWResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:sendSTPW#sendSTPW", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Return, nil
}

// Send TDG From Client
func (p *serverBinusPortType) SendTDG(username string, password string, xmlTDG string) (string, error) {
	α := struct {
		M OperationSendTDGRequest `xml:"tns:sendTDG"`
	}{
		OperationSendTDGRequest{
			&username,
			&password,
			&xmlTDG,
		},
	}

	γ := struct {
		M OperationSendTDGResponse `xml:"sendTDGResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:sendTDG#sendTDG", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Return, nil
}

// Send TDP From Client
func (p *serverBinusPortType) SendTDP(username string, password string, xmlTDP string) (string, error) {
	α := struct {
		M OperationSendTDPRequest `xml:"tns:sendTDP"`
	}{
		OperationSendTDPRequest{
			&username,
			&password,
			&xmlTDP,
		},
	}

	γ := struct {
		M OperationSendTDPResponse `xml:"sendTDPResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:sendTDP#sendTDP", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Return, nil
}

// Send Test Parameter From Client
func (p *serverBinusPortType) SendTEST(string string) (string, error) {
	α := struct {
		M OperationSendTESTRequest `xml:"tns:sendTEST"`
	}{
		OperationSendTESTRequest{
			&string,
		},
	}

	γ := struct {
		M OperationSendTESTResponse `xml:"sendTESTResponse"`
	}{}
	if err := p.cli.RoundTripWithAction("urn:sendTEST#sendTEST", α, &γ); err != nil {
		return "", err
	}
	return *γ.M.Return, nil
}

//INSERT INTO qa(question, `group`) VALUES ("Setiap karyawan percaya bahwa dirinya dapat memberi dampak positif ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Kemampuan karyawan dipandang sebagai sumber utama keunggulan yang kompetitif", 1);
//INSERT INTO qa(question, `group`) VALUES ("Para pemimpin mempraktekkan apa yang mereka putuskan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Tidak sulit untuk mengkoordinasikan pekerjaan yang melibatkan unit – unit kerja yang berbeda", 1);
//INSERT INTO qa(question, `group`) VALUES ("Manajemen memastikan bahwa semua kegiatan kerja dilakukan secara transparan ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Ada arah dan tujuan jangka panjang dalam Perusahaan ini", 1);
//INSERT INTO qa(question, `group`) VALUES ("Tujuan organisasi disepakati bersama oleh seluruh karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Setiap karyawan disarankan untuk berani mengambil keputusan yang telah dipertimbangkan dengan matang pada saat yang tepat", 1);
//INSERT INTO qa(question, `group`) VALUES ("Ada penghargaan atau dorongan – dorongan yang sering dilakukan oleh pimpinan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Garis wewenang di Perusahaan selalu dijalankan sebagaimana mestinya ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Setiap karyawan mengetahui peluang untuk berkembang ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Selalu ada penghargaan bagi karyawan yang berprestasi cemerlang ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Komunikasi interpersonal karyawan dengan rekan – rekannya sangat intensif", 1);
//INSERT INTO qa(question, `group`) VALUES ("Dana yang ada sebagian besar untuk investasi pengembangan karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Pihak manajemen beranggapan bahwa karyawan memiliki kemampuan yang sama ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Perilaku manajemen mencerminkan keadilan ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Dirasakan adanya tuntutan untuk terus menerus untuk memperbaiki prestasi kerja pribadi maupun kelompok ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Selalu dibicarakan isu – isu kritis ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Adanya keyakinan bahwa Perusahaan merupakan satu – satunya harapan untuk hidup lebih baik  ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Aktivitas perusahaan sering melibatkan para tenaga professional", 1);
//INSERT INTO qa(question, `group`) VALUES ("Inisiatif dari karyawan selalu ditindaklanjuti oleh pihak manajemen", 1);
//INSERT INTO qa(question, `group`) VALUES ("Pihak manajemen menginvestasikan dana untuk pengembangan perusahaaan dalam memenuhi tantangan masa depan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Peraturan yang berlaku selalu mencerminkan visi dan misi Perusahaan ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Hambatan selalu diterjemahkan sebagai peluang", 1);
//INSERT INTO qa(question, `group`) VALUES ("Kerjasama antar unit kerja mendapat dukungan sepenuhnya ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Ada peraturan organisasi yang dengan jelas mengatur perilaku apa yang boleh dan tidak boleh dilakukan karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Karyawan sangat tanggap terhadap persaingan dengan Perusahaan lainnya", 1);
//INSERT INTO qa(question, `group`) VALUES ("Unit – unit kerja yang berbeda mudah bekerjasama untuk membuat perubahan yang sesuai dengan visi dan misi organisasi", 1);
//INSERT INTO qa(question, `group`) VALUES ("Para pemimpin manajemen sangat jelas dalam menetapkan sasaran – sasaran yang ingin diraih", 1);
//INSERT INTO qa(question, `group`) VALUES ("Apabila karyawan melakukan kesalahan maka akan segera ditegur. ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Sangat terasa adanya kebersamaan di dalam Perusahaan ini ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Kebijakan dalam Perusahaan ini dapat dimengerti dengan jelas oleh para karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Ada rasa saling mempercayai diantara karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Perusahaan ini selalu membuka kesempatan untuk berkomunikasi dengan atasan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Produktivitas kerja jarang terganggu karena kurang perencanaan dan koordinasi", 1);
//INSERT INTO qa(question, `group`) VALUES ("Perilaku karyawan yang tidak baik selalu diingatkan oleh pihak manajemen ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Saran dan usul selalu diperhatikan ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Setiap ada keputusan selallu dituangkan dalam aturan organisasi", 1);
//INSERT INTO qa(question, `group`) VALUES ("Adanya kebebasan untuk melakukan apapun yang ingin dilakukan ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Jarang terjadi pemborosan dana untuk keperluan non perusahaan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Karyawan bekerja layaknya seperti bagian dari satu tim ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Sangat mudah mencapai consensus, bahkan mengenai persoalan yang sangat sulit sekalipun", 1);
//INSERT INTO qa(question, `group`) VALUES ("Karyawan memiliki pemahaman mendalam tentang keinginnan dan kebutuhan kliennya. ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Para pemimpin Perusahaan memiliki pandangan yang jauh kedepan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Para karyawan selalu bekerja berdasarkan aturan yang ada ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Dirasakan adanya kesempatan untuk meningkatkan cara kerja maupun pengembangan technologi", 1);
//INSERT INTO qa(question, `group`) VALUES ("Setiap karyawan diberi tanggung jawab penuh untuk menjalankan tugas – tugasnya", 1);
//INSERT INTO qa(question, `group`) VALUES ("Setiap karyawan mengetahui dengan baik apa yang diharapkan oleh organisasi", 1);
//INSERT INTO qa(question, `group`) VALUES ("Standar kerja telah dibakukan ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Pihak manajemen selalu mengingatkan pentingnya aktualisasi visi dan misi", 1);
//INSERT INTO qa(question, `group`) VALUES ("Adanya jenjang karir yang jelas ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Dalam mencapai tujuan perusahaan dirasakan adanya kesempatan bagi karyawan untuk ikut serta secara aktif", 1);
//INSERT INTO qa(question, `group`) VALUES ("Secara periodic diadakan gathering untuk keluarga karyawan ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Selalu diingatkan oleh pihak manajemen akan perkembangan perusahaan lainnya", 1);
//INSERT INTO qa(question, `group`) VALUES ("Tugas – tugas yang diberikan mengandung tantangan ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Adanya pola kerja baku antar unit – unit kerja", 1);
//INSERT INTO qa(question, `group`) VALUES ("Dalam melaksanakan tugas, karyawan dapat melakukannya dengan cara kerja sendiri tanpa persetujuan atasan.", 1);
//INSERT INTO qa(question, `group`) VALUES ("Karyawan selalu dimotivasi untuk membuat Perusahaan ini terkenal ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Dirasakan adanya perhatian dari pihak manajemen terhadap persoalan pribadi. ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Investasi untuk meningkatkan pengetahuan dilakukan terus menerus ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Perilaku organisasi yang dijalankan oleh manajemen ini sangat konsisten dan dapat diprediksi", 1);
//INSERT INTO qa(question, `group`) VALUES ("Hampir semua pelaksanaan tugas terkoordinasi dengan rapi.√", 1);
//INSERT INTO qa(question, `group`) VALUES ("Pemikiran jangka pendek selalu selaras dengan cisi jangka panjang. ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Setiap rencana selalu dikoordinasikan dengan baik ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Terjalin kerjasama yang baik antara atasan dan bawahan, sehingga tugas – tugas berjalan dengan lancer. ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Setiap karyawan diberi kesempatan untuk menyampaikan pendapatnya", 1);
//INSERT INTO qa(question, `group`) VALUES ("Jarang terjadi pertentangan diantara karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Visi dan misi selalu disosialisasikan kepada seluruh karyawan ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Pertemuan antara karyawan dan pihak manajemen selalu diadakan secara periode ", 1);
//INSERT INTO qa(question, `group`) VALUES ("Dalam pelaksanaan tugas, karyawan diberi tanggungjawab penuh untuk melaksanakan tugas selama menggunakan cara kerja yang benar", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya memimpin dengan “melakukan” bukan hanya sekedar “berbicara”", 1);
//INSERT INTO qa(question, `group`) VALUES ("Pemimpin di perusahaan selalu menjadi panutan kebanggaan dan loyalitas bekeja di mata karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya dapat dijadikan acuan / panutan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya meminpin dengan contoh", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya menunjukkan bahwa banyak berharap dari karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya menuntut kinerja yang terbaik dari karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan di perusahaan selalu berhubungan dengan karyawan secara personal", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya bertindak tanpa mempedulikan perasaan karyawannya", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya menginspirasi orang lain dengan rencana di masa depan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan menggantikan tugas rekan yang tidak hadir", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan selalu bekerja optimal, agar hasil pekerjaan lebih baik dari teman-teman", 1);
//INSERT INTO qa(question, `group`) VALUES ("Pemimpin selalu mendorong ekspresi ide dari karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Pemimpin di perusahaan saya selalu berbicara optimis dengan antusias", 1);
//INSERT INTO qa(question, `group`) VALUES ("Pemimpin selalu mendorong motivasi karyawan dalam bekerja", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan menuangkan kreatifitas dan ide saya dalam meningkatkan produktivitas kerja", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan belajar dari teman yang telah berhasil untuk meningkatkan keterampilan saya", 1);
//INSERT INTO qa(question, `group`) VALUES ("Pemimpin di perusahaan selalu menekankan pentingnya tujuan serta komitmen bekerja", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya memahami ke mana organisasi akan dibawa", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya memupuk kerja sama", 1);
//INSERT INTO qa(question, `group`) VALUES ("Pemimpin selalu mendorong ekspresi ide dari karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya memiliki karyawan yang dapat diajak kerja sama untuk tujuan yang sama", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya memperlakukan karyawan tanpa melihat perasaan pribadi karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Saya percaya bahwa atasan saya memiliki integritas yang tinggi", 1);
//INSERT INTO qa(question, `group`) VALUES ("Pemimpin selalu mempertimbangkan segala kebutuhan karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya selalu mencari peluang baru untuk organisasi", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya memberi gambaran yang menarik tentang masa depan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Pemimpin selalu merangsang perspektif baru kepada karyawan", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya mampu mendorong karyawan untuk lebih aktif", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya memiliki ide yang membuat karyawan mampu memikirkan ide-ide baru", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan saya mampu membuat karyawan berpikir tentang sebuah masalah lama dengan cara yang baru", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan menolong rekan agar menjadi lebih produktif", 1);
//INSERT INTO qa(question, `group`) VALUES ("Atasan memberikan saran kepada organisasi sebagai perbaikan", 1);

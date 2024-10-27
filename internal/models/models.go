package models

type Flow struct {
	FlowStage   string
	Province    Option
	District    []Option
	Clinic      Option
	Hospital    []Option
	Doctor      []Option
	StartDate   string
	EndDate     string
	WantedHours []string
}

type Option struct {
	ID   string
	Name string
}

type LoginResponse struct {
	Success bool `json:"success"`
	Data    struct {
		JWT string `json:"jwt"`
	} `json:"data"`
}

type SearchAppointment struct {
	AksiyonID         string   `json:"aksiyonId"`
	Cinsiyet          string   `json:"cinsiyet"`
	MHRSHekimID       string   `json:"mhrsHekimId"`
	MHRSIlID          string   `json:"mhrsIlId"`
	MHRSIlceID        string   `json:"mhrsIlceId"`
	MHRSKlinikID      string   `json:"mhrsKlinikId"`
	MHRSKurumID       string   `json:"mhrsKurumId"`
	MuayeneYeriID     string   `json:"muayeneYeriId"`
	TumRandevular     bool     `json:"tumRandevular"`
	EkRandevu         bool     `json:"ekRandevu"`
	RandevuZamaniList []string `json:"randevuZamaniList"`
}

type AppointmentResponse struct {
	Data struct {
		Hastane []SingleAppointment `json:"hastane"`
		Semt    []SingleAppointment `json:"semt"`
	}
}

type SingleAppointment struct {
	Date     string   `json:"baslangicZamani"`
	Doctor   Doctor   `json:"hekim"`
	Hospital Hospital `json:"kurumBilgileri"`
}

type Hospital struct {
	Province   string  `json:"ilAdi"`
	District   string  `json:"ilceAdi"`
	Name       string  `json:"kurumAdi"`
	PublicName string  `json:"halkDilindekiAdi"`
	Latitude   float64 `json:"enlem"`
	Longitude  float64 `json:"boylam"`
}

type Doctor struct {
	Name    string `json:"ad"`
	Surname string `json:"soyad"`
	Gender  struct {
		Value string `json:"valText"`
	} `json:"cinsiyet"`
}

type SearchProvinceResponse struct {
	Value    int    `json:"value"`
	Text     string `json:"text"`
	Children []struct {
		Value int    `json:"value"`
		Text  string `json:"text"`
	}
}

type SearchDistrictResponse struct {
	Value string `json:"value"`
	Text  string `json:"text"`
}

type SearchClinicResponse struct {
	Data []NumericResponse `json:"data"`
}

type NumericResponse struct {
	Value int    `json:"value"`
	Text  string `json:"text"`
}

type SearchDoctorResponse struct {
	Data []NumericResponse `json:"data"`
}

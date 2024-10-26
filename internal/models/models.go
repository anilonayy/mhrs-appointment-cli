package models

type Flow struct {
	FlowStage string
	Province  Option
	District  Option
	Clinic    Option
	Hospital  Option
	Doctor    Option
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

type SearchSlot struct {
	AksiyonID         string   `json:"aksiyonId"`
	Cinsiyet          string   `json:"cinsiyet"`
	MHRSHekimID       int      `json:"mhrsHekimId"`
	MHRSIlID          int      `json:"mhrsIlId"`
	MHRSIlceID        int      `json:"mhrsIlceId"`
	MHRSKlinikID      int      `json:"mhrsKlinikId"`
	MHRSKurumID       int      `json:"mhrsKurumId"`
	MuayeneYeriID     int      `json:"muayeneYeriId"`
	TumRandevular     bool     `json:"tumRandevular"`
	EkRandevu         bool     `json:"ekRandevu"`
	RandevuZamaniList []string `json:"randevuZamaniList"`
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

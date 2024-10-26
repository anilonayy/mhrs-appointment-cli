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

type SearchSlotResponse struct {
	Lang     string   `json:"lang"`
	Success  bool     `json:"success"`
	Infos    []string `json:"infos"`
	Warnings []string `json:"warnings"`
	Errors   []struct {
		Kodu  string `json:"kodu"`
		Mesaj string `json:"mesaj"`
	} `json:"errors"`
	Data struct {
		Hastane       []string `json:"hastane"`
		Semt          []string `json:"semt"`
		Alternatif    []string `json:"alternatif"`
		SemtAramasi   bool     `json:"semtAramasi"`
		EkVar         bool     `json:"ekVar"`
		AcilacakSekme int      `json:"acilacakSekme"`
	} `json:"data"`
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
	Data []struct {
		Value int    `json:"value"`
		Text  string `json:"text"`
	} `json:"data"`
}

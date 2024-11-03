package models

type Flow struct {
	FlowStage   string
	Province    Option
	Districts   []Option
	Clinic      Option
	Hospitals   []Option
	Doctors     []Option
	StartDate   string
	EndDate     string
	SlotTime    string
	Appointment MakeAppointmentPayload
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
	StartDate struct {
		Date string `json:"date"`
	} `json:"baslangicZamaniStr"`
	Doctor   Doctor   `json:"hekim"`
	Hospital Hospital `json:"kurumBilgileri"`
}

type Hospital struct {
	ID         int     `json:"mhrsKurumId"`
	Province   string  `json:"ilAdi"`
	District   string  `json:"ilceAdi"`
	Name       string  `json:"kurumAdi"`
	PublicName string  `json:"halkDilindekiAdi"`
	Latitude   float64 `json:"enlem"`
	Longitude  float64 `json:"boylam"`
}

type Doctor struct {
	ID      int    `json:"mhrsHekimId"`
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

type SearchSlot struct {
	AksiyonID         string   `json:"aksiyonId"`
	Cinsiyet          string   `json:"cinsiyet"`
	MHRSHekimID       string   `json:"mhrsHekimId"`
	MHRSIlID          string   `json:"mhrsIlId"`
	MHRSKlinikID      string   `json:"mhrsKlinikId"`
	MHRSKurumID       string   `json:"mhrsKurumId"`
	MuayeneYeriID     string   `json:"muayeneYeriId"`
	TumRandevular     bool     `json:"tumRandevular"`
	EkRandevu         bool     `json:"ekRandevu"`
	RandevuZamaniList []string `json:"randevuZamaniList"`
}

type SearchSlotResponse struct {
	Data []struct {
		HekimSlotList []struct {
			MuayeneYeriSlotList []struct {
				SaatSlotList []struct {
					SlotList []struct {
						Id                int    `json:"id"`
						FkCetvelId        int64  `json:"fkCetvelId"`
						BaslangicZamani   string `json:"baslangicZamani"`
						BitisZamani       string `json:"bitisZamani"`
						CetvelIsKurallari struct {
							Id         int         `json:"id"`
							Cetvel     interface{} `json:"cetvel"`
							EnBuyukYas interface{} `json:"enBuyukYas"`
							EnKucukYas interface{} `json:"enKucukYas"`
							Cinsiyet   struct {
								Val     string `json:"val"`
								ValText string `json:"valText"`
							} `json:"cinsiyet"`
							CetvelOzellikModel interface{} `json:"cetvelOzellikModel"`
						} `json:"cetvelIsKurallari"`
						Slot struct {
							Id              int         `json:"id"`
							FkCetvelId      int64       `json:"fkCetvelId"`
							MhrsKurumId     int         `json:"mhrsKurumId"`
							MhrsKlinikId    int         `json:"mhrsKlinikId"`
							MhrsHekimId     int         `json:"mhrsHekimId"`
							MuayeneYeriId   int         `json:"muayeneYeriId"`
							FkAksiyonId     int         `json:"fkAksiyonId"`
							SlotHastaSayisi int         `json:"slotHastaSayisi"`
							RandevuSuresi   int         `json:"randevuSuresi"`
							Kullanim        int         `json:"kullanim"`
							KalanKullanim   int         `json:"kalanKullanim"`
							BaslangicZamani string      `json:"baslangicZamani"`
							BitisZamani     string      `json:"bitisZamani"`
							KayitZamani     interface{} `json:"kayitZamani"`
							AksiyonAdi      string      `json:"aksiyonAdi"`
						} `json:"slot"`
						Bos                       bool        `json:"bos"`
						IsKurali                  bool        `json:"isKurali"`
						Kapasite                  int         `json:"kapasite"`
						BosKapasite               int         `json:"bosKapasite"`
						Ek                        bool        `json:"ek"`
						UygunRandevuGecmisSlot    bool        `json:"uygunRandevuGecmisSlot"`
						RezerveTuruData           interface{} `json:"rezerveTuruData"`
						UzaktanDegerlendirmeVarmi bool        `json:"uzaktanDegerlendirmeVarmi"`
						BaslangicZamanStr         struct {
							Date         string `json:"date"`
							Tarih        string `json:"tarih"`
							Gun          string `json:"gun"`
							Saat         string `json:"saat"`
							GunAyGunIsmi string `json:"gunAyGunIsmi"`
							TarihAy      string `json:"tarihAy"`
							Zaman        string `json:"zaman"`
						} `json:"baslangicZamanStr"`
						BitisZamanStr struct {
							Date         string `json:"date"`
							Tarih        string `json:"tarih"`
							Gun          string `json:"gun"`
							Saat         string `json:"saat"`
							GunAyGunIsmi string `json:"gunAyGunIsmi"`
							TarihAy      string `json:"tarihAy"`
							Zaman        string `json:"zaman"`
						} `json:"bitisZamanStr"`
						KalanGunMesaj string `json:"kalanGunMesaj"`
					} `json:"slotList"`
					Saat                      string `json:"saat"`
					Bos                       bool   `json:"bos"`
					Ek                        bool   `json:"ek"`
					ToplamKapasite            int    `json:"toplamKapasite"`
					BosKapasite               int    `json:"bosKapasite"`
					UzaktanDegerlendirmeVarmi bool   `json:"uzaktanDegerlendirmeVarmi"`
					SaatStr                   string `json:"saatStr"`
				} `json:"saatSlotList"`
				MuayeneYeri struct {
					Id  int    `json:"id"`
					Adi string `json:"adi"`
				} `json:"muayeneYeri"`
				Bos   bool `json:"bos"`
				BosEk bool `json:"bosEk"`
			} `json:"muayeneYeriSlotList"`
			Hekim Doctor `json:"hekim"`
			Kurum struct {
				MhrsKurumId    int    `json:"mhrsKurumId"`
				MhrsAnaKurumId int    `json:"mhrsAnaKurumId"`
				KurumAdi       string `json:"kurumAdi"`
				KurumKisaAdi   string `json:"kurumKisaAdi"`
				KurumTurId     int    `json:"kurumTurId"`
				IlIlce         struct {
					MhrsIlId   int    `json:"mhrsIlId"`
					DkIlKodu   int    `json:"dkIlKodu"`
					IlAdi      string `json:"ilAdi"`
					MhrsIlceId int    `json:"mhrsIlceId"`
					DkIlceKodu int    `json:"dkIlceKodu"`
					IlceAdi    string `json:"ilceAdi"`
				} `json:"ilIlce"`
			} `json:"kurum"`
			Klinik struct {
				MhrsKlinikId  int           `json:"mhrsKlinikId"`
				MhrsKlinikAdi string        `json:"mhrsKlinikAdi"`
				KisaAdi       string        `json:"kisaAdi"`
				LcetvelTipi   int           `json:"lcetvelTipi"`
				YandalList    []interface{} `json:"yandalList"`
				AnadalList    []interface{} `json:"anadalList"`
			} `json:"klinik"`
		} `json:"hekimSlotList"`
		Gun           string `json:"gun"`
		Bos           bool   `json:"bos"`
		Kapasite      int    `json:"kapasite"`
		Kullanim      int    `json:"kullanim"`
		KalanKullanim int    `json:"kalanKullanim"`
		Cakisma       int    `json:"cakisma"`
		Istisna       int    `json:"istisna"`
		GunStr        struct {
			Date         string `json:"date"`
			Tarih        string `json:"tarih"`
			Gun          string `json:"gun"`
			Saat         string `json:"saat"`
			GunAyGunIsmi string `json:"gunAyGunIsmi"`
			TarihAy      string `json:"tarihAy"`
			Zaman        string `json:"zaman"`
		} `json:"gunStr"`
	} `json:"data"`
}

type MakeAppointmentPayload struct {
	FkSlotId        int    `json:"fkSlotId"`
	FkCetvelId      int64  `json:"fkCetvelId"`
	Yenidogan       bool   `json:"yenidogan"`
	MuayeneYeriId   int    `json:"muayeneYeriId"`
	BaslangicZamani string `json:"baslangicZamani"`
	BitisZamani     string `json:"bitisZamani"`
	RandevuNotu     string `json:"randevuNotu"`
}

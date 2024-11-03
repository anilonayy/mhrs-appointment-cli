package config

import (
	"os"
)

type URLConfig struct {
	LoginURL             string
	ProvinceSearchURL    string
	DistrictSearchURL    string
	ClinicSearchURL      string
	HospitalSearchURL    string
	DoctorSearchURL      string
	AppointmentSearchURL string
	SlotSearchURL        string
	CreateAppointmentURL string
	FileName             string
}

var urlConfig *URLConfig

func loadConfig() {
	urlConfig = &URLConfig{
		LoginURL:             os.Getenv("LOGIN_URL"),
		ProvinceSearchURL:    os.Getenv("PROVINCE_SEARCH_URL"),
		DistrictSearchURL:    os.Getenv("DISTRICT_SEARCH_URL"),
		ClinicSearchURL:      os.Getenv("CLINIC_SEARCH_URL"),
		HospitalSearchURL:    os.Getenv("HOSPITAL_SEARCH_URL"),
		DoctorSearchURL:      os.Getenv("DOCTOR_SEARCH_URL"),
		AppointmentSearchURL: os.Getenv("APPOINTMENT_SEARCH_URL"),
		SlotSearchURL:        os.Getenv("SLOT_SEARCH_URL"),
		CreateAppointmentURL: os.Getenv("CREATE_APPOINTMENT_URL"),
		FileName:             os.Getenv("FILE_NAME"),
	}
}

func GetConfig() *URLConfig {
	if urlConfig == nil {
		loadConfig()
	}

	if urlConfig.LoginURL == "" || urlConfig.ProvinceSearchURL == "" || urlConfig.DistrictSearchURL == "" ||
		urlConfig.ClinicSearchURL == "" || urlConfig.HospitalSearchURL == "" || urlConfig.DoctorSearchURL == "" ||
		urlConfig.AppointmentSearchURL == "" || urlConfig.SlotSearchURL == "" || urlConfig.CreateAppointmentURL == "" ||
		urlConfig.FileName == "" {
		panic("One or more URL configurations are not set")
	}

	return urlConfig
}

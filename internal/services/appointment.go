package services

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/anilonayy/mhrs-appointment-bot/config"
	"github.com/anilonayy/mhrs-appointment-bot/internal/constants"
	"github.com/anilonayy/mhrs-appointment-bot/internal/models"
	"github.com/anilonayy/mhrs-appointment-bot/pkg/resty"
)

func GetProvinces() (response []models.SearchProvinceResponse, err error) {
	err = WithSafeAuthorization(func() error {
		token, err := GetJWTToken()
		if err != nil {
			return err
		}

		resp, err := resty.GetClient().R().
			SetAuthToken(token).
			SetResult(&response).
			Get(config.GetConfig().ProvinceSearchURL)

		if err != nil {
			return err
		}

		if resp.IsError() {
			return errors.New(resp.String())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func SelectProvince(province *models.Option) (err error) {
	provinces, err := GetProvinces()
	if err != nil {
		return err
	}

	var provinceOptions []string

	for _, p := range provinces {
		provinceOptions = append(provinceOptions, p.Text)

		if len(p.Children) > 0 {
			for _, cP := range p.Children {
				provinceOptions = append(provinceOptions, cP.Text)
			}
		}
	}

	inputProvince := ""

	SelectOption("Please enter your province: ", provinceOptions, &inputProvince)

	for _, p := range provinces {
		if inputProvince == p.Text {
			(*province).Name = p.Text
			(*province).ID = strconv.Itoa(p.Value)

			break
		}

		if len(p.Children) > 0 {
			for _, cP := range p.Children {
				if inputProvince == cP.Text {
					(*province).Name = cP.Text
					(*province).ID = strconv.Itoa(cP.Value)

					break
				}
			}
		}

	}

	return err
}

func GetDistricts(province models.Option) (response []models.SearchDistrictResponse, err error) {
	err = WithSafeAuthorization(func() error {
		token, err := GetJWTToken()
		if err != nil {
			return err
		}

		resp, err := resty.GetClient().R().
			SetAuthToken(token).
			SetResult(&response).
			Get(config.GetConfig().DistrictSearchURL + province.ID)

		if err != nil {
			return err
		}

		if resp.IsError() {
			return errors.New(resp.String())
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func SelectDistrict(district *models.Option, province models.Option) (err error) {
	districts, err := GetDistricts(province)
	if err != nil {
		return err
	}

	districts = append([]models.SearchDistrictResponse{{Value: constants.NO_SELECTION, Text: constants.NO_SELECTION}}, districts...)

	districtOptions := make([]string, len(districts))

	for i, d := range districts {
		districtOptions[i] = d.Text
	}

	inputDistrict := ""
	SelectOption("Please enter your district: ", districtOptions, &inputDistrict)

	if inputDistrict == constants.NO_SELECTION {
		(*district).Name = constants.NO_SELECTION
		(*district).ID = "-1"

		return nil
	}

	for _, d := range districts {
		if inputDistrict == d.Text {
			(*district).Name = d.Text
			(*district).ID = d.Value

			return err
		}
	}

	return errors.New("invalid district")
}

func GetClinics(flow *models.Flow) (response models.SearchClinicResponse, err error) {
	err = WithSafeAuthorization(func() error {
		token, err := GetJWTToken()
		if err != nil {
			return err
		}

		resp, err := resty.GetClient().R().
			SetAuthToken(token).
			SetResult(&response).
			Get(fmt.Sprintf(config.GetConfig().ClinicSearchURL, flow.Province.ID, flow.District.ID))

		if err != nil {
			return err
		}

		if resp.IsError() {
			return errors.New(resp.String())
		}

		return nil
	})

	if err != nil {
		return response, err
	}

	return response, nil
}

func SelectClinic(flow *models.Flow) (err error) {
	clinics, err := GetClinics(flow)
	if err != nil {
		return err
	}

	var clinicOptions []string

	for _, p := range clinics.Data {
		clinicOptions = append(clinicOptions, p.Text)
	}

	inputProvince := ""

	SelectOption("Please enter your clinic: ", clinicOptions, &inputProvince)

	for _, p := range clinics.Data {
		if inputProvince == p.Text {
			(*flow).Clinic.Name = p.Text
			(*flow).Clinic.ID = strconv.Itoa(p.Value)

			break
		}
	}

	return nil
}

func GetHospitals(flow *models.Flow) (response models.SearchClinicResponse, err error) {
	err = WithSafeAuthorization(func() error {
		token, err := GetJWTToken()
		if err != nil {
			return err
		}

		resp, err := resty.GetClient().R().
			SetAuthToken(token).
			SetResult(&response).
			Get(fmt.Sprintf(config.GetConfig().HospitalSearchURL, flow.Province.ID, flow.District.ID, flow.Clinic.ID))

		if err != nil {
			return err
		}

		if resp.IsError() {
			return errors.New(resp.String())
		}

		return nil
	})

	if err != nil {
		return response, err
	}

	return response, nil
}

func SelectHospital(flow *models.Flow) (err error) {
	hospitals, err := GetHospitals(flow)
	if err != nil {
		return err
	}

	hospitals.Data = append([]models.NumericResponse{{Value: -1, Text: constants.NO_SELECTION}}, hospitals.Data...)

	var hospitalOptions []string

	for _, p := range hospitals.Data {
		hospitalOptions = append(hospitalOptions, p.Text)
	}

	inputProvince := ""

	SelectOption("Please enter your hospital: ", hospitalOptions, &inputProvince)

	for _, p := range hospitals.Data {
		if inputProvince == p.Text {
			(*flow).Hospital.Name = p.Text
			(*flow).Hospital.ID = strconv.Itoa(p.Value)

			break
		}
	}

	return nil
}

func GetDoctors(flow *models.Flow) (response models.SearchDoctorResponse, err error) {
	err = WithSafeAuthorization(func() error {
		token, err := GetJWTToken()
		if err != nil {
			return err
		}

		resp, err := resty.GetClient().R().
			SetAuthToken(token).
			SetResult(&response).
			Get(fmt.Sprintf(config.GetConfig().DoctorSearchURL, flow.Hospital.ID, flow.Clinic.ID))

		if err != nil {
			return err
		}

		if resp.IsError() {
			return errors.New(resp.String())
		}

		return nil
	})

	if err != nil {
		return response, err
	}

	return response, nil
}

func SelectDoctor(flow *models.Flow) (err error) {
	doctors, err := GetDoctors(flow)
	if err != nil {
		return err
	}

	doctors.Data = append([]models.NumericResponse{{Value: -1, Text: constants.NO_SELECTION}}, doctors.Data...)

	var doctorOptions []string

	for _, p := range doctors.Data {
		doctorOptions = append(doctorOptions, p.Text)
	}

	inputProvince := ""

	SelectOption("Please enter your doctor: ", doctorOptions, &inputProvince)

	for _, p := range doctors.Data {
		if inputProvince == p.Text {
			(*flow).Doctor.Name = p.Text
			(*flow).Doctor.ID = strconv.Itoa(p.Value)

			break
		}
	}

	return nil
}

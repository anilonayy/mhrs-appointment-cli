package appointment

import (
	"errors"
	"fmt"
	"github.com/anilonayy/mhrs-appointment-bot/internal/services/auth"
	"github.com/anilonayy/mhrs-appointment-bot/internal/ui"
	"strconv"
	"time"

	"github.com/anilonayy/mhrs-appointment-bot/config"
	"github.com/anilonayy/mhrs-appointment-bot/internal/constants"
	"github.com/anilonayy/mhrs-appointment-bot/internal/models"
	"github.com/anilonayy/mhrs-appointment-bot/pkg/resty"
)

func GetProvinces() (response []models.SearchProvinceResponse, err error) {
	err = auth.WithSafeAuthorization(func() error {
		token, err := auth.GetJWTToken()
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

func SelectProvince(flow *models.Flow) (err error) {
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

	var provinceSelection string
	ui.SelectOption("Please enter your province: ", provinceOptions, &provinceSelection)

	for _, p := range provinces {
		if provinceSelection == p.Text {
			(*flow).Province.Name = p.Text
			(*flow).Province.ID = strconv.Itoa(p.Value)

			break
		}

		if len(p.Children) > 0 {
			for _, cP := range p.Children {
				if provinceSelection == cP.Text {
					(*flow).Province.Name = cP.Text
					(*flow).Province.ID = strconv.Itoa(cP.Value)

					break
				}
			}
		}

	}

	return err
}

func GetDistricts(province models.Option) (response []models.SearchDistrictResponse, err error) {
	err = auth.WithSafeAuthorization(func() error {
		token, err := auth.GetJWTToken()
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

func SelectDistrict(flow *models.Flow) (err error) {
	districts, err := GetDistricts(flow.Province)
	if err != nil {
		return err
	}

	districts = append([]models.SearchDistrictResponse{{Value: constants.NO_SELECTION, Text: constants.NO_SELECTION}}, districts...)

	districtOptions := make([]string, len(districts))

	for i, d := range districts {
		districtOptions[i] = d.Text
	}

	var selectedDistricts []string

	ui.SelectOptions("Please select your districts: ", districtOptions, &selectedDistricts)

	for _, d := range districts {
		for _, inputDistrict := range selectedDistricts {
			if inputDistrict == constants.NO_SELECTION {
				(*flow).District = append((*flow).District, models.Option{Name: constants.NO_SELECTION, ID: constants.NO_SELECTION_CODE})

				return nil
			}

			if inputDistrict == d.Text {
				(*flow).District = append((*flow).District, models.Option{Name: d.Text, ID: d.Value})
			}
		}
	}

	return errors.New("invalid flow")
}

func GetClinics(flow *models.Flow) (response []models.NumericResponse, err error) {
	err = auth.WithSafeAuthorization(func() error {
		token, err := auth.GetJWTToken()
		if err != nil {
			return err
		}

		for _, district := range flow.District {
			var singleResponse models.SearchClinicResponse

			resp, err := resty.GetClient().R().
				SetAuthToken(token).
				SetResult(&singleResponse).
				Get(fmt.Sprintf(config.GetConfig().ClinicSearchURL, flow.Province.ID, district.ID))

			if err != nil {
				return err
			}

			if resp.IsError() {
				return errors.New(resp.String())
			}

			response = append(response, singleResponse.Data...)
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

	for _, p := range clinics {
		clinicOptions = append(clinicOptions, p.Text)
	}

	var clinicSelection string
	ui.SelectOption("Please enter your clinic: ", clinicOptions, &clinicSelection)

	for _, p := range clinics {
		if clinicSelection == p.Text {
			(*flow).Clinic.Name = p.Text
			(*flow).Clinic.ID = strconv.Itoa(p.Value)

			break
		}
	}

	return nil
}

func GetHospitals(flow *models.Flow) (response []models.NumericResponse, err error) {
	err = auth.WithSafeAuthorization(func() error {
		token, err := auth.GetJWTToken()
		if err != nil {
			return err
		}

		for _, district := range flow.District {
			var singleResponse models.SearchClinicResponse

			resp, err := resty.GetClient().R().
				SetAuthToken(token).
				SetResult(&singleResponse).
				Get(fmt.Sprintf(config.GetConfig().HospitalSearchURL, flow.Province.ID, district.ID, flow.Clinic.ID))

			if err != nil {
				return err
			}

			if resp.IsError() {
				return errors.New(resp.String())
			}

			response = append(response, singleResponse.Data...)
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

	hospitals = append([]models.NumericResponse{{Value: -1, Text: constants.NO_SELECTION}}, hospitals...)

	var hospitalOptions []string

	for _, p := range hospitals {
		hospitalOptions = append(hospitalOptions, p.Text)
	}

	var hospitalSelections []string

	ui.SelectOptions("Please select wanted hospitals: ", hospitalOptions, &hospitalSelections)

	for _, p := range hospitals {
		for _, hospitalSelections := range hospitalSelections {
			if hospitalSelections == constants.NO_SELECTION {
				(*flow).Hospital = append((*flow).Hospital, models.Option{Name: constants.NO_SELECTION, ID: constants.NO_SELECTION_CODE})
				break
			}

			if hospitalSelections == p.Text {
				(*flow).Hospital = append((*flow).Hospital, models.Option{Name: p.Text, ID: strconv.Itoa(p.Value)})
			}
		}
	}

	return nil
}

func GetDoctors(flow *models.Flow) (response []models.NumericResponse, err error) {
	err = auth.WithSafeAuthorization(func() error {
		token, err := auth.GetJWTToken()
		if err != nil {
			return err
		}

		for _, hospital := range flow.Hospital {
			var singleResponse struct {
				Data []models.NumericResponse `json:"data"`
			}

			resp, err := resty.GetClient().R().
				SetAuthToken(token).
				SetResult(&singleResponse).
				Get(fmt.Sprintf(config.GetConfig().DoctorSearchURL, hospital.ID, flow.Clinic.ID))

			if err != nil {
				return err
			}

			if resp.IsError() {
				return errors.New(resp.String())
			}

			response = append(response, singleResponse.Data...)
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

	doctors = append([]models.NumericResponse{{Value: -1, Text: constants.NO_SELECTION}}, doctors...)

	var doctorOptions []string

	for _, p := range doctors {
		doctorOptions = append(doctorOptions, p.Text)
	}

	var doctorSelections []string
	ui.SelectOptions("Please enter your doctor: ", doctorOptions, &doctorSelections)

	for _, p := range doctors {
		for _, doctor := range doctorSelections {
			if doctor == constants.NO_SELECTION {
				(*flow).Doctor = append((*flow).Doctor, models.Option{Name: constants.NO_SELECTION, ID: constants.NO_SELECTION_CODE})
				break
			}

			if doctor == p.Text {
				(*flow).Doctor = append((*flow).Doctor, models.Option{Name: p.Text, ID: strconv.Itoa(p.Value)})
			}
		}
	}

	return nil
}

func SelectDateRanges(flow *models.Flow) (err error) {
	ui.GetInput("Please enter wanted appointment start date (YYYY-MM-DD): ", &flow.StartDate)
	ui.GetInput("Please enter wanted appointment end date (YYYY-MM-DD): ", &flow.EndDate)

	if flow.StartDate == "" || flow.EndDate == "" {
		return nil
	}

	if _, err := time.Parse(time.DateOnly, flow.StartDate); err != nil {
		return errors.New("invalid date")
	}

	if _, err := time.Parse(time.DateOnly, flow.EndDate); err != nil {
		return errors.New("invalid date")
	}

	return nil
}
package appointment

import (
	"errors"
	"fmt"
	"github.com/anilonayy/mhrs-appointment-bot/internal/services/auth"
	"github.com/anilonayy/mhrs-appointment-bot/internal/ui"
	"github.com/anilonayy/mhrs-appointment-bot/internal/utils"
	"strconv"
	"time"

	"github.com/anilonayy/mhrs-appointment-bot/config"
	"github.com/anilonayy/mhrs-appointment-bot/internal/constants"
	"github.com/anilonayy/mhrs-appointment-bot/internal/models"
	"github.com/anilonayy/mhrs-appointment-bot/pkg/resty"
)

var (
	ErrNoAppointmentsFound = errors.New("no appointments found")
	ErrDateRangeExpired    = errors.New("date range expired")
)

func getProvinces() (response []models.SearchProvinceResponse, err error) {
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
	provinces, err := getProvinces()
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

func getDistricts(province models.Option) (response []models.SearchDistrictResponse, err error) {
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
	districts, err := getDistricts(flow.Province)
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
				(*flow).Districts = append((*flow).Districts, models.Option{Name: constants.NO_SELECTION, ID: constants.NO_SELECTION_CODE})

				return nil
			}

			if inputDistrict == d.Text {
				(*flow).Districts = append((*flow).Districts, models.Option{Name: d.Text, ID: d.Value})
			}
		}
	}

	if len((*flow).Districts) == 0 {
		return errors.New("no district selected")
	}

	return nil
}

func getClinics(flow *models.Flow) (response []models.NumericResponse, err error) {
	err = auth.WithSafeAuthorization(func() error {
		token, err := auth.GetJWTToken()
		if err != nil {
			return err
		}

		for _, district := range flow.Districts {
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
	clinics, err := getClinics(flow)
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

func getHospitals(flow *models.Flow) (response []models.NumericResponse, err error) {
	err = auth.WithSafeAuthorization(func() error {
		token, err := auth.GetJWTToken()
		if err != nil {
			return err
		}

		for _, district := range flow.Districts {
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
	hospitals, err := getHospitals(flow)
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
				(*flow).Hospitals = append((*flow).Hospitals, models.Option{Name: constants.NO_SELECTION, ID: constants.NO_SELECTION_CODE})
				break
			}

			if hospitalSelections == p.Text {
				(*flow).Hospitals = append((*flow).Hospitals, models.Option{Name: p.Text, ID: strconv.Itoa(p.Value)})
			}
		}
	}

	return nil
}

func getDoctors(flow *models.Flow) (response []models.NumericResponse, err error) {
	err = auth.WithSafeAuthorization(func() error {
		token, err := auth.GetJWTToken()
		if err != nil {
			return err
		}

		for _, hospital := range flow.Hospitals {
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
	doctors, err := getDoctors(flow)
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
				(*flow).Doctors = append((*flow).Doctors, models.Option{Name: constants.NO_SELECTION, ID: constants.NO_SELECTION_CODE})
				break
			}

			if doctor == p.Text {
				(*flow).Doctors = append((*flow).Doctors, models.Option{Name: p.Text, ID: strconv.Itoa(p.Value)})
			}
		}
	}

	return nil
}

func SelectDateRanges(flow *models.Flow) (err error) {
	ui.GetInput("Please enter wanted appointment start date (YYYY-MM-DD): ", &flow.StartDate)
	ui.GetInput("Please enter wanted appointment end date (YYYY-MM-DD): ", &flow.EndDate)

	if flow.StartDate == "" {
		flow.StartDate = time.Now().Format(time.DateOnly)
	}

	if flow.EndDate == "" {
		flow.EndDate = time.Now().AddDate(0, 0, 30).Format(time.DateOnly)
	}

	if _, err := time.Parse(time.DateOnly, flow.StartDate); err != nil {
		return errors.New("invalid date")
	}

	if _, err := time.Parse(time.DateOnly, flow.EndDate); err != nil {
		return errors.New("invalid date")
	}

	return nil
}

func SelectSlotTimes(flow *models.Flow) (err error) {
	slots := []string{constants.MORNING_SLOT, constants.AFTERNOON_SLOT}

	ui.SelectOption("Please select your slot time: ", slots, &flow.SlotTime)

	return nil
}

func getAppointments(flow *models.Flow) ([]models.SingleAppointment, error) {
	var response models.AppointmentResponse
	var payload = models.SearchAppointment{
		AksiyonID:         "200",
		Cinsiyet:          "F",
		MHRSIlID:          flow.Province.ID,
		MHRSIlceID:        constants.NO_SELECTION_CODE,
		MHRSKlinikID:      flow.Clinic.ID,
		MHRSKurumID:       constants.NO_SELECTION_CODE,
		MuayeneYeriID:     constants.NO_SELECTION_CODE,
		MHRSHekimID:       constants.NO_SELECTION_CODE,
		TumRandevular:     false,
		EkRandevu:         false,
		RandevuZamaniList: []string{},
	}

	err := auth.WithSafeAuthorization(func() error {
		token, err := auth.GetJWTToken()
		if err != nil {
			return err
		}

		resp, err := resty.GetClient().R().
			SetAuthToken(token).
			SetBody(payload).
			SetResult(&response).
			Post(config.GetConfig().AppointmentSearchURL)

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

	if len(response.Data.Hastane) == 0 {
		return nil, errors.New("no appointments found")
	}

	return response.Data.Hastane, nil
}

func filterAppointments(appointments []models.SingleAppointment, flow *models.Flow) (filteredAppointments []models.SingleAppointment) {
	for _, appointment := range appointments {
		var (
			doctorFilter   = false
			hospitalFilter = false
			dateFilter     = false
		)

		if len(flow.Doctors) > 0 && flow.Doctors[0].ID != constants.NO_SELECTION_CODE {
			for _, doctor := range flow.Doctors {
				if strconv.Itoa(appointment.Doctor.ID) == doctor.ID {
					doctorFilter = true
				}
			}
		} else {
			doctorFilter = true
		}

		if flow.Hospitals[0].ID != constants.NO_SELECTION_CODE {
			for _, hospital := range flow.Hospitals {
				if appointment.Hospital.Name == hospital.Name {
					hospitalFilter = true
				}
			}
		} else {
			hospitalFilter = true
		}

		startDate, _ := time.Parse(time.DateOnly, flow.StartDate)
		endDate, _ := time.Parse(time.DateOnly, flow.EndDate)
		appointmentDate, appointmentErr := time.Parse(time.DateTime, appointment.StartDate.Date)

		if appointmentDate.After(startDate) && appointmentDate.Before(endDate) {
			dateFilter = true
		}

		if appointmentErr != nil {
			dateFilter = false
		}

		if doctorFilter && hospitalFilter && dateFilter {
			filteredAppointments = append(filteredAppointments, appointment)
		}

	}

	return filteredAppointments
}

func getSlots(flow *models.Flow) (response models.SearchSlotResponse, err error) {
	var payload = models.SearchSlot{
		AksiyonID:         "200",
		Cinsiyet:          "F",
		MHRSIlID:          flow.Province.ID,
		MHRSKlinikID:      flow.Clinic.ID,
		MHRSKurumID:       flow.Hospitals[0].ID,
		MuayeneYeriID:     constants.NO_SELECTION_CODE,
		MHRSHekimID:       flow.Doctors[0].ID,
		TumRandevular:     false,
		EkRandevu:         false,
		RandevuZamaniList: []string{},
	}

	err = auth.WithSafeAuthorization(func() error {
		token, err := auth.GetJWTToken()
		if err != nil {
			return err
		}

		resp, err := resty.GetClient().R().
			SetAuthToken(token).
			SetResult(&response).
			SetBody(payload).
			Post(config.GetConfig().SlotSearchURL)

		if err != nil {
			return err
		}

		if resp.IsError() {
			if utils.CheckNeedAdvancedExpertError(resp.String()) {
				panic("Need advanced expert advise to make appointment for this clinic!")
			}
		}

		return nil
	})

	if err != nil {
		return response, err
	}

	return response, nil
}

func Do(flow *models.Flow) error {
	isDateExpired, err := utils.CheckDateRangeExpire(flow.EndDate)
	if err != nil {
		return err
	}

	if isDateExpired {
		return ErrDateRangeExpired
	}

	appointments, err := getAppointments(flow)
	if err != nil {
		return err
	}

	if len(appointments) == 0 {
		return ErrNoAppointmentsFound
	}

	appointments = filterAppointments(appointments, flow)

	if len(appointments) == 0 {
		return ErrNoAppointmentsFound
	}

	ui.PrintInfoMessage(fmt.Sprintf("Found %d appointments, slots searching..", len(appointments)))

	for _, appointment := range appointments {
		(*flow).Doctors = []models.Option{{Name: appointment.Doctor.Name, ID: strconv.Itoa(appointment.Doctor.ID)}}
		(*flow).Hospitals = []models.Option{{Name: appointment.Hospital.Name, ID: strconv.Itoa(appointment.Hospital.ID)}}

		slot, err := getSlots(flow)
		if err != nil {
			return err
		}

		for _, slotList := range slot.Data {
			doctorSlotList := slotList.HekimSlotList[0]
			examinationPlace := doctorSlotList.MuayeneYeriSlotList[0]

			if examinationPlace.Bos == false {
				continue
			}

			for _, hourSlotList := range examinationPlace.SaatSlotList {
				if hourSlotList.Bos == false {
					continue
				}

				for _, hour := range hourSlotList.SlotList {
					if hour.Bos == false {
						continue
					}

					ok, err := utils.CheckTimeSlot(flow.SlotTime, hour.BaslangicZamani)
					if err != nil {
						return err
					}

					if ok {
						infoMessage := fmt.Sprintf(
							"Found available slot: %s\nDoctor: %s\nHospital: %s\nExamination Place: %s\nSlot: %s",
							hour.BaslangicZamani,
							appointment.Doctor.Name+" "+appointment.Doctor.Surname,
							appointment.Hospital.Name,
							examinationPlace.MuayeneYeri.Adi,
							hour.BaslangicZamani)
						ui.PrintInfoMessage(infoMessage)

						(*flow).Appointment = models.MakeAppointmentPayload{
							BaslangicZamani: hour.BaslangicZamani,
							BitisZamani:     hour.BitisZamani,
							FkCetvelId:      hour.FkCetvelId,
							FkSlotId:        hour.Id,
							MuayeneYeriId:   examinationPlace.MuayeneYeri.Id,
							RandevuNotu:     "",
							Yenidogan:       false,
						}

						if err := makeAppointment(flow); err != nil {
							return err
						}

						return nil
					}
				}
			}
		}
	}

	return ErrNoAppointmentsFound
}

func makeAppointment(flow *models.Flow) error {
	var response struct {
		Success bool `json:"success"`
		Errors  []struct {
			Message string `json:"mesaj"`
			Code    string `json:"kodu"`
		}
		Infos []struct {
			Message string `json:"mesaj"`
			Code    string `json:"kodu"`
		}
	}
	err := auth.WithSafeAuthorization(func() error {
		token, err := auth.GetJWTToken()
		if err != nil {
			return err
		}

		resp, err := resty.GetClient().R().
			SetAuthToken(token).
			SetBody((*flow).Appointment).
			SetResult(&response).
			Post(config.GetConfig().CreateAppointmentURL)

		if err != nil {
			return err
		}

		if resp.IsError() {
			return errors.New(resp.String())
		}

		return nil
	})

	if err != nil {
		return err
	}

	if !response.Success {
		return errors.New("appointment could not be created here is the reason:" + response.Errors[0].Message)
	}

	ui.PrintInfoMessage("Appointment created successfully! Message:" + response.Infos[0].Message)

	return nil
}

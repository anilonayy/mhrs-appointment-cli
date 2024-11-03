package flow

import (
	"errors"
	"os"
	"time"

	"github.com/anilonayy/mhrs-appointment-bot/internal/models"
	"github.com/anilonayy/mhrs-appointment-bot/internal/services/appointment"
	"github.com/anilonayy/mhrs-appointment-bot/internal/ui"
	"github.com/anilonayy/mhrs-appointment-bot/internal/utils"
)

func SetFlowStage(flow *models.Flow) {
	switch (*flow).FlowStage {
	case "1": // Select Province
		if err := appointment.SelectProvince(flow); err != nil {
			panic(err)
		}

		(*flow).FlowStage = "2"

		SetFlowStage(flow)
	case "2": // Select Districts (Optional)
		if err := appointment.SelectDistrict(flow); err != nil {
			panic(err)
		}

		(*flow).FlowStage = "3"
		SetFlowStage(flow)
	case "3": // Select Clinic
		if err := appointment.SelectClinic(flow); err != nil {
			panic(err)
		}

		(*flow).FlowStage = "4"
		SetFlowStage(flow)
	case "4": // Select Hospitals (Optional)
		if err := appointment.SelectHospital(flow); err != nil {
			panic(err)
		}

		if utils.HasDefaultSelection((*flow).Hospitals) {
			ui.PrintInfoMessage("No hospital selected. Skipping doctor selection.")

			(*flow).FlowStage = "6"
			SetFlowStage(flow)

			break
		}

		(*flow).FlowStage = "5"
		SetFlowStage(flow)
	case "5": // Select Doctors (Optional)
		if err := appointment.SelectDoctor(flow); err != nil {
			panic(err)
		}

		(*flow).FlowStage = "6"
		SetFlowStage(flow)
	case "6": // Select StartDate Range (Optional)
		ui.PrintInfoMessage("If you enter empty dates, the system will automatically select the next available date.")
		if err := appointment.SelectDateRanges(flow); err != nil {
			panic(err)
		}

		(*flow).FlowStage = "7"
		SetFlowStage(flow)
	case "7": // Select slot times
		if err := appointment.SelectSlotTimes(flow); err != nil {
			panic(err)
		}

		(*flow).FlowStage = "8"
		SetFlowStage(flow)
	case "8": // Search for appointments
		err := appointment.Do(flow)

		if err != nil {
			if errors.Is(err, appointment.ErrDateRangeExpired) {
				ui.PrintInfoMessage("Selected date range is expired.")

				os.Exit(0)
			} else if errors.Is(err, appointment.ErrNoAppointmentsFound) {
				(*flow).FlowStage = "8"

				ui.PrintInfoMessage("No available slot found. Retrying in 15 minutes. Time: " + time.Now().Format(time.DateTime))
				time.Sleep(time.Minute * 15)

				SetFlowStage(flow)
			} else {
				panic(err)
			}
		}

		(*flow).FlowStage = "9"
		SetFlowStage(flow)
	case "9": // Show success message and exit
		ui.PrintInfoMessage("Appointment is successfully created.")

		os.Exit(0)
	}
}

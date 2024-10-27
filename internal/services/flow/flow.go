package flow

import (
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
	case "2": // Select District (Optional)
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
	case "4": // Select Hospital (Optional)
		if err := appointment.SelectHospital(flow); err != nil {
			panic(err)
		}

		if utils.HasDefaultSelection((*flow).Hospital) {
			ui.PrintInfoMessage("No hospital selected. Skipping doctor selection.")

			(*flow).FlowStage = "6"
			SetFlowStage(flow)

			break
		}

		(*flow).FlowStage = "5"
		SetFlowStage(flow)
	case "5": // Select Doctor
		if err := appointment.SelectDoctor(flow); err != nil {
			panic(err)
		}

		(*flow).FlowStage = "6"
		SetFlowStage(flow)
	case "6": // Select Date Range
		ui.PrintInfoMessage("If you enter empty dates, the system will automatically select the next available date.")
		if err := appointment.SelectDateRanges(flow); err != nil {
			panic(err)
		}

		(*flow).FlowStage = "7"
		SetFlowStage(flow)
		//case "7": // Select Slot Times
	}
}

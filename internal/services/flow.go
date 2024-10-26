package services

import (
	"github.com/anilonayy/mhrs-appointment-bot/internal/models"
)

func SetFlowStage(flow *models.Flow) {
	switch (*flow).FlowStage {
	case "1": // Select Province
		if err := SelectProvince(&flow.Province); err != nil {
			panic(err)
		}

		(*flow).FlowStage = "2"

		SetFlowStage(flow)
	case "2": // Select District (Optional)
		if err := SelectDistrict(&flow.District, flow.Province); err != nil {
			panic(err)
		}

		(*flow).FlowStage = "3"
		SetFlowStage(flow)
	case "3": // Select Clinic
		if err := SelectClinic(flow); err != nil {
			panic(err)
		}

		(*flow).FlowStage = "4"
		SetFlowStage(flow)
	case "4": // Select Hospital (Optional)
		if err := SelectHospital(flow); err != nil {
			panic(err)
		}

		(*flow).FlowStage = "5"
		SetFlowStage(flow)
	}
}

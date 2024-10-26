package services

import (
	"fmt"
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
	case "2": // Select District
		if err := SelectDistrict(&flow.District, flow.Province); err != nil {
			fmt.Println(err)

			SetFlowStage(flow)
		}

		(*flow).FlowStage = "3"
	}
}

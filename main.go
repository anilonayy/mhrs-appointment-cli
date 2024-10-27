package main

import (
	"fmt"
	"os"

	"github.com/anilonayy/mhrs-appointment-bot/internal/models"
	flowservice "github.com/anilonayy/mhrs-appointment-bot/internal/services/flow"
	"github.com/anilonayy/mhrs-appointment-bot/internal/ui"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	flow := models.Flow{
		FlowStage: "1",
	}
	menuStage := ""

	ui.SelectMenu(&menuStage)

	switch menuStage {
	case "1": // Search Appointment
		flowservice.SetFlowStage(&flow)
		break
	case "3":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		ui.SelectMenu(&menuStage)
	}
}

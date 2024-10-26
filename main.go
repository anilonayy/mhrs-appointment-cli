package main

import (
	"fmt"
	"os"

	"github.com/anilonayy/mhrs-appointment-bot/internal/models"
	"github.com/anilonayy/mhrs-appointment-bot/internal/services"
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

	services.SelectMenu(&menuStage)

	switch menuStage {
	case "1": // Search Appointment
		services.SetFlowStage(&flow)
		break
	case "3":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		services.SelectMenu(&menuStage)
	}
}

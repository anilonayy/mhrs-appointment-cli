package main

import (
	"github.com/anilonayy/mhrs-appointment-bot/internal/models"
	flowservice "github.com/anilonayy/mhrs-appointment-bot/internal/services/flow"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if os.Getenv("MHRS_USERNAME") == "" || os.Getenv("MHRS_PASSWORD") == "" {
		panic("MHRS_USERNAME or MHRS_PASSWORD is not set")
	}
}

func main() {
	flow := models.Flow{
		FlowStage: "1",
	}

	flowservice.SetFlowStage(&flow)
}

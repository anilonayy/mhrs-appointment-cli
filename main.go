package main

import (
	"github.com/anilonayy/mhrs-appointment-bot/internal/models"
	flowservice "github.com/anilonayy/mhrs-appointment-bot/internal/services/flow"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	flow := models.Flow{
		FlowStage: "1",
	}

	flowservice.SetFlowStage(&flow)
}

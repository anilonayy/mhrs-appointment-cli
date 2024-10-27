package utils

import (
	"github.com/anilonayy/mhrs-appointment-bot/internal/constants"
	"github.com/anilonayy/mhrs-appointment-bot/internal/models"
)

func HasDefaultSelection(opts []models.Option) bool {
	for _, v := range opts {
		if v.ID == constants.NO_SELECTION_CODE {
			return true
		}
	}

	return false
}

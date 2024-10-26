package utils

import (
	"encoding/json"
	"github.com/anilonayy/mhrs-appointment-bot/internal/constants"
)

func CheckUnauthorizedError(resp string) bool {
	data := make(map[string]any)

	if err := json.Unmarshal([]byte(resp), &data); err != nil {
		return false
	}

	if data["errors"] != nil {
		if data["errors"].([]any)[0].(map[string]any)["kodu"] == constants.UNAUTHORIZED_CODE {
			return true
		}
	}

	return false
}

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
		code := data["errors"].([]any)[0].(map[string]any)["kodu"]

		if code == constants.UNAUTHORIZED_CODE || code == constants.ANOTHER_LOGIN_CODE {
			return true
		}
	}

	return false
}

func CheckNeedAdvancedExpertError(resp string) bool {
	data := make(map[string]any)

	if err := json.Unmarshal([]byte(resp), &data); err != nil {
		return false
	}

	if data["errors"] != nil {
		code := data["errors"].([]any)[0].(map[string]any)["kodu"]

		if code == constants.NEED_ADVANCED_EXPERT {
			return true
		}
	}

	return false
}

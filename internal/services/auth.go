package services

import (
	"encoding/json"
	"errors"
	"github.com/anilonayy/mhrs-appointment-bot/config"
	"github.com/anilonayy/mhrs-appointment-bot/internal/constants"
	"github.com/anilonayy/mhrs-appointment-bot/internal/utils"
	"github.com/anilonayy/mhrs-appointment-bot/pkg/resty"
	"github.com/anilonayy/mhrs-appointment-bot/pkg/retry"
	"os"
	"strconv"
	"time"
)

func UpdateJWTToken() error {
	var response struct {
		Success bool `json:"success"`
		Data    struct {
			JWT string `json:"jwt"`
		} `json:"data"`
	}

	payload := map[string]any{
		"kullaniciAdi": os.Getenv("MHRS_USERNAME"),
		"parola":       os.Getenv("MHRS_PASSWORD"),
		"islemKanali":  "VATANDAS_RESPONSIVE",
		"girisTipi":    "PAROLA",
		"captchaKey":   nil,
	}

	resp, err := resty.GetClient().R().
		SetBody(payload).
		SetResult(&response).
		Post(config.GetConfig().LoginURL)

	if err != nil {
		return err
	}

	if resp.IsError() {
		return errors.New("Error while updating JWT token: " + resp.String())
	}

	fileData, err := utils.ReadFromFile()
	if err != nil {
		return err
	}
	fileData["jwt"] = response.Data.JWT
	fileData["expires"] = strconv.FormatInt(time.Now().Add(time.Minute*10).Unix(), 10)

	return utils.SaveToFile(fileData)
}

func GetJWTToken() (string, error) {
	data, err := utils.ReadFromFile()
	if err != nil {
		return "", err
	}

	expiresInSec, err := strconv.Atoi(data["expires"])
	if err != nil {
		return "", err
	}

	if data["jwt"] == "" || data["expires"] == "" || time.Now().Unix() > int64(expiresInSec) {
		err := UpdateJWTToken()
		if err != nil {
			return "", err
		}

		return GetJWTToken()
	}

	return data["jwt"], err
}

func CheckUnauthorizedError(resp string) bool {
	data := make(map[string]any)

	if err := json.Unmarshal([]byte(resp), &data); err != nil {
		return false
	}

	if data["Errors"] != nil {
		if data["Errors"].([]any)[0].(map[string]any)["Kodu"] == constants.UNAUTHORIZED_CODE {
			return true
		}
	}

	return false
}

func WithSafeAuthorization(callback func() error) (err error) {
	err = retry.Do(func() error {
		err := callback()

		if err != nil && utils.CheckUnauthorizedError(err.Error()) {
			if updateErr := UpdateJWTToken(); updateErr != nil {
				return updateErr
			}
		}

		return err
	})

	return err
}

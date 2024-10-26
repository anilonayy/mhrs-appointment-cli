package utils

import (
	"encoding/json"
	"errors"
	"github.com/anilonayy/mhrs-appointment-bot/config"
	"os"
)

func SaveToFile(data any) error {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(config.GetConfig().FileName, jsonByte, 0644)

	if errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(config.GetConfig().FileName)
		if err != nil {
			return err
		}

		return SaveToFile(data)
	}

	return nil
}

func ReadFromFile() (map[string]string, error) {
	fileName := config.GetConfig().FileName
	data := make(map[string]string)

	file, err := os.ReadFile(fileName)
	if errors.Is(err, os.ErrNotExist) {
		if _, err := os.Create(fileName); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	if len(file) == 0 {
		return data, nil
	}

	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

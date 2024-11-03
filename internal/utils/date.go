package utils

import (
	"fmt"
	"strings"
	"time"
)

func CheckTimeSlot(slot string, checkTimeStr string) (bool, error) {
	timeRange := strings.Split(slot, "-")
	if len(timeRange) != 2 {
		return false, fmt.Errorf("invalid slot format, expected 'start-end'")
	}

	start, err := time.Parse(time.Kitchen, timeRange[0])
	if err != nil {
		return false, fmt.Errorf("invalid start time format: %w", err)
	}
	end, err := time.Parse(time.Kitchen, timeRange[1])
	if err != nil {
		return false, fmt.Errorf("invalid end time format: %w", err)
	}

	checkTime, err := time.Parse(time.DateTime, checkTimeStr)
	if err != nil {
		return false, fmt.Errorf("invalid check time format: %w", err)
	}

	checkTime = time.Date(0, 1, 1, checkTime.Hour(), checkTime.Minute(), 0, 0, time.UTC)

	return checkTime.After(start) && checkTime.Before(end), nil
}

func CheckDateRangeExpire(endDate string) (bool, error) {
	end, err := time.Parse(time.DateOnly, endDate)
	if err != nil {
		return false, fmt.Errorf("invalid end date format: %w", err)
	}

	return time.Now().After(end), nil
}

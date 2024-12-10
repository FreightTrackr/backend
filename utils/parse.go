package utils

import "time"

func ParseDate(dateStr string, isEndDate bool) (time.Time, error) {
	if dateStr == "" {
		if isEndDate {
			return time.Now(), nil
		}
		return time.Time{}, nil
	}

	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

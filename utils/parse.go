package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

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

func WriteJSONResponse(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	jsonData, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

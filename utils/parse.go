package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/FreightTrackr/backend/models"
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

func ParseBody(w http.ResponseWriter, r *http.Request, v interface{}) {
	err := json.NewDecoder(r.Body).Decode(&v)

	if err != nil {
		WriteJSONResponse(w, http.StatusBadRequest, models.Pesan{
			Status:  http.StatusBadRequest,
			Message: "Error parsing application/json: " + err.Error(),
		})
		return
	}
}

func GetUrlQuery(r *http.Request, queryKey string, defaultValue string) string {
	query := r.URL.Query()
	v := query.Get(queryKey)
	if v == "" {
		return defaultValue
	}
	return v
}

func LogRequestDetails(start time.Time, message string) {
	duration := time.Since(start)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Response Time: %.2f ms | Alloc = %.2f MiB | TotalAlloc = %.2f MiB | Sys = %.2f MiB | NumGC = %v | Message: %s\n",
		duration.Seconds()*1000,
		float64(m.Alloc)/1024/1024,
		float64(m.TotalAlloc)/1024/1024,
		float64(m.Sys)/1024/1024,
		m.NumGC,
		message)
}

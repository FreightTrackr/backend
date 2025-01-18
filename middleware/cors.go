package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
	"https://freighttrackr.github.io",
	"https://freighttrackr.befous.com",
}

var headers = []string{
	"Content-Type",
	"Authorization",
	"Accept",
	"Origin",
}

var FiberCors = cors.Config{
	AllowOrigins:     strings.Join(origins, ","),
	AllowHeaders:     strings.Join(headers, ","),
	AllowMethods:     "GET, POST, PUT, PATCH, DELETE",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
	MaxAge:           int((2 * time.Hour).Seconds()),
}

var FiberLocalCors = cors.Config{
	AllowOrigins:  "*",
	AllowHeaders:  strings.Join(headers, ","),
	AllowMethods:  "GET, POST, PUT, PATCH, DELETE",
	ExposeHeaders: "Content-Length",
	MaxAge:        int((2 * time.Hour).Seconds()),
}

func StdCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" && contains(origins, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
			w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Max-Age", "7200")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func StdLocalCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		w.Header().Set("Access-Control-Max-Age", "7200")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

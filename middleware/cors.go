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
}

var FiberCors = cors.Config{
	AllowOrigins:     strings.Join(origins, ","),
	AllowHeaders:     strings.Join(headers, ","),
	AllowMethods:     "GET, POST, PUT, PATCH, DELETE",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
	MaxAge:           int((2 * time.Hour).Seconds()),
}

func StdCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", strings.Join(origins, ","))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		next.ServeHTTP(w, r)
	})
}

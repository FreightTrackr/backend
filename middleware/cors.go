package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
	"https://freighttrackr.github.io",
	"https://freighttrackr.befous.com",
}

var headers = []string{
	"Origin",
	"Content-Type",
	"Accept",
	"Authorization",
}

var Cors = cors.Config{
	AllowOrigins:     strings.Join(origins, ","),
	AllowHeaders:     strings.Join(headers, ","),
	AllowMethods:     "GET, POST, PUT, PATCH, DELETE",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
	MaxAge:           int((2 * time.Hour).Seconds()),
}

var CorsAllowAll = cors.Config{
	AllowOrigins:  "*",
	AllowHeaders:  strings.Join(headers, ","),
	AllowMethods:  "GET, POST, PUT, PATCH, DELETE",
	ExposeHeaders: "Content-Length",
	MaxAge:        int((2 * time.Hour).Seconds()),
}

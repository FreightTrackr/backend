package config

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

var origins = []string{
	"https://freighttrackr.github.io",
}

var headers = []string{
	"Origin",
	"Content-Type",
	"Accept",
	"Authorization",
	"Access-Control-Request-Headers",
	"Token",
	"Login",
	"Access-Control-Allow-Origin",
	"Bearer",
	"X-Requested-With",
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
	AllowOrigins:     "*",
	AllowHeaders:     strings.Join(headers, ","),
	AllowMethods:     "GET, POST, PUT, PATCH, DELETE",
	ExposeHeaders:    "Content-Length",
	AllowCredentials: true,
	MaxAge:           int((2 * time.Hour).Seconds()),
}

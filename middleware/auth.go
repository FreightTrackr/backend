package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/FreightTrackr/backend/models"
	"github.com/FreightTrackr/backend/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(models.Pesan{
				Status:  fiber.StatusUnauthorized,
				Message: "Authorization token required",
			})
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(models.Pesan{
				Status:  fiber.StatusUnauthorized,
				Message: "Invalid authorization format",
			})
		}

		tokenString = parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return utils.ReadPublicKeyFromEnv("PUBLIC_KEY")
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(models.Pesan{
				Status:  fiber.StatusUnauthorized,
				Message: "Invalid or expired token",
			})
		}
		c.Locals("user", token)
		return c.Next()
	}
}

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			utils.WriteJSONResponse(w, http.StatusUnauthorized, models.Pesan{
				Status:  http.StatusUnauthorized,
				Message: "Authorization token required",
			})
			return
		}
		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.WriteJSONResponse(w, http.StatusUnauthorized, models.Pesan{
				Status:  http.StatusUnauthorized,
				Message: "Invalid authorization format",
			})
			return
		}
		tokenString = parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return utils.ReadPublicKeyFromEnv("PUBLIC_KEY")
		})
		if err != nil || !token.Valid {
			utils.WriteJSONResponse(w, http.StatusUnauthorized, models.Pesan{
				Status:  http.StatusUnauthorized,
				Message: "Invalid or expired token",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

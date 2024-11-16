package config

import (
	"fmt"
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
			return utils.ReadPublicKeyFromFile("./keys/public.pem")
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

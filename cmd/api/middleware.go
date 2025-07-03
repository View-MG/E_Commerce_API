package main

import (
	"strings"

	"github.com/View-MG/be-project/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(secret []byte) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.UnauthorizedResponse(c, "Missing Authorization header")
		}

		//get token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return response.UnauthorizedResponse(c, "Invalid Authorization header format")
		}

		// Parse token
		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return secret, nil
		})

		if err != nil || !token.Valid {
			response.UnauthorizedResponse(c, "Invalid or expired token")
		}
		// Save claims to locals (optional)
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Locals("userClaims", claims)
		}

		return c.Next()
	}
}

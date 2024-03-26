package middleware

import (
	"github.com/gofiber/fiber/v2"
	"news-service/model"
)

func Authorize() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")

		if token == "" {
			return c.
				Status(fiber.StatusBadRequest).
				JSON(model.GeneralResponse{
					Code:    400,
					Message: "Bad Request",
					Data:    "Missing token",
				})
		}

		if token != "something_value" {
			return c.
				Status(fiber.StatusUnauthorized).
				JSON(model.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Invalid token",
				})
		}

		return c.Next()
	}
}

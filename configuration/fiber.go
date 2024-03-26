package configuration

import (
	"github.com/gofiber/fiber/v2"
	"news-service/exception"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"news-service/configuration"
	"news-service/controller"
	"news-service/exception"
	repository "news-service/repository/impl"
	service "news-service/service/impl"
)

func main() {
	config := configuration.New()
	database := configuration.NewDatabase(config)

	newsRepository := repository.NewNewsRepositoryImpl(database)

	newsService := service.NewNewsServiceImpl(&newsRepository)

	newsController := controller.NewNewsController(&newsService, config)

	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	newsController.Route(app)

	err := app.Listen(config.Get("SERVER_PORT"))
	exception.PanicLogging(err)
}

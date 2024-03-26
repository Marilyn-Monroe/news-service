package controller

import (
	"github.com/gofiber/fiber/v2"
	"news-service/configuration"
	"news-service/exception"
	"news-service/middleware"
	"news-service/model"
	"news-service/service"
	"strconv"
)

type NewsController struct {
	service.NewsService
	configuration.Config
}

func NewNewsController(newsService *service.NewsService, config configuration.Config) *NewsController {
	return &NewsController{NewsService: *newsService, Config: config}
}

func (controller NewsController) Route(app *fiber.App) {
	app.Post("/edit/:Id", middleware.Authorize(), controller.Update)
	app.Get("/list", middleware.Authorize(), controller.FindAll)
}

func (controller NewsController) Update(c *fiber.Ctx) error {
	var request model.NewsUpdateModel
	idString := c.Params("Id")
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		exception.PanicLogging(err)
	}

	response := controller.NewsService.Update(c.Context(), request, id)
	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}

func (controller NewsController) FindAll(c *fiber.Ctx) error {
	result := controller.NewsService.FindAll(c.Context())
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Success": true,
		"News":    result,
	})
}

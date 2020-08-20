package routes

import (
	handler "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/handlers"
	"github.com/gofiber/fiber"
)

//SetupRoutes inicializar as routas da aplicação
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/companies")
	api.Get("/", handler.GetAll)
	api.Get("/:id", handler.GetByID)
	api.Post("/new", handler.NewCompany)
	api.Post("/upload", handler.Upload)
	api.Delete("/:id", handler.Delete)
}

package routes

import (
	control "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/controllers"
	"github.com/gofiber/fiber"
)

//SetupRoutes inicializar as routas da aplicação
func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/companies")
	api.Get("/", control.GetAll)
	api.Get("/:id", control.GetByID)
	api.Post("/new", control.NewCompany)
	api.Post("/upload", control.Upload)
	api.Delete("/:id", control.Delete)
}

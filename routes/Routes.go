package routes

import (
	"github.com/gofiber/fiber"
)

//InitializeRoutes inicializar as routas da aplicação
func InitializeRoutes(app *fiber.App) {
	api := app.Group("/companies")

	//api.get("")

}

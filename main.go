package main

import (
	"log"
	// control "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/controllers"
	myDatabase "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/db"
	handler "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/handlers"
	route "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/routes"
	"github.com/gofiber/fiber"
)

func main() {
	err := myDatabase.DatabaseConnection("CompaniesDb")

	if err != nil {
		log.Fatalln("Error in Database Connection: ", err)
		return
	}

	handler.LoadData("./assets/", "q1_catalog.csv")

	app := fiber.New()
	app.Static("/", "./views/")
	route.SetupRoutes(app)

	app.Listen(3333)

}

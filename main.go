package main

import (
	"log"

	"github.com/gofiber/fiber"

	myDatabase "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/db"
)

func main() {
	err := myDatabase.DatabaseConnection("CompaniesDb")

	if err != nil {
		log.Fatalln("Error in Database Connection: ", err)
		return
	}

	conn := myDatabase.GetDB() //get database connection

	defer conn.Close()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})

	app.Listen(3000)

}

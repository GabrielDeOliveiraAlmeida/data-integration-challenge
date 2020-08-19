package main

import (
	"log"
	// control "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/controllers"
	myDatabase "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/db"
	handler "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/handlers"
)

func main() {
	err := myDatabase.DatabaseConnection("CompaniesDb")

	if err != nil {
		log.Fatalln("Error in Database Connection: ", err)
		return
	}

	handler.LoadData("q1_catalog.csv")
	// conn := myDatabase.GetDB() //get database connection

	// defer conn.Close()

	//app := fiber.New()

	// company := models.Company{
	// 	Name:    "tesju2",
	// 	Zipcode: "22245",
	// 	Website: "",
	// }

	// comp, isThere := control.Index(company)
	// if isThere == false {
	// 	control.Store(comp)
	// }

	//app.Listen(3000)

}

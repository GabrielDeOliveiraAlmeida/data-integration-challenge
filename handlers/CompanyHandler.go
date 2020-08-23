package handlers

import (
	"os"
	"strconv"

	control "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/controllers"
	model "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/models"
	"github.com/gofiber/fiber"
)

//GetAll - listar todas as companhias
func GetAll(c *fiber.Ctx) {
	companies := control.IndexAll()
	c.JSON(companies)
}

//GetCompany - pegar companhia que está nos parametros
func GetCompany(c *fiber.Ctx) {

	var company model.Company
	company.ID = 0
	company.CompanyName = c.Query("name")
	company.Zipcode = c.Query("zip")

	comp, err := control.Index(company, false)
	c.Set("Content-Type", "application/json")

	if err {
		c.Status(404)
	} else {
		// c.Status(200).JSON(map[string]interface{}{
		// 	"ID":          strconv.Itoa(comp.ID),
		// 	"CompanyName": comp.CompanyName,
		// 	"Zipcode":     comp.Zipcode,
		// 	"Website":     comp.Website,
		// })
		c.Status(200).JSON(comp)
	}
}

//GetByID - Retornar apenas uma compnhia do banco de dados
func GetByID(c *fiber.Ctx) {
	var company model.Company
	id := c.Params("id")
	company.ID, _ = strconv.Atoi(id)
	company.CompanyName = ""
	company.Zipcode = ""
	comp, err := control.Index(company, true)

	c.Set("Content-Type", "application/json")

	if err {
		c.Status(404)
	} else {
		if len(comp) > 0 {
			c.Status(200).JSON(map[string]interface{}{
				"ID":          strconv.Itoa(comp[0].ID),
				"CompanyName": comp[0].CompanyName,
				"Zipcode":     comp[0].Zipcode,
				"Website":     comp[0].Website,
			})
		}
	}

}

//Upload - Armazenar várias companhias
func Upload(c *fiber.Ctx) {
	file, err := c.FormFile("file")

	if err != nil {
		c.Status(400)
		return
	}
	tempName := "./assets/temp/"
	c.SaveFile(file, tempName+file.Filename)
	LoadData(tempName, file.Filename)

	os.Remove(tempName + file.Filename)
	GetAll(c)
}

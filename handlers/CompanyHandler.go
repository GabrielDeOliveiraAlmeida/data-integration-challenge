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

	comp, isFound := control.Index(company, false)

	c.Set("Content-Type", "application/json")

	if isFound {
		c.Status(404)
	} else {

		c.Status(200).JSON(map[string]interface{}{
			"ID":          strconv.Itoa(comp.ID),
			"CompanyName": comp.CompanyName,
			"Zipcode":     comp.Zipcode,
			"Website":     comp.Website,
		})
	}
}

//GetByID - Retornar apenas uma compnhia do banco de dados
func GetByID(c *fiber.Ctx) {
	var company model.Company
	id := c.Params("id")
	company.ID, _ = strconv.Atoi(id)
	company.CompanyName = ""
	company.Zipcode = ""
	comp, notFound := control.Index(company, true)

	c.Set("Content-Type", "application/json")

	if notFound {
		c.Status(404)
	} else {
		c.Status(200).JSON(map[string]interface{}{
			"ID":          strconv.Itoa(comp.ID),
			"CompanyName": comp.CompanyName,
			"Zipcode":     comp.Zipcode,
			"Website":     comp.Website,
		})
	}

}

//NewCompany - Armezenar uma nova companhia
func NewCompany(c *fiber.Ctx) {

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

//Delete uma companhia
func Delete(c *fiber.Ctx) {

}

package controllers

import (
	"strconv"

	model "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/models"
	"github.com/gofiber/fiber"
)

//GetAll - listar todas as companhias
func GetAll(c *fiber.Ctx) {
	companies := IndexAll()
	c.JSON(companies)
}

//GetByID - Retornar apenas uma compnhia do banco de dados
func GetByID(c *fiber.Ctx) {
	var company model.Company
	id := c.Params("id")
	company.ID, _ = strconv.Atoi(id)
	company.Name = ""
	company.Zipcode = ""
	comp, isFound := Index(company, true)

	c.Set("Content-Type", "application/json")

	if isFound {
		c.Status(200).JSON(map[string]interface{}{
			"ID":      strconv.Itoa(comp.ID),
			"Name":    comp.Name,
			"Zipcode": comp.Zipcode,
			"Website": comp.Website,
		})
	} else {
		c.Status(404)
	}

}

//NewCompany - Armezenar uma nova companhia
func NewCompany(c *fiber.Ctx) {

}

//Upload - Armazenar várias companhias
func Upload(c *fiber.Ctx) {

}

//Delete uma companhia
func Delete(c *fiber.Ctx) {

}

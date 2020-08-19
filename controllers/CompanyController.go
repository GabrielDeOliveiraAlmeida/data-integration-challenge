package controllers

import (
	"strings"

	myDatabase "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/db"
	model "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/models"
)

//Index busca alguma companhia no banco de dados
func Index(company model.Company) (model.Company, bool) {
	conn := myDatabase.GetDB()

	stmt := model.Company{
		Name:    strings.ToUpper(company.Name),
		Zipcode: company.Zipcode,
		Website: company.Website,
	}

	companyInDb := model.Company{}
	if conn.Where(stmt).First(&companyInDb).RecordNotFound() {
		return company, false
	}
	return companyInDb, true
}

//Store armazena a companhia no banco de dados
func Store(company model.Company) {
	conn := myDatabase.GetDB()

	stmt := model.Company{
		Name:    strings.ToUpper(company.Name),
		Zipcode: company.Zipcode,
		Website: company.Website,
	}
	conn.Create(&stmt)
}

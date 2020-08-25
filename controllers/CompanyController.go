package controllers

import (
	"strconv"
	"strings"

	myDatabase "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/db"
	model "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/models"
)

//AddChar Preencher a string de caracteres ch até o tamanho leng
func AddChar(str string, ch string, leng int) string {
	for i := len(str); i < leng; i++ {
		str = ch + str
	}
	return str
}

//IndexAll retorna todas as companhias
func IndexAll() []model.Company {
	conn := myDatabase.GetDB()

	var companiesInDb []model.Company
	conn.Find(&companiesInDb)
	return companiesInDb
}

//Index busca companhia pelo name e zipcode no banco de dados
func Index(company model.Company, fullName bool) []model.Company {
	conn := myDatabase.GetDB()

	var stmt string

	if company.ID != 0 {
		stmt = " id = ? "
	} else if company.CompanyName != "" {
		if fullName {
			stmt = "company_name = ? "
		} else {
			stmt = "company_name LIKE ?  "
			company.CompanyName = "%" + company.CompanyName + "%"
		}
		if company.Zipcode != "" {
			stmt = stmt + " AND zipcode = ? "
		}
	}
	var companyInDb []model.Company
	if company.ID != 0 { //Se possui ID, buscar pelo ID, apenas
		conn.Where(stmt, strconv.Itoa(company.ID)).Find(&companyInDb)
	} else if company.Zipcode == "" { //buscar pelo nome
		conn.Where(stmt, strings.ToUpper(company.CompanyName)).Find(&companyInDb)
	} else { // buscar pelo nome e zipcode
		conn.Where(stmt, strings.ToUpper(company.CompanyName), AddChar(company.Zipcode, "0", 5)).Find(&companyInDb)
	}
	if len(companyInDb) > 0 && companyInDb[0].ID != 0 { //foi encontrado
		return companyInDb
	}
	return nil
}

//Store armazena a companhia no banco de dados
func Store(company model.Company) {
	conn := myDatabase.GetDB()

	stmt := model.Company{
		CompanyName: strings.ToUpper(company.CompanyName),
		Zipcode:     AddChar(company.Zipcode, "0", 5),
		Website:     company.Website,
	}
	conn.Create(&stmt)
}

//Update todos os campos do tabela
func Update(company model.Company) {
	conn := myDatabase.GetDB()

	conn.Model(&company).Update("website", strings.ToLower(company.Website))
}

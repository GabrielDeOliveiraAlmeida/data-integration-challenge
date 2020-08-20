package controllers

import (
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
func Index(company model.Company, fullName bool) (model.Company, bool) {
	conn := myDatabase.GetDB()

	var stmt string
	if fullName {
		stmt = "name = ? "
	} else {
		stmt = "name like = ?  "
	}
	stmt = stmt + " AND zipcode = ? "

	var companyInDb model.Company
	notFound := false
	//buscar pelo nome e zipcode
	if conn.Where(stmt, strings.ToUpper(company.Name), AddChar(company.Zipcode, "0", 5)).First(&companyInDb).RecordNotFound() {
		notFound = true
	}

	if notFound { //se não foi encontrado
		return company, notFound
	}
	return companyInDb, notFound //foi encontrado
}

//Index busca companhia pelo name e zipcode no banco de dados
// func Index(company model.Company, fullName bool) (model.Company, bool) {
// 	conn := myDatabase.GetDB()

// 	var stmt string

// 	if company.ID != 0 {
// 		stmt = " id = ? "
// 	} else if company.Name != "" {
// 		if fullName {
// 			stmt = "name = ? "
// 		} else {
// 			stmt = "name like = ?  "
// 		}
// 		stmt = stmt + " AND zipcode = ? "
// 	}

// 	var companyInDb model.Company
// 	notFound := true
// 	if company.ID != 0 { //Se possui ID, buscar pelo ID, apenas
// 		if conn.Where(stmt, strconv.Itoa(company.ID)).First(&companyInDb).RecordNotFound() {
// 			notFound = false
// 		}
// 	} else { //buscar pelo nome e zipcode
// 		if conn.Where(stmt, strings.ToUpper(company.Name), company.Zipcode).First(&companyInDb).RecordNotFound() {
// 			notFound = false
// 		}
// 	}

// 	if notFound { //se não foi encontrado
// 		return companyInDb, notFound
// 	}
// 	return company, notFound //foi encontrado
// }

//Store armazena a companhia no banco de dados
func Store(company model.Company) {
	conn := myDatabase.GetDB()

	stmt := model.Company{
		Name:    strings.ToUpper(company.Name),
		Zipcode: AddChar(company.Zipcode, "0", 5),
		Website: company.Website,
	}
	conn.Create(&stmt)
}

//Update todos os campos do tabela
func Update(company model.Company) {
	conn := myDatabase.GetDB()

	conn.Model(&company).Update("website", strings.ToLower(company.Website))
}

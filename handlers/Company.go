package handlers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	control "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/controllers"
	models "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/models"
)

//LoadCompanyFromCSV deve carregar o ler o arquivo e carregar no banco de dados
func LoadCompanyFromCSV(file string) {
	fmt.Println("Abrir arquivo: " + file)
	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo ")
		log.Fatalln(err)
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	csvReader.Comma = ';'
	csvReader.FieldsPerRecord = -1

	csvLines, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	//Linha a linha do arquivo e enviar ao banco de dados
	for i, line := range csvLines {
		if i != 0 {

			name := line[0]
			zipcode := line[1]
			website := ""

			if len(line) > 2 {
				website = line[2]
			}
			//Modelo da tabela do banco
			company := models.Company{
				Name:    name,
				Zipcode: zipcode,
				Website: website,
			}

			//Checar se já existe, se não, salvaar no banco de dados
			comp, isThere := control.Index(company)
			if isThere == false {
				control.Store(comp)
			}

		}
	}
}

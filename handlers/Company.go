package handlers

import (
	"encoding/csv"
	"log"
	"os"
	models "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/models"
)

//Ler o arquivo CSV e carregar no banco de dados
func LoadCompanyFromCSV(file string) {
	csvFile, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}

	defer csvFile.Close()

	csvLines, err := csv.NewReader(csv).readAll()
	if err != nil {
		log.Fatalln(err)
	}

	//Linha a linha do arquivo e enviar ao banco de dados
	for i, line := range cvsLines {
		if i != 0 {

			name := line[0]
			zipcode := line[1]
			website := ""

			if len(line) > 2 {
				website = line[2]
			}
			//Modelo da tabela do banco
			company := models.Company{
				Name:	name,
				Zipcode: zipcode,
				Website: website,
			}

			//Checar se já existe se não salvaar no banco de dados
			if conn.

	}



}

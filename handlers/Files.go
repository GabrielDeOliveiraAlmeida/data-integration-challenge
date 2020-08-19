package handlers

//LoadData , ler o arquivo e carregar tudo ao banco de daodos
func LoadData(filename string) {
	file := "./assets/" + filename

	LoadCompanyFromCSV(file)

}

package handlers

//LoadData , ler o arquivo e carregar tudo ao banco de daodos
func LoadData(filepath string) {
	file := filepath + "q1_catalog"



	
	LoadCompanyFromCSV(file)

}

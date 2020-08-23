package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	//postgress
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

//URIDatabase serve para formatar a string para conexão com o banco de dados
func URIDatabase() string {

	// username := handler.GetVar("USER")
	// password := handler.GetVar("PASSWORD")
	// dbName := handler.GetVar("DBNAME")
	// dbPort := handler.GetVar("PORT")
	// dbHost := handler.GetVar("HOST")
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")
	dbPort := os.Getenv("PORT")
	dbHost := os.Getenv("HOST")

	url := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", dbHost, username, dbName, dbPort, password)

	return url
}

//DatabaseConnection realiza a conexão com o banco de dados
func DatabaseConnection(database string) error {
	url := URIDatabase()
	conn, error := gorm.Open("postgres", url)

	if error != nil {
		return error
	}

	db = conn

	return nil
}

//GetDB returna o objeto de base de dados
func GetDB() *gorm.DB {
	return db
}

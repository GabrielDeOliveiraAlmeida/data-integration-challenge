package handlers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//GetVar carrega variaveis de desenvolvimento
func GetVar(variable string) string {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	return os.Getenv(variable)
}

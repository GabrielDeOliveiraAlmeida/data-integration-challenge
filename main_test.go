package main

import (
	"io"
	"net/http"
	"testing"

	myDatabase "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/db"
	handler "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/handlers"
	route "github.com/GabrielDeOliveiraAlmeida/data-integration-challenge/routes"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/assert"
)

var conn *gorm.DB

func TestDatabase(t *testing.T) {
	err := myDatabase.DatabaseConnection("CompaniesDb")
	assert.NoError(t, err)
	conn = myDatabase.GetDB()
}

func TestFile(t *testing.T) {
	err := handler.LoadData("./assets/", "q1_catalog.csv")
	assert.NoError(t, err)
}

func TestGetCompanies(t *testing.T) {
	app := fiber.New()
	route.SetupRoutes(app)

	apitest.New().
		HandlerFunc(FiberToHandlerFunc(app)).
		Get("/api/companies/1").
		Expect(t).
		Body(`{"ID": 1, CompanyName: }`)
	Status(http.StatusOK).
		End()

}

func FiberToHandlerFunc(app *fiber.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := app.Test(r)
		if err != nil {
			panic(err)
		}

		// copy headers
		for k, vv := range resp.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			panic(err)
		}
	}
}

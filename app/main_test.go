package main

import (
	"errors"
	"log"
	"os"
	"testing"
	"time"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/db"
	"github.com/jarcoal/httpmock"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	godotenv.Load("./env/.env.test")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")
	db := db.Init()
	db.Exec("TRUNCATE TABLE outpatient_histories")
	db.Exec("TRUNCATE TABLE tmp_outpatient_histories")
	os.Exit(exitVal)
}

func TestHandler(t *testing.T) {
	t.Run("http request success", func(t *testing.T) {
		time := time.Now().AddDate(0, 0, -1).Format("20060102")
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		mock_json, _ := os.ReadFile("../test_data/20221230.json")
		httpmock.RegisterResponder("GET", "https://opendata.corona.go.jp/api/covid19DailySurvey/"+time,
			httpmock.NewStringResponder(200, string(mock_json)),
		)
		err := CreateOutpatientHistoires()
		if err != nil {
			t.Fatal("Error failed to trigger with an invalid request")
		}
	})

	t.Run("http request failer", func(t *testing.T) {
		time := time.Now().AddDate(0, 0, -1).Format("20060102")
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		response_error := errors.New("errror")
		httpmock.RegisterResponder("GET", "https://opendata.corona.go.jp/api/covid19DailySurvey/"+time,
			httpmock.NewErrorResponder(response_error),
		)
		err := CreateOutpatientHistoires()
		if err == nil {
			t.Fatal("Error failed to trigger with an invalid request")
		}
	})
}

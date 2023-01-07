package main

import (
	"net/http"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"github.com/aws/aws-lambda-go/lambda"
)

func CreateOutpatientHistoires() error {
	go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=:rocket:start_batch_for_create_last_day_city_patient")
	err := service.CreateCityOutPatients()
	if err != nil {
		go http.Get("http://127.0.0.1:3000/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
	}
	go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=:rocket:finish_batch_for_create_last_day_city_patient")
	return nil
}

func main() {
	lambda.Start(CreateOutpatientHistoires)
}

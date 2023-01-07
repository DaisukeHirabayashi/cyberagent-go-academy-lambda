package main

import (
	"net/http"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"github.com/aws/aws-lambda-go/lambda"
)

func CreateOutpatientHistoires() error {
	go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=:rocket:start_batch_for_create_last_day_outpatient_hisotory")
	err := service.CreateCityOutPatients()
	if err != nil {
		go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
	}
	go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=:tada:finish_batch_for_create_last_day_outpatient_hisotory")
	return nil
}

func main() {
	lambda.Start(CreateOutpatientHistoires)
}

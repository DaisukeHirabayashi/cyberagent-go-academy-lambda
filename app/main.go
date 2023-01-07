package main

import (
	"net/http"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"github.com/aws/aws-lambda-go/lambda"
)

func CreateOutpatientHistoires() error {
	go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + ":rocket:start_batch_for_create_last_day_city_patient")

	outpatientHistories, err := service.GetLastDayOutpatientHistory()
	if err != nil {
		go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
	}

	err = service.CreateOutpatientHistoires(outpatientHistories)

	if err != nil {
		go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
	}

	go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + ":tada: finish start batch for create last day patient history")
	return nil
}

func main() {
	lambda.Start(CreateOutpatientHistoires)
}

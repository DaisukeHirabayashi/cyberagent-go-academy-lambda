package main

import (
	"net/http"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/client"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"github.com/aws/aws-lambda-go/lambda"
)

func CreateOutpatientHistoires() error {
	go http.Get("https://wgmjrdremf.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + ":rocket:start_batch_for_create_last_day_outpatient_hisotory")

	service := service.HospitalService{Client: client.CoronaClient{}}
	outpatientHistories, err := service.GetLastDayOutpatientHistory()
	if err != nil {
		go http.Get("https://wgmjrdremf.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
		return err
	}

	err = service.CreateOutpatientHistoires(outpatientHistories)

	if err != nil {
		go http.Get("https://wgmjrdremf.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
		return err
	}

	go http.Get("https://wgmjrdremf.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + ":tada:finish_batch_for_create_last_day_outpatient_hisotory")
	return nil
}

func main() {
	lambda.Start(CreateOutpatientHistoires)
}

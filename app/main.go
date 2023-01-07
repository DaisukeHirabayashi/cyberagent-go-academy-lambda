package main

import (
	"net/http"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"github.com/aws/aws-lambda-go/lambda"
)

func CreateOutpatientHistoires() error {
	outpatientHistories, err := service.GetLastDayOutpatientHistory()
	if err != nil {
		go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
	}

	err = service.CreateOutpatientHistoires(outpatientHistories)

	if err != nil {
		go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
	}
	return nil
}

func main() {
	lambda.Start(CreateOutpatientHistoires)
}

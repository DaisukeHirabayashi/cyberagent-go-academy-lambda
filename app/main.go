package main

import (
	"net/http"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"github.com/aws/aws-lambda-go/lambda"
)

func CreateOutpatientHistoires() error {
	go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + "昨日のデータ取得開始")

	outpatientHistories, err := service.GetLastDayOutpatientHistory()
	if err != nil {
		go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
	}

	err = service.CreateOutpatientHistoires(outpatientHistories)

	if err != nil {
		go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
	}

	go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + "昨日のデータ取得終了")
	return nil
}

func main() {
	lambda.Start(CreateOutpatientHistoires)
}

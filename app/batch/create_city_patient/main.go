package main

import (
	"net/http"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"github.com/aws/aws-lambda-go/lambda"
)

func CreateOutpatientHistoires() error {
	go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + "患者数のデータ整形開始")
	err := service.CreateCityOutPatients()
	if err != nil {
		go http.Get("http://127.0.0.1:3000/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
	}
	go http.Get("https://a5pky2m9q6.execute-api.ap-northeast-1.amazonaws.com/Prod/notification?message=" + "患者数のデータ整形終了")
	return nil
}

func main() {
	lambda.Start(CreateOutpatientHistoires)
}

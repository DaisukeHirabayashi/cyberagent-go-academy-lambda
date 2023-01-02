package main

import (
	"net/http"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"github.com/aws/aws-lambda-go/lambda"
)

func CreateOutpatientHistoires() error {
	err := service.CreateCityOutPatients()
	if err != nil {
		go http.Get("http://127.0.0.1:3000/notification?message=" + "error GetLastDayOutpatientHistory:" + err.Error())
	}
	return nil
}

func main() {
	lambda.Start(CreateOutpatientHistoires)
}

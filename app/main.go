package main

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func OnlyErrors() error {
	// service.GetBooks()
	outpatientHistories, err := service.GetLastDayOutpatientHistory()
	if err != nil {
		http.PostForm("https://hooks.slack.com/services/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", url.Values{"payload": {string(err.Error())}})
	}

	err = service.CreateOutpatientHistoires(outpatientHistories)

	if err != nil {
		http.PostForm("https://hooks.slack.com/services/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", url.Values{"payload": {string(err.Error())}})
	}
	return nil
}

func main() {
	lambda.Start(OnlyErrors)
}

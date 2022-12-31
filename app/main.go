package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"github.com/aws/aws-lambda-go/lambda"
)

type payload struct {
	Text string `json:"text"`
}

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

func OnlyErrors() error {
	//service.GetBooks()
	outpatientHistories, err := service.GetLastDayOutpatientHistory()
	if err != nil {
		err_json, _ := json.Marshal(payload{Text: err.Error()})
		http.PostForm("https://hooks.slack.com/services/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", url.Values{"payload": {string(err_json)}})
	}

	err = service.CreateOutpatientHistoires(outpatientHistories)

	if err != nil {
		err_json, _ := json.Marshal(payload{Text: err.Error()})
		http.PostForm("https://hooks.slack.com/services/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", url.Values{"payload": {string(err_json)}})
	}
	return nil
}

func main() {
	lambda.Start(OnlyErrors)
}

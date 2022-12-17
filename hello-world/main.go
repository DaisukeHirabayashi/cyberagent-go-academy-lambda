package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

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
	p, err := json.Marshal(payload{Text: time.Now().Format("2006-01-02 15:04:05") + "08: Hira"})
	if err != nil {
		return err
	}
	resp, err := http.PostForm("https://hooks.slack.com/services/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", url.Values{"payload": {string(p)}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func main() {
	lambda.Start(OnlyErrors)
}

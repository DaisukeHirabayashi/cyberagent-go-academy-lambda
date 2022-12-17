package main

import (
	"errors"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/db"
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
	db.Init()
	defer db.Close()
	return nil
}

func main() {
	lambda.Start(OnlyErrors)
}

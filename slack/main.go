package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
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

type payload struct {
	Text string `json:"text"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	p, err := json.Marshal(payload{Text: request.PathParameters["message"]})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	resp, err := http.PostForm("https://hooks.slack.com/services/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", url.Values{"payload": {string(p)}})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	defer resp.Body.Close()
	return events.APIGatewayProxyResponse{
		Body:       "success",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

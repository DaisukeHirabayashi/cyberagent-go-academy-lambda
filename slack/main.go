package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type payload struct {
	Text string `json:"text"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	p, err := json.Marshal(payload{Text: request.MultiValueQueryStringParameters["message"][0]})

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

package main

import (
	"encoding/json"
	"fmt"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/service"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	city_outpatients, err := service.GetCityPatients()

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	res, err := json.Marshal(city_outpatients)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("%s", res),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

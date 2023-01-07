package service

import (
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/api/repository"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/db"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/entity"
	"github.com/aws/aws-lambda-go/events"
)

func GetCityPatients(request events.APIGatewayProxyRequest) ([]entity.City_Outpatient, error) {
	db := db.Init()
	pref_name := request.QueryStringParameters["pref_name"]
	city_name := request.QueryStringParameters["city_name"]

	var city_outpatients []entity.City_Outpatient
	if err := db.Scopes(repository.SearchByPrefName(pref_name)).Scopes(repository.SearchByCityName(city_name)).Find(&city_outpatients).Error; err != nil {
		return city_outpatients, err
	}

	return city_outpatients, nil
}

package service

import (
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/db"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/entity"
)

func GetCityPatients() ([]entity.City_Outpatient, error) {
	db := db.Init()

	var city_outpatients []entity.City_Outpatient
	if err := db.Find(&city_outpatients).Error; err != nil {
		return city_outpatients, err
	}

	return city_outpatients, nil
}

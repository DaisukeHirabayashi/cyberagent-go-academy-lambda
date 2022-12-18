package service

import (
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/client"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/db"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/entity"
)

// 前日の履歴を取ってくる
func GetLastDayOutpatientHistory() error {
	db := db.GetDB()
	var prefectures []entity.Prefecture

	if err := db.Find(&prefectures).Error; err != nil {
		return err
	}

	for _, prefecture := range prefectures {
		pref := prefecture.Name
		client.GetMedicalSystem(&pref, nil)
	}

	return nil
}

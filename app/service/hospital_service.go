package service

import (
	"encoding/json"
	"log"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/client"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/dao"
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
		reponse_body, _ := client.GetMedicalSystem(&pref, nil)
		var hospitals []dao.Hospital
		json.Unmarshal(reponse_body, &hospitals)
		log.Println("Error Request:", hospitals[0])
	}

	return nil
}

package service

import (
	"encoding/json"
	"log"
	"time"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/client"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/dao"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/db"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/mapper"
)

// 前日の履歴を取ってくる
func GetLastDayOutpatientHistory() []dao.Hospital {
	time := time.Now().Format("20060102")
	reponse_body, _ := client.GetMedicalSystem(&time)
	var hospitals []dao.Hospital
	json.Unmarshal(reponse_body, &hospitals)
	return hospitals
}

func CreateOutpatientHistoires(dao_hospitals []dao.Hospital) error {
	db := db.GetDB()
	outpatinet_histories, err := mapper.HospitalDaosToOutpatientHistorys(dao_hospitals)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	for i, outpatinet_history := range outpatinet_histories {
		if err := db.Create(&outpatinet_history).Error; err != nil {
			log.Println("Error:", i, "番目:", outpatinet_history)
			return err
		}
	}

	return nil
}

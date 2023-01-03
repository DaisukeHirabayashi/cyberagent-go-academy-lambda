package service

import (
	"encoding/json"
	"log"
	"time"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/client"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/dao"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/db"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/entity"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/mapper"
)

// 前日の履歴を取ってくる
func GetLastDayOutpatientHistory() ([]dao.Hospital, error) {
	time := time.Now().AddDate(0, 0, -1).Format("20060102")
	reponse_body, err := client.GetMedicalSystem(&time)
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}

	var hospitals []dao.Hospital
	json.Unmarshal(reponse_body, &hospitals)
	return hospitals, nil
}

func CreateOutpatientHistoires(dao_hospitals []dao.Hospital) error {
	db := db.Init()
	outpatinet_histories, err := mapper.HospitalDaosToOutpatientHistorys(dao_hospitals)
	tmp_outpatinet_histories := mapper.OutpatientHistoryEntityToTmpOutPatientHistory(outpatinet_histories)

	if err != nil {
		log.Println("Error:", err)
		return err
	}

	if err := db.CreateInBatches(&tmp_outpatinet_histories, 1000).Error; err != nil {
		log.Println("Error:", err)
		return err
	}

	if err := db.CreateInBatches(&outpatinet_histories, 1000).Error; err != nil {
		log.Println("Error:", err)
		return err
	}
	return nil
}

func CreateCityOutPatients() error {
	db := db.Init()
	var city_outpatients []entity.City_Outpatient

	err := db.Table("outpatient_histories").Select("pref_name, city_name, submit_date, count(*) as number").Group("pref_name, city_name, submit_date").Scan(&city_outpatients).Error
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	if err := db.CreateInBatches(&city_outpatients, 100).Error; err != nil {
		log.Println("Error:", err)
		return err
	}
	log.Println("Create Success:", err)
	return nil
}

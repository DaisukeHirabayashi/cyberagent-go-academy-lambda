package service

import (
	"encoding/json"
	"log"
	"time"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/client"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/dao"
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
	outpatinet_histories, err := mapper.HospitalDaosToOutpatientHistorys(dao_hospitals)
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	log.Println("outpatinethistory:", outpatinet_histories[0])

	return nil
}

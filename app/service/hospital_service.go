package service

import (
	"encoding/json"
	"log"
	"time"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/client"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/dao"
)

// 前日の履歴を取ってくる
func GetLastDayOutpatientHistory() error {
	time := time.Now().Format("20060102")
	reponse_body, _ := client.GetMedicalSystem(&time)
	var hospitals []dao.Hospital
	json.Unmarshal(reponse_body, &hospitals)
	log.Println("hospital:", hospitals[0])

	return nil
}

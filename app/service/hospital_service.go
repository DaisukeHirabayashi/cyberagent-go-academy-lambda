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

	prefecture := "石川県"
	client.GetMedicalSystem(&prefecture, nil)

	return nil
}

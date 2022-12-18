package service

import (
	"log"

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

	log.Print(prefectures)

	return nil
}

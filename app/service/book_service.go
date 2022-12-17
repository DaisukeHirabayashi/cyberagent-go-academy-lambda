package service

import (
	"log"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/db"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/entity"
)

func GetBooks() error {
	db := db.GetDB()
	var books []entity.Book

	if err := db.Find(&books).Error; err != nil {
		return err
	}

	log.Print(books)

	return nil
}

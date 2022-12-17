package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	log.Print("DB init...")
	databaseUrl := "academy08:78fqoLQpJ9XXpAab8fVuQmbzdmhACZ@(ca-academy-db.c9ml7do7yvmn.ap-northeast-1.rds.amazonaws.com:3306)/academy08"
	db, err = gorm.Open("mysql", databaseUrl)
	if err != nil {
		panic(err)
	}
	log.Print("DB connection success!")
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

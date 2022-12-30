package db

import (
	"log"

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	log.Print("DB init...")
	databaseUrl := "academy08:78fqoLQpJ9XXpAab8fVuQmbzdmhACZ@(ca-academy-db.c9ml7do7yvmn.ap-northeast-1.rds.amazonaws.com:3306)/academy08"
	sqlDB, err := sql.Open("mysql", databaseUrl)

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	log.Print("DB connection success!")
	return db
}

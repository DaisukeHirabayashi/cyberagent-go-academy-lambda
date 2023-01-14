package db

import (
	"log"
	"os"

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	databaseUrl := os.Getenv("DATABASEURL")
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

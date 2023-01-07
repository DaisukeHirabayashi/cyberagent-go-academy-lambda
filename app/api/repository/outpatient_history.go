package repository

import (
	"gorm.io/gorm"
)

func SearchByPrefName(pref_name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pref_name != "" {
			db.Where("pref_name=?", pref_name)
		}
		return db
	}
}

func SearchByCityName(city_name string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if city_name != "" {
			db.Where("city_name=?", city_name)
		}
		return db
	}
}

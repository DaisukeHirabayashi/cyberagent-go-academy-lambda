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
	"gorm.io/gorm"
)

type HospitalServiceInterface interface {
	GetLastDayOutpatientHistory() ([]dao.Hospital, error)
	CreateOutpatientHistoires(dao_hospitals []dao.Hospital) error
}

type HospitalService struct {
	Client client.CoronaClientInterface
}

// 前日の履歴を取ってくる
func (s *HospitalService) GetLastDayOutpatientHistory() ([]dao.Hospital, error) {
	time := time.Now().AddDate(0, 0, -1).Format("20060102")
	reponse_body, err := s.Client.GetMedicalSystem(&time)
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}

	var hospitals []dao.Hospital
	json.Unmarshal(reponse_body, &hospitals)
	return hospitals, nil
}

func (s *HospitalService) CreateOutpatientHistoires(dao_hospitals []dao.Hospital) error {
	db := db.Init()
	outpatinet_histories, err := mapper.HospitalDaosToOutpatientHistorys(dao_hospitals)
	tmp_outpatinet_histories := mapper.OutpatientHistoryEntityToTmpOutPatientHistory(outpatinet_histories)

	if err != nil {
		log.Println("Error:", err)
		return err
	}

	tr_err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(&tmp_outpatinet_histories, 1000).Error; err != nil {
			log.Println("Error:", err)
			return err
		}

		if err := db.CreateInBatches(&outpatinet_histories, 1000).Error; err != nil {
			log.Println("Error:", err)
			return err
		}
		return nil
	})

	if tr_err != nil {
		log.Println("Error:", err)
		return tr_err
	}
	return nil
}

func CreateCityOutPatients() error {
	db := db.Init()
	var city_outpatients []entity.City_Outpatient

	err := db.Table("tmp_outpatient_histories").Select("pref_name, city_name, submit_date, count(*) as number").Group("pref_name, city_name, submit_date").Scan(&city_outpatients).Error
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	tr_err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(&city_outpatients, 100).Error; err != nil {
			log.Println("Error:", err)
			return err
		}
		tx.Exec("TRUNCATE TABLE tmp_outpatient_histories")

		return nil
	})

	if tr_err != nil {
		log.Println("Error:", err)
		return tr_err
	}

	log.Println("Create Success")
	return nil
}

package mapper

import (
	"log"
	"time"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/dao"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/entity"
)

func HospitalDaosToOutpatientHistorys(dao_hospitals []dao.Hospital) ([]entity.OutpatientHistory, error) {
	var outpatient_histries []entity.OutpatientHistory
	for _, dao_hospital := range dao_hospitals {
		submit_date, err := time.Parse("2006-01-02", dao_hospital.SubmitDate)

		if err != nil {
			log.Println("Error:", err, dao_hospital)
			return nil, err
		}
		entity_hospital := entity.OutpatientHistory{FacilityId: dao_hospital.FacilityId, FacilityName: dao_hospital.FacilityName, ZipCode: dao_hospital.ZipCode, PrefName: dao_hospital.PrefName, FacilityAddr: dao_hospital.FacilityAddr, FacilityTel: dao_hospital.FacilityTel, SubmitDate: submit_date, FacilityType: dao_hospital.FacilityType, AnsType: dao_hospital.AnsType, LocalGovCode: dao_hospital.LocalGovCode, CityName: dao_hospital.CityName, FaciltyCode: dao_hospital.FaciltyCode}
		outpatient_histries = append(outpatient_histries, entity_hospital)
	}

	return outpatient_histries, nil
}

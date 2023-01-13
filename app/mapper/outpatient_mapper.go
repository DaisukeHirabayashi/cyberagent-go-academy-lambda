package mapper

import (
	"log"
	"time"

	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/dao"
	"github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/entity"
	pb "github.com/DaisukeHirabayashi/cyberagent-go-academy-lambda/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func HospitalDaosToOutpatientHistorys(dao_hospitals []dao.Hospital) ([]entity.OutpatientHistory, error) {
	var outpatient_histries []entity.OutpatientHistory
	for _, dao_hospital := range dao_hospitals {
		submit_date, err := time.Parse("2006-01-02", dao_hospital.SubmitDate)

		if err != nil {
			log.Println("Error:", err, dao_hospital)
			return nil, err
		}
		entity_hospital := entity.OutpatientHistory{FacilityId: dao_hospital.FacilityId, FacilityName: dao_hospital.FacilityName, ZipCod: dao_hospital.ZipCode, PrefName: dao_hospital.PrefName, FacilityAddr: dao_hospital.FacilityAddr, FacilityTel: dao_hospital.FacilityTel, SubmitDate: submit_date, FacilityType: dao_hospital.FacilityType, AnsType: dao_hospital.AnsType, LocalGovCode: dao_hospital.LocalGovCode, CityName: dao_hospital.CityName, FacilityCode: dao_hospital.FaciltyCode}
		outpatient_histries = append(outpatient_histries, entity_hospital)
	}

	return outpatient_histries, nil
}

func OutpatientHistoryEntityToTmpOutPatientHistory(outpatient_histories []entity.OutpatientHistory) []entity.TmpOutpatientHistory {
	var tmp_outpatient_histories []entity.TmpOutpatientHistory
	for _, outpatient_history := range outpatient_histories {
		tmp_outpatient_history := entity.TmpOutpatientHistory(outpatient_history)
		tmp_outpatient_histories = append(tmp_outpatient_histories, tmp_outpatient_history)
	}
	return tmp_outpatient_histories
}

func OutpatientHistoryEntityToPbOutpatientHistory(outpatient_history entity.OutpatientHistory) *pb.OutpatientHistory {
	return &pb.OutpatientHistory{
		Id:           uint64(outpatient_history.Id),
		FacilityId:   outpatient_history.FacilityId,
		ZipCod:       outpatient_history.ZipCod,
		PrefName:     outpatient_history.PrefName,
		FacilityAddr: outpatient_history.FacilityAddr,
		FacilityTel:  outpatient_history.FacilityTel,
		SubmitDate:   timestamppb.New(outpatient_history.SubmitDate),
		FacilityType: outpatient_history.FacilityType,
		AnsType:      outpatient_history.AnsType,
		LocalGovCode: outpatient_history.LocalGovCode,
		CityName:     outpatient_history.CityName,
		FacilityCode: outpatient_history.FacilityCode,
		CreatedAt:    timestamppb.New(outpatient_history.CreatedAt),
		UpdatedAt:    timestamppb.New(outpatient_history.UpdatedAt),
	}
}

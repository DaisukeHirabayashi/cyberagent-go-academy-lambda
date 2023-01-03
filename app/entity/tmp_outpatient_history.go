package entity

import "time"

type TmpOutpatientHistory struct {
	Id           uint      `json:"id"`
	FacilityId   string    `json:"facility_id"`
	FacilityName string    `json:"facility_name"`
	ZipCod       string    `json:"zip_cod"`
	PrefName     string    `json:"pref_name"`
	FacilityAddr string    `json:"facility_addr"`
	FacilityTel  string    `json:"facility_tel"`
	SubmitDate   time.Time `json:"submit_date"`
	FacilityType string    `json:"facility_type"`
	AnsType      string    `json:"ans_type"`
	LocalGovCode string    `json:"local_gov_code"`
	CityName     string    `json:"city_name"`
	FacilityCode string    `json:"facility_code"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

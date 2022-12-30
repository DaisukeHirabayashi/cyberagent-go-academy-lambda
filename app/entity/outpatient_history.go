package entity

import "time"

type OutpatientHistory struct {
	Id           uint      `json:"id"`
	FacilityId   string    `json:"facility_id"`
	FacilityName string    `json:"facility_name"`
	ZipCode      string    `json:"zip_code"`
	PrefName     string    `json:"pref_name"`
	FacilityAddr string    `json:"facility_addr"`
	FacilityTel  string    `json:"facility_tel"`
	SubmitDate   time.Time `json:"submit_date"`
	FacilityType string    `json:"facility_type"`
	AnsType      string    `json:"ans_type"`
	LocalGovCode string    `json:"local_gov_code"`
	CityName     string    `json:"city_name"`
	FaciltyCode  string    `json:"facility_code"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

package entity

import "time"

type City_Outpatient struct {
	Id         uint      `json:"id"`
	PrefName   string    `json:"pref_name"`
	CityName   string    `json:"city_name"`
	Number     uint      `json:"number"`
	SubmitDate time.Time `json:"submit_date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

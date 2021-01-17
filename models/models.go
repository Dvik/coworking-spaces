package models

type CoworkingSpace struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Speed    string `json:"speed"`
	Charging string `json:"charging"`
}

package models

type Player struct {
	Nome string  `json:"nome" bson:"nome"`
	Ppg  float64 `json:"ppg" bson:"ppg"`
}

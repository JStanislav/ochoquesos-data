package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Goals     int    `json:"goals"`
}

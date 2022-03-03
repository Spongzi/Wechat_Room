package models

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	UUID     int64  `json:"UUID" gorm:"type:int(10);"`
	Name     string `json:"Name" gorm:"type:varchar(30)"`
	Sing     string `json:"Sing" gorm:"type:varchar(30)"`
	Photo    string `json:"Photo" gorm:"type:varchar(30)"`
	Address  string `json:"Address" gorm:"type:varchar(30)"`
	Tel      int    `json:"Tel" gorm:"type:int(20);"`
	Password string `json:"Password" gorm:"type:varchar(30)"`
	Email    string `json:"Email" gorm:"type:varchar(30)"`
	Status   string `json:"Status" gorm:"type:varchar(30)"`
}

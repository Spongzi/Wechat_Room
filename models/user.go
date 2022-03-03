package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UUID     int64    `json:"UUID" gorm:"type:int(10);"`
	LoginId  string   `json:"LoginId" gorm:"type:varchar(30);"`
	Name     string   `json:"Name" gorm:"type:varchar(30);"`
	Sing     string   `json:"Sing" gorm:"type:varchar(30);"`
	Photo    string   `json:"Photo" gorm:"type:varchar(30);"`
	Address  string   `json:"Address" gorm:"type:varchar(30);"`
	Tel      int      `json:"Tel" gorm:"type:int(20);"`
	Password string   `json:"Password" gorm:"type:varchar(30);"`
	Account  string   `json:"Account" gorm:"type:varchar(30);"`
	Email    string   `json:"Email" gorm:"type:varchar(30);"`
	Status   string   `json:"Status" gorm:"type:varchar(30);"`
	Friends  []Friend `json:"friends" gorm:"foreignKey:ID"`
}

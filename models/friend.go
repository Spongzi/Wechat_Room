package models

import "gorm.io/gorm"

type Friend struct {
	gorm.Model
	UserId   uint
	FriendId uint
}

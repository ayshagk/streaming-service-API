package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	UserID uint
	MovieID uint
	Rating int
}
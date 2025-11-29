package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title string 
	Genre string 
	Description string
	ReleaseYear int
	ImageURL string
}
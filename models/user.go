package models

import "gorm.io/gorm"

type User struct { 
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type RegisterUser struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Role     string `json:"role"`
}
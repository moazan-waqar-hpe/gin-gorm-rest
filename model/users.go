package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

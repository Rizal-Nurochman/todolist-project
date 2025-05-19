package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string		`json:"email" binding:"required" gorm:"unique"`
	Password string	`json:"password" binding:"required"`
	Todos []Todos		`gorm:"foreignKey:userid"`
}
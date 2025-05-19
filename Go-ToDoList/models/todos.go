package models

import "gorm.io/gorm"

type Todos struct {
	gorm.Model
	UserID uint			`json:"userid"`
	Name string			`json:"name" binding:"required"`
	IsDone bool			`json:"isdone" binding:"required"`
}
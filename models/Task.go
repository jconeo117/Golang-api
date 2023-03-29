package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Title       string `gorm:"not null;unique" json:"title"`
	Description string `json:"description"`
	Done        bool   `gorm:"default:false"`
	UserId      uint   `json:"user_id"`
}

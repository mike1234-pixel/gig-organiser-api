package models

import (
	"gorm.io/gorm"
)

type Job struct {
	gorm.Model
	Title        string `json:"title"`
	Organisation string `json:"organisation"`
	Description  string `json:"description"`
	Priority     int    `json:"priority"`
	Status       string `json:"status"`
}

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	Password string `gorm:"type:varchar(255)"`
}

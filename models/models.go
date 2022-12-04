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

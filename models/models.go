package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email"`
	Password string `gorm:"type:varchar(255)" json:"password"`
	Jobs     []Job  `gorm:"ForeignKey:UserID" json:"jobs"`
}

type Job struct {
	gorm.Model
	UserID       uint     `gorm:"primary_key" json:"userid"`
	Title        string   `json:"title"`
	Organisation string   `json:"organisation"`
	Description  string   `json:"description"`
	Priority     int      `json:"priority"`
	Status       string   `json:"status"`
	Actions      []Action `gorm:"ForeignKey:JobID" json:"actions"`
}

type Action struct {
	gorm.Model
	UserID      uint   `gorm:"primary_key" json:"userid"`
	JobID       uint   `json:"jobid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CompleteBy  string `json:"complete_by"`
	Completed   bool   `json:"completed"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	ID          uint       `json:"id"`
	TaskNama    string     `json:"tasknama"`
	Completed   string     `json:"completed"`
	IsiListData []ListData `gorm:"foreignKey:ListData"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// IdList uint `gorm :"foreignKey:"id" json: Idlist`
	// Created   time.Time `json:"created at"`
}

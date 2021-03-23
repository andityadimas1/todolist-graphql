package models

import (
	"time"

	"gorm.io/gorm"
)

type ListData struct {
	gorm.Model
	ID        uint   `json:"id"`
	ListNama  string `json:"listnama"`
	Completed string `json:"completed"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	// IdList uint `gorm :"foreignKey:"id" json: Idlist`
	// Created   time.Time `json:"created at"`
}

var IsiListData = []ListData{
	{
		ID:        1,
		ListNama:  "Makan Malam",
		Completed: "True",
	},
	{
		ID:        2,
		ListNama:  "Minum Kopi",
		Completed: "False",
	},
	{
		ID:        3,
		ListNama:  "Ronda",
		Completed: "True",
	},
}

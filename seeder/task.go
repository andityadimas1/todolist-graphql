package seeder

import (
	"fmt"
	"todolist-graphql/models"

	"gorm.io/gorm"
)

func TaskSeeder(db *gorm.DB) {
	var AddtaskArray = [...][2]string{
		{"Makan Pagi", "true"},
		{"Mandi pagi", "true"},
		{"Pergi ke sekolah", "true"},
		{"Napping", "false"},
		{"main kelereng", "false"},
		{"Nonton netflix", "false"},
		{"main dota2 sampai pagi", "true"},
	}

	var task models.Task

	for _, data := range AddtaskArray {
		// Get Data from Array
		task.ID = 0
		task.TaskNama = data[0]
		task.Completed = data[1]
		db.Create(&task)
	}
	fmt.Println("seed!")
}

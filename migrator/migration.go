package migrator

import (
	"fmt"
	"todolist-graphql/models"

	"gorm.io/gorm"
)

//function untuk migrasi ke database
func Migrations(db *gorm.DB) {
	db.Migrator().DropTable(&models.User{})
	if check := db.Migrator().HasTable(&models.User{}); !check { // kalau belum ada di db postgre
		db.Migrator().CreateTable(&models.User{})
		fmt.Println("Table berhasil tercreate")
	}
	db.Migrator().DropTable(&models.Task{})
	if check := db.Migrator().HasTable(&models.Task{}); !check { // kalau belum ada di db postgre
		db.Migrator().CreateTable(&models.Task{})
		fmt.Println("Table berhasil tercreate")
	}
	db.Migrator().DropTable(&models.ListData{})
	if check := db.Migrator().HasTable(&models.ListData{}); !check { // kalau belum ada di db postgre
		db.Migrator().CreateTable(&models.ListData{})
	}
}

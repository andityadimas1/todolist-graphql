package middleware

import "gorm.io/gorm"

type StrDB struct {
	DB *gorm.DB
}

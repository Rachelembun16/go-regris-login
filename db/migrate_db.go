package db

import (
	"GO-REGRIS-LOGIN/entity"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&entity.User{})
}

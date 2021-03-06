package db

import (
	"GO-REGRIS-LOGIN/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDatabase() *gorm.DB {
	cfg, _ := config.NewConfig(".env")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Name)
	db, err := gorm.Open(postgres.Open(string(dsn)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

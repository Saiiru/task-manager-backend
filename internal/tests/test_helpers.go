package tests

import (
	"task-manager-app/backend/internal/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.Task{}, &domain.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

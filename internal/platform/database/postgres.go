package database

import (
	"maestro/internal/shared"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ...

func NewPostgresDB(dsn shared.DSN) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(string(dsn)), &gorm.Config{})
}

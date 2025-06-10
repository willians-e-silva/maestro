package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Opcional: Migrações automáticas (para desenvolvimento, cuidado em produção)
	// err = db.AutoMigrate(&domain.User{}) // Importe suas entidades aqui
	// if err != nil {
	// 	log.Printf("Falha na migração do DB: %v", err)
	// }

	log.Println("Conectado ao banco de dados PostgreSQL")
	return db, nil
}

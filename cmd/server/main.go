// cmd/server/main.go
package main

import (
	"log"

	"github.com/willians-e-silva/maestro/internal/config"
	"github.com/willians-e-silva/maestro/internal/handler/repository"
	"github.com/willians-e-silva/maestro/internal/handler/server"
	"github.com/willians-e-silva/maestro/internal/platform/database"
	usecase "github.com/willians-e-silva/maestro/internal/usecase/user"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	userRepo := repository.NewPostgresUserRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepo)

	server.ServerGrpc(cfg.Port, userUsecase)
}

// cmd/server/main.go
package main

import (
	"log"

	"github.com/willians-e-silva/maestro/internal/adapter/repository"
	"github.com/willians-e-silva/maestro/internal/adapter/server"
	"github.com/willians-e-silva/maestro/internal/config"
	"github.com/willians-e-silva/maestro/internal/platform/database"
	"github.com/willians-e-silva/maestro/internal/usecase"
)

func main() {
	// Carregar configurações (ex: do .env)
	cfg := config.LoadConfig()

	// 1. Inicializar o Banco de Dados (Plataforma/Infraestrutura)
	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}

	// 2. Inicializar o Repositório (Adaptador)
	userRepo := repository.NewPostgresUserRepository(db)

	// 3. Inicializar os Casos de Uso (Core da Aplicação)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// 4. Iniciar o Servidor GRPC (Adaptador)
	server.ServerGrpc(cfg.Port, userUsecase)
}

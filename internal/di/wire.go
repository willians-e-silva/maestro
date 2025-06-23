//go:build wireinject
// +build wireinject

package di

import (
	"maestro/internal/config"
	"maestro/internal/handler/repository"
	"maestro/internal/handler/server"
	"maestro/internal/platform/database"
	"maestro/internal/shared"
	taskusecase "maestro/internal/usecase/task"

	"github.com/google/wire"
)

// Tipos distintos para evitar conflito de bindings com string
type Port string
type DSN string

func InitializeServer() (*server.GRPCServer, error) {
	wire.Build(
		config.LoadConfig,
		provideDSN,
		providePort,
		database.NewPostgresDB,
		repository.NewPostgresTaskRepository,
		taskusecase.NewTaskUsecase,
		server.NewGRPCServer,
	)

	return &server.GRPCServer{}, nil
}

func providePort(cfg *config.Config) shared.Port {
	return shared.Port(":" + cfg.Port)
}

func provideDSN(cfg *config.Config) shared.DSN {
	return shared.DSN(cfg.DatabaseURL)
}

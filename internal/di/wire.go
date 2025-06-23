package di

import (
	"maestro/internal/config"
	"maestro/internal/handler/repository"
	"maestro/internal/handler/server"
	"maestro/internal/platform/database"
	taskusecase "maestro/internal/usecase/task"

	"github.com/google/wire"
)

func InitializeServer() (*server.GRPCServer, error) {
	wire.Build(
		config.LoadConfig,
		database.NewPostgresDB,
		repository.NewPostgresTaskRepository,
		taskusecase.NewTaskUsecase,
		providePort,
		server.NewGRPCServer,
	)
	return &server.GRPCServer{}, nil
}

func providePort(cfg *config.Config) string {
	return cfg.Port
}

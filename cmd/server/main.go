// cmd/server/main.go
package main

import (
	"log"

	"maestro/internal/config"
	"maestro/internal/handler/repository"
	"maestro/internal/handler/server"
	"maestro/internal/platform/database"
	taskusecase "maestro/internal/usecase/task"
	userusecase "maestro/internal/usecase/user"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDB, _ := db.DB()
	database.RunMigrations(sqlDB)

	userRepo := repository.NewPostgresUserRepository(db)
	taskRepo := repository.NewPostgresTaskRepository(db)

	userUsecase := userusecase.NewUserUsecase(userRepo)
	taskUsecase := taskusecase.NewTaskUsecase(taskRepo)

	server.ServerGrpc(cfg.Port, userUsecase, taskUsecase)
}

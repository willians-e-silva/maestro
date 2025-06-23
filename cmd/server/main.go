package main

import (
	"log"
	"maestro/internal/di"
)

func main() {
	srv, err := di.InitializeServer()
	if err != nil {
		log.Fatalf("Erro ao inicializar o servidor: %v", err)
	}

	srv.Start()
}

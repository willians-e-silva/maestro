// server/server.go
package server

import (
	"log"
	"net"

	taskpb "maestro/internal/infra/grpc/task"
	userpb "maestro/internal/infra/grpc/user"

	taskusecase "maestro/internal/usecase/task"
	userusecase "maestro/internal/usecase/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ServerGrpc(port string, userUsecase *userusecase.UserUsecase, taskUsecase *taskusecase.TaskUsecase) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Falha ao iniciar o listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	userpb.RegisterUserServiceServer(grpcServer, userUsecase)
	taskpb.RegisterTaskServiceServer(grpcServer, taskUsecase)

	reflection.Register(grpcServer)

	log.Printf("Servidor gRPC iniciado na porta %s", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Falha ao iniciar o servidor gRPC: %v", err)
	}
}

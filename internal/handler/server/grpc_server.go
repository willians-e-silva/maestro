package server

import (
	"log"
	"net"

	taskpb "maestro/internal/infra/grpc/task"
	"maestro/internal/shared"
	taskusecase "maestro/internal/usecase/task"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	taskUsecase *taskusecase.TaskUsecase
	port        string
}

func NewGRPCServer(taskUsecase *taskusecase.TaskUsecase, port shared.Port) *GRPCServer {
	return &GRPCServer{
		taskUsecase: taskUsecase,
		port:        string(port),
	}
}

func (s *GRPCServer) Start() {
	listener, err := net.Listen("tcp", s.port)
	if err != nil {
		log.Fatalf("Falha ao iniciar o listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	taskpb.RegisterTaskServiceServer(grpcServer, s.taskUsecase)
	reflection.Register(grpcServer)

	log.Printf("Servidor gRPC iniciado na porta %s", s.port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Falha ao iniciar o servidor gRPC: %v", err)
	}
}

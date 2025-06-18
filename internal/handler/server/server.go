package server

import (
	"log"
	"net"

	pb "github.com/willians-e-silva/maestro/internal/infra/grpc/user"
	usecase "github.com/willians-e-silva/maestro/internal/usecase/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ServerGrpc(port string, userUsecase *usecase.UserUsecase) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Falha ao iniciar o listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userUsecase)

	reflection.Register(grpcServer)

	log.Printf("Servidor gRPC iniciado na porta %s", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Falha ao iniciar o servidor gRPC: %v", err)
	}
}

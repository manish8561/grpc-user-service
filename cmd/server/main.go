package main

import (
	"log"
	"net"

	handler "github.com/manish8561/grpc-user-service/internal/handler/grpc"
	"github.com/manish8561/grpc-user-service/internal/repository"
	"github.com/manish8561/grpc-user-service/internal/service"
	"github.com/manish8561/grpc-user-service/pb"
	"google.golang.org/grpc"
)

// This is the main function for the server
// It creates a new listener on port 50051
// It then creates a new gRPC server
// It then creates a new user repository and user service
// It then creates a new user handler
// It then registers the user service with the gRPC server
// It then logs a message to the console indicating that the server is running
// It then serves the gRPC server on the listener
// It then logs a message to the console indicating that the server failed to serve
func main() {
	// Create a new listener on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Dependency Injection
	repo := repository.NewUserRepository()
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	pb.RegisterUserServiceServer(grpcServer, h)

	log.Println("gRPC server running on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

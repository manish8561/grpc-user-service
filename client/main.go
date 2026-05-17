package main

import (
	"context"
	"log"
	"time"

	"github.com/manish8561/grpc-user-service/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// This is the main function for the client
// It creates a new connection to the server and a new client
// It then creates a new context with a timeout of 500 milliseconds
// It then calls the GetUser method on the client
// It then prints the response
func main() {

	// Create a new connection to the server
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		500*time.Millisecond,
	)

	defer cancel()

	resp, err := client.GetUser(ctx, &pb.GetUserRequest{
		Id: 1,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(resp)
}

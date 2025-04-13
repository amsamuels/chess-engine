package main

import (
	pb "chess-engine/gen"                       // Update this path to match your actual module
	cheServer "chess-engine/internal/cheServer" // Update this to wherever your server code lives
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 1. Listen on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 2. Create a new gRPC server
	grpcServer := grpc.NewServer()

	// 3. Register your chess server
	gameServer := cheServer.New()
	pb.RegisterChessServiceServer(grpcServer, gameServer)

	// Register reflection for grpcurl to work
	reflection.Register(grpcServer)

	log.Println("Chess gRPC server is listening on :50051...")

	// 4. Start serving
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

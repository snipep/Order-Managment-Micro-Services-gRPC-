package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewGRPCClient creates a new grpc client connection
func NewGRPCClient(addr string) *grpc.ClientConn {
	// NewClient creates a new grpc client connection
	// grpc.WithTransportCredentials creates a new transport credentials
	// insecure.NewCredentials() creates a new insecure transport credentials
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil{
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

// main starts the http server
func main() {
	// NewHttpServer creates a new http server
	httpServer := NewHttpServer(":1000")
	// Run starts the http server
	httpServer.Run() 
}

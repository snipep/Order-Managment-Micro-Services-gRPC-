package main

import (
	"log"
	"net"

	handler "github.com/snipep/kitchen/services/orders/handler/orders"
	service "github.com/snipep/kitchen/services/orders/service"
	grpc "google.golang.org/grpc"
)

// gRPCServer is the struct that holds the address of the gRPC server
type gRPCServer struct {
	addr string
}

// NewGRPCServer creates a new gRPC server
func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{
		addr: addr,
	}
}

// Run starts the gRPC server and listens on the given address
func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer()

	// register our grpc services
	OrderService := service.NewOrderService()
	handler.NewGRPCOrdersService(grpcServer, OrderService)

	// register our grpc services
	log.Println("starting gRPC server on", s.addr)
	return grpcServer.Serve(lis)
}

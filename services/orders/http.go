package main

import (
	"log"
	"net/http"

	handler "github.com/snipep/kitchen/services/orders/handler/orders"
	"github.com/snipep/kitchen/services/orders/service"
)

// httpServer is the struct that holds the address of the server
type httpServer struct {
	addr string
}

// NewHttpServer creates a new http server
func NewHttpServer(add string) *httpServer {
	return &httpServer{
		addr: add,
	}
}

// Run starts the http server and listens on the given address
func (s *httpServer) Run() error {
	// create a new router
	router := http.NewServeMux()

	// create a new order service
	orderService := service.NewOrderService()

	// create a new http order handler
	orderHandler := handler.NewHttpOrderHandler(orderService)

	// register the router with the order handler
	orderHandler.RegisterRouter(router)

	// start the http server
	log.Println("starting http server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

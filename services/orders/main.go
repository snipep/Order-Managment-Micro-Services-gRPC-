package main

// main function starts the http server and the grpc server
func main() {
	// start the http server on port 8000
	httpServer := NewHttpServer(":8000")
	go httpServer.Run()

	// start the grpc server on port 9000
	gRPCServer := NewGRPCServer(":9000")
	gRPCServer.Run()
}


package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/snipep/kitchen/services/common/genproto/orders"
)

// httpServer is the struct that holds the address of the server
type httpServer struct {
	addr string
}

// NewHttpServer creates a new http server
func NewHttpServer(addr string) *httpServer {
	return &httpServer{
		addr: addr,
	}
}

// Run starts the http server and listens on the given address
func (s *httpServer) Run() error {
	router := http.NewServeMux()

	// Connect to the gRPC server
	conn := NewGRPCClient(":9000")
	defer conn.Close()

	// Handle the "/" path
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Create a new client with the gRPC connection
		c := orders.NewOrderServiceClient(conn)

		// Create a new context with a timeout of 2 seconds
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		// Create an order and check if it was created successfully
		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 24,
			ProductID:  3123,
			Quantity:   1,
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		// Get the list of orders from the gRPC server
		res, err := c.GetOrders(ctx, &orders.GetOrdersRequest{
			CustomerID: 42,
		})
		if err != nil {
			log.Fatalf("client error: %v", err)
		}

		// Parse the template and execute it with the orders list
		t := template.Must(template.New("orders").Parse(ordersTemplate))
		if err := t.Execute(w, res.GetOrders()); err != nil {
			log.Fatalf("template error: %v", err)
		}
	})

	log.Println("starting http server on", s.addr)
	return http.ListenAndServe(s.addr, router)
}

// ordersTemplate is the template for the list of orders
var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>
`


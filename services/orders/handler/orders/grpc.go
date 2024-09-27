package handler

import (
	"context"

	orders "github.com/snipep/kitchen/services/common/genproto/orders"
	types "github.com/snipep/kitchen/services/orders/types"
	grpc "google.golang.org/grpc"
)

// OrderGRPCHandler is the handler for the orders gRPC service
type OrderGRPCHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

// NewGRPCOrdersService creates a new handler for the orders gRPC service
func NewGRPCOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrderGRPCHandler{
		ordersService: ordersService,
	}
	//register the OrderServiceServer
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)
}

// GetOrders returns all orders from the database
func (h *OrderGRPCHandler) GetOrders(ctx context.Context, req *orders.GetOrdersRequest) (*orders.GetOrdersResponse, error) {
	o := h.ordersService.GetOrder(ctx)
	res := &orders.GetOrdersResponse{
		Orders: o,
	}
	return res, nil
}

// CreateOrder creates a new order and adds it to the database
func (h *OrderGRPCHandler) CreateOrder(ctx context.Context, req *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    42,
		CustomerID: 2,
		ProductID:  1,
		Quantity:   10,
	}

	err := h.ordersService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	res := &orders.CreateOrderResponse{
		Status: "success",
	}
	return res, nil
}


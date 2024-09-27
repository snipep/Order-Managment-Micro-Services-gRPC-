package service

import (
	"context"

	"github.com/snipep/kitchen/services/common/genproto/orders"
)
// ordersDB is the in-memory database for orders
var ordersDB = make([]*orders.Order, 0)

// OrderService is the interface for the order service
type OrderService struct {
	// store is the in-memory database for orders
}

// NewOrderService creates a new order service
func NewOrderService() *OrderService {
	return &OrderService{}
}

// CreateOrder creates a new order and adds it to the database
func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDB = append(ordersDB, order)
	return nil
}

// GetOrder returns all orders from the database
func (s *OrderService) GetOrder(ctx context.Context) []*orders.Order {
	return ordersDB
}

package types

import (
	"context"

	orders "github.com/snipep/kitchen/services/common/genproto/orders"
)

// OrderService is the interface for the order service
type OrderService interface {
	// CreateOrder creates a new order and adds it to the database
	CreateOrder(context.Context, *orders.Order) error
	// GetOrder returns all orders from the database
	GetOrder(context.Context) []*orders.Order	
}

package handler

import (
	"net/http"

	orders "github.com/snipep/kitchen/services/common/genproto/orders"
	"github.com/snipep/kitchen/services/common/util"
	"github.com/snipep/kitchen/services/orders/types"
)

// OrdersHttpHandler is the struct that holds the order service
// and is used to register the router for the order service
type OrdersHttpHandler struct {
	orderService types.OrderService
}

// NewHttpOrderHandler creates a new http order handler
func NewHttpOrderHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		orderService: orderService,
	}
	return handler
}

// RegisterRouter registers the router for the order service
func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	// handle the POST /orders path
	router.HandleFunc("POST /orders", h.CreateOrder)
}

// CreateOrder handles the POST /orders path
func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req orders.CreateOrderRequest
	err := util.ParseJSON(r, &req)
	if err != nil {
		// return an error if the json could not be parsed
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	order := &orders.Order{
		OrderID:    42,
		CustomerID: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	err = h.orderService.CreateOrder(r.Context(), order)
	if err != nil {
		// return an error if the order could not be created
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	res := &orders.CreateOrderResponse{Status: "success"}
	// return a success response
	util.WriteJSON(w, http.StatusOK, res)
}


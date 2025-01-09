package service

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	orderv1alpha1 "github.com/cuminandpaprika/go-monorepo-example/gen/order/v1alpha1"
)

type OrderService struct {
	orders map[string]*orderv1alpha1.Order
}

func NewOrderService() *OrderService {
	return &OrderService{
		orders: make(map[string]*orderv1alpha1.Order),
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orderv1alpha1.Order) (*orderv1alpha1.Order, error) {
	if order.Id == "" {
		return nil, errors.New("order ID is required")
	}
	if _, exists := s.orders[order.Id]; exists {
		return nil, fmt.Errorf("order with ID %s already exists", order.Id)
	}
	s.orders[order.Id] = order
	return order, nil
}

func (s *OrderService) GetOrder(ctx context.Context, id string) (*orderv1alpha1.Order, error) {
	order, exists := s.orders[id]
	if !exists {
		return nil, fmt.Errorf("order with ID %s not found", id)
	}
	return order, nil
}

type OrderServiceHandler struct {
	service *OrderService
}

func NewOrderServiceHandler(service *OrderService) *OrderServiceHandler {
	return &OrderServiceHandler{service: service}
}

func (h *OrderServiceHandler) CreateOrder(ctx context.Context, req *connect.Request[orderv1alpha1.CreateOrderRequest]) (*connect.Response[orderv1alpha1.CreateOrderResponse], error) {
	order, err := h.service.CreateOrder(ctx, req.Msg.Order)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return connect.NewResponse(&orderv1alpha1.CreateOrderResponse{Order: order}), nil
}

func (h *OrderServiceHandler) GetOrder(ctx context.Context, req *connect.Request[orderv1alpha1.GetOrderRequest]) (*connect.Response[orderv1alpha1.GetOrderResponse], error) {
	order, err := h.service.GetOrder(ctx, req.Msg.Id)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
	return connect.NewResponse(&orderv1alpha1.GetOrderResponse{Order: order}), nil
}

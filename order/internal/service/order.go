package service

import (
	"context"
	"errors"
	"fmt"
	"log"

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
	if s == nil {
		log.Println("OrderService is nil")
		return nil, errors.New("OrderService is nil")
	}
	if order == nil {
		log.Println("Order is nil")
		return nil, errors.New("order is nil")
	}
	log.Printf("Creating order with ID: %s", order.Id)
	if order.Id == "" {
		return nil, errors.New("order ID is required")
	}
	if _, exists := s.orders[order.Id]; exists {
		return nil, fmt.Errorf("order with ID %s already exists", order.Id)
	}
	s.orders[order.Id] = order
	log.Printf("Order with ID: %s created successfully", order.Id)
	return order, nil
}

func (s *OrderService) GetOrder(ctx context.Context, id string) (*orderv1alpha1.Order, error) {
	if s == nil {
		log.Println("OrderService is nil")
		return nil, errors.New("OrderService is nil")
	}
	log.Printf("Fetching order with ID: %s", id)
	order, exists := s.orders[id]
	if !exists {
		return nil, fmt.Errorf("order with ID %s not found", id)
	}
	log.Printf("Order with ID: %s fetched successfully", id)
	return order, nil
}

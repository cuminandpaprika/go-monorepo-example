package service

import (
	"context"
	"errors"
	"log"

	"connectrpc.com/connect"
	orderv1alpha1 "github.com/cuminandpaprika/go-monorepo-example/gen/order/v1alpha1"
)

type OrderServiceHandler struct {
	service *OrderService
}

func NewOrderServiceHandler(service *OrderService) *OrderServiceHandler {
	if service == nil {
		panic("service cannot be nil")
	}
	return &OrderServiceHandler{service: service}
}

func (h *OrderServiceHandler) CreateOrder(ctx context.Context, req *connect.Request[orderv1alpha1.CreateOrderRequest]) (*connect.Response[orderv1alpha1.CreateOrderResponse], error) {
	if h == nil || h.service == nil {
		log.Println("OrderServiceHandler or its service is nil")
		return nil, connect.NewError(connect.CodeInternal, errors.New("OrderServiceHandler or its service is nil"))
	}
	if req.Msg == nil || req.Msg.Order == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("request or order is nil"))
	}
	log.Printf("Received CreateOrder request for order ID: %s", req.Msg.Order.Id)
	order, err := h.service.CreateOrder(ctx, req.Msg.Order)
	if err != nil {
		log.Printf("Error creating order: %v", err)
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	log.Printf("CreateOrder request for order ID: %s processed successfully", req.Msg.Order.Id)
	return connect.NewResponse(&orderv1alpha1.CreateOrderResponse{Order: order}), nil
}

func (h *OrderServiceHandler) GetOrder(ctx context.Context, req *connect.Request[orderv1alpha1.GetOrderRequest]) (*connect.Response[orderv1alpha1.GetOrderResponse], error) {
	if h == nil || h.service == nil {
		log.Println("OrderServiceHandler or its service is nil")
		return nil, connect.NewError(connect.CodeInternal, errors.New("OrderServiceHandler or its service is nil"))
	}
	if req.Msg == nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("request is nil"))
	}
	log.Printf("Received GetOrder request for order ID: %s", req.Msg.Id)
	order, err := h.service.GetOrder(ctx, req.Msg.Id)
	if err != nil {
		log.Printf("Error fetching order: %v", err)
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
	log.Printf("GetOrder request for order ID: %s processed successfully", req.Msg.Id)
	return connect.NewResponse(&orderv1alpha1.GetOrderResponse{Order: order}), nil
}

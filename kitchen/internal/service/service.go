package service

import (
	"context"

	kitchenpb "github.com/cuminandpaprika/go-monorepo-example/gen/kitchen/v1alpha1"
)

func New() kitchenpb.KitchenServiceServer {
	return &service{}
}

type service struct {
	kitchenpb.UnimplementedKitchenServiceServer
}

func (s *service) CookFood(ctx context.Context, req *kitchenpb.CookFoodRequest) (*kitchenpb.CookFoodResponse, error) {
	// Implement your logic here
	return &kitchenpb.CookFoodResponse{
		Status:  "success",
		Message: "Food cooked successfully",
	}, nil
}

func (s *service) PrepareFood(ctx context.Context, req *kitchenpb.PrepareFoodRequest) (*kitchenpb.PrepareFoodResponse, error) {
	// Implement your logic here
	return &kitchenpb.PrepareFoodResponse{
		Status:  "success",
		Message: "Food prepared successfully",
	}, nil
}

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	kitchenpb "github.com/cuminandpaprika/go-monorepo-example/gen/kitchen/v1alpha1"
	"github.com/cuminandpaprika/go-monorepo-example/kitchen/internal/service"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type ExampleRouter struct {
	*mux.Router
}

func NewExampleRouter() *ExampleRouter {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./web"))
	r.PathPrefix("/").Handler(fs)

	return &ExampleRouter{
		Router: r,
	}
}

func main() {
	fmt.Println("hello world! Kitchen service")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	healthcheck := health.NewServer()
	healthcheck.SetServingStatus("kitchenpb.KitchenService", healthgrpc.HealthCheckResponse_SERVING)
	healthgrpc.RegisterHealthServer(s, healthcheck)
	kitchenpb.RegisterKitchenServiceServer(s, service.New())
	reflection.Register(s)

	log.Println("Serving gRPC on port 8080")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	kitchenpb "github.com/cuminandpaprika/go-monorepo-example/gen/kitchen/v1alpha1"
	"github.com/cuminandpaprika/go-monorepo-example/kitchen/internal/service"
	"github.com/gorilla/mux"
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

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	kitchenpb.RegisterKitchenServiceServer(s, service.New())
	reflection.Register(s)

	log.Println("Serving gRPC on port 8000")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

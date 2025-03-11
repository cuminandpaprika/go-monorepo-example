package main

import (
	"fmt"
	"log"
	"net/http"

	kitchenv1alpha1 "github.com/cuminandpaprika/go-monorepo-example/gen/kitchen/v1alpha1"
	orderv1alpha1connect "github.com/cuminandpaprika/go-monorepo-example/gen/order/v1alpha1/orderv1alpha1connect"
	"github.com/cuminandpaprika/go-monorepo-example/order/internal/service"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
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
	fmt.Println("hello world! Fast")
	http.Handle("/", NewExampleRouter())

	// Connect to KitchenService
	conn, err := grpc.Dial("kitchen:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to KitchenService: %v", err)
	}
	defer conn.Close()
	kc := kitchenv1alpha1.NewKitchenServiceClient(conn)

	orderService := service.NewOrderService(kc)
	orderServiceHandler := service.NewOrderServiceHandler(orderService)

	path, handler := orderv1alpha1connect.NewOrderServiceHandler(orderServiceHandler)
	http.Handle(path, handler)

	log.Println("Serving on port 8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatalf("Server exited with: %v", err)
	}
}

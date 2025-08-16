package cmd

import (
	middleware "ecommerce/Middleware"
	"ecommerce/global_router"
	"ecommerce/hendlers"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	mux := http.NewServeMux()
	manager.Use(middleware.Logger,middleware.Hudai)

	// mux.Handle("GET /route", middleware.Logger(http.HandlerFunc(hendlers.Test)))

	mux.Handle("GET /route", manager.With(
		http.HandlerFunc(hendlers.Test),
	))

	// mux.Handle("GET /Products", http.HandlerFunc(getProduct))
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(hendlers.GetProducts),
	
	))

	mux.Handle("GET /products/{productId}", manager.With(
		http.HandlerFunc(hendlers.GetProductById),
	
	))

	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(hendlers.CreateProduct),
	
	))

	fmt.Println("server running on :8080")

	GlobarRouter := global_router.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", GlobarRouter)
	if err != nil {
		fmt.Println("Error Starting the server", err)
	}

}

package main

import (
	"assignment-dua/handler"
	"assignment-dua/repository"
	"assignment-dua/service"
	"assignment-dua/utils"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	pgpool, err := utils.NewPostgresPool(
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully connected to database")

	r := mux.NewRouter()

	itemRepo := repository.NewItemRepository(pgpool)
	orderRepo := repository.NewOrdersRepository(pgpool)
	serviceOrder := service.NewService(itemRepo, orderRepo)

	handlerOrder := handler.NewHandler(serviceOrder)

	r.HandleFunc("/order", handlerOrder.CreateOrders).Methods("POST")
	r.HandleFunc("/order", handlerOrder.GetAllData).Methods("GET")
	r.HandleFunc("/order/{id}", handlerOrder.DeleteData).Methods("DELETE")
	r.HandleFunc("/order/{id}", handlerOrder.UpdateData).Methods("PUT")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

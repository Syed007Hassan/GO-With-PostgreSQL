package router

import (
	"go-postgres/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Roouter{

	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}",middleware.GetStock).Methods("GET","OPTIONS");
	router.HandleFunc("/api/stocks",middleware.getAllStocks).Methods("GET","OPTIONS");
	router.HandleFunc("/api/stock",middleware.CreateStock).Methods("POST","OPTIONS");
	router.HandleFunc("/api/stock/{id}",middleware.UpdateStock).Methods("PUT","OPTIONS");
	router.HandleFunc("/api/stock/{id}",middleware.DeleteStock).Methods("DELETE","OPTIONS");

}
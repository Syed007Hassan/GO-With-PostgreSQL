package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-postgres/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Response is a struct that contains a message and an ID
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	errr := db.Ping()

	if errr != nil {
		panic(errr)
	}

	fmt.Println("Successfully connected to database!")

	return db

}

// CreateStock is a function that creates a stock

func CreateStock(w http.ResponseWriter, r *http.Request) {

	var stock models.Stock

	// Decode the request body into the struct and if there is an error, return a 400 bad request
	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatal("Error decoding the request body", err)
	}

	insertID := InsertStock(stock)

	// Format a response object
	res := response{
		ID:      insertID,
		Message: "Stock created successfully",
	}

	// Send the response
	json.NewEncoder(w).Encode(res)

}

// GetStock is a function that gets a stock

func GetStock(w http.ResponseWriter, r *http.Request) {

	// Get the id from the incoming url
	params := mux.Vars(r)

	// Convert the id from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	// Call the GetStock function with the id to retrieve a single stock
	stock, err := getStock(int64(id))

	if err != nil {
		log.Fatalf("Unable to get stock. %v", err)
	}

	// Send the response
	json.NewEncoder(w).Encode(stock)

}

// GetAllStocks is a function that gets all stocks

func GetAllStocks(w http.ResponseWriter, r *http.Request) {

	// Call the GetAllStocks function to retrieve all stocks
	stocks, err := getAllStocks()

	if err != nil {
		log.Fatalf("Unable to get all stocks. %v", err)
	}

	// Send the response
	json.NewEncoder(w).Encode(stocks)
}

// UpdateStock is a function that updates a stock

func UpdateStock(w http.ResponseWriter, r *http.Request) {

	// get the id from the incoming url
	params := mux.Vars(r)

	// convert the id from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	var stock models.Stock

	// decode the request body into the struct and if there is an error, return a 400 bad request
	err = json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatal("Error decoding the request body", err)
	}

	// call the updateStock function with the id to update a single stock
	updatedRows := updateStock(int64(id), stock)

	// format a response object
	msg := fmt.Sprintf("Stock updated successfully. Total rows/record affected %v", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

// DeleteStock is a function that deletes a stock

func DeleteStock() {

}

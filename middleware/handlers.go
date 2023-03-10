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

	_ "github.com/lib/pq"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Response is a struct that contains a message and an ID
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func createConnection() *sql.DB {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		log.Fatal("Error due to environment variables not set properly. Please check the .env file.")
		panic(err)
	}

	errr := db.Ping()

	if errr != nil {
		log.Fatal("Error in ping")
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

	insertID := insertStock(stock)

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

func DeleteStock(w http.ResponseWriter, r *http.Request) {

	// get the id from incoming url
	params := mux.Vars(r)

	// convert the id from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	// call the deleteStock function with the id to delete a single stock
	deletedRows := deleteStock(int64(id))

	// format a response object
	msg := fmt.Sprintf("Stock updated successfully. Total rows/record affected %v", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

func insertStock(stock models.Stock) int64 {

	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO stocks (name,price,company) VALUES ($1, $2, $3) RETURNING stockid`
	var id int64
	err := db.QueryRow(sqlStatement, stock.Name, stock.Price, stock.Company).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query %v", err)
	}
	fmt.Println("New record ID is:", id)
	return id

}

func getStock(id int64) (models.Stock, error) {

	db := createConnection()
	defer db.Close()

	sqlStatement := `SELECT * FROM stocks WHERE stockid=$1`
	var stock models.Stock
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return stock, nil
	case nil:
		return stock, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return stock, err

}

func getAllStocks() ([]models.Stock, error) {

	db := createConnection()
	defer db.Close()

	var stocks []models.Stock
	sqlStatement := `Select * from stocks`

	// Execute the SQL statement and get the rows in return
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	// Iterate over the rows and append the stocks in the slice
	for rows.Next() {
		var stock models.Stock

		// Unmarshal the row object to stock struct
		err = rows.Scan(&stock.StockID, &stock.Name, &stock.Price, &stock.Company)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// Append the stock in the slice stocks
		stocks = append(stocks, stock)
	}

	return stocks, err

}

func updateStock(id int64, stock models.Stock) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `UPDATE stocks SET name=$2, price=$3, company=$4 WHERE stockid=$1`

	res, err := db.Exec(sqlStatement, id, stock.Name, stock.Price, stock.Company)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected

}

func deleteStock(id int64) int64 {

	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM stocks WHERE stockid=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected

}

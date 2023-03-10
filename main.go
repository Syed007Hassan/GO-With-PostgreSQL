package main

import (
	"fmt"
	"go-postgres/router"
	"log"
	"net/http"
)

func main(){
	r := router.Router()

	fmt.Println("Server running on port 5000...")

	log.Fatal(http.ListenAndServe(":5000",r))
}
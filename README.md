# GO-With-PostgresSQL

* Implemented REST APIs that perform CRUD operations with PostgreSQL 

```
router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
router.HandleFunc("/api/stock", middleware.GetAllStocks).Methods("GET", "OPTIONS")
router.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")
router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
router.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")
```

<p align="centre" >
  
  <img src="https://user-images.githubusercontent.com/104893311/224471343-f439402f-7006-4b24-bc50-3c4e8379414d.png">

</p>  

<p align="centre" >
  
  <img src="https://user-images.githubusercontent.com/104893311/224471468-e23ebf82-fc56-4fc7-9393-3ece766d50dc.png">

</p>  

<p align="centre" >
  
  <img src="https://user-images.githubusercontent.com/104893311/224471599-48fd0da1-b80c-4542-9462-00376acf83e9.png">

</p> 

<p align="centre" >
  
  <img src="https://user-images.githubusercontent.com/104893311/224471753-e72039d4-fce5-4a45-815f-d356f6cee781.png">

</p> 

<p align="centre" >
  
  <img src="https://user-images.githubusercontent.com/104893311/224471807-43c9cbfd-87cc-47a7-b86d-d7a54852a3aa.png">

</p> 


* To configure DB connection use:

```
POSTGRES_URL="postgres://<username>:<password>@localhost:5432/<db_name>?sslmode=disable"
```

* To run:

```
go run main.go
```

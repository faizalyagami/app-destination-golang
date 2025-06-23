package main

import (
	"app-destination/database"
	"app-destination/router"
	"log"
	"net/http"
)

func main() {
	database.InitDB()

	r := router.NewRouter()

	log.Println("Starting server on port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
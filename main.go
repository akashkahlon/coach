package main

import (
	"coach/api"
	"coach/db"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	  err := godotenv.Load()
	  if err != nil {
			log.Fatalf("error loading .env file: %v", err)
	  }
		
    db, err := db.InitDB()
		if err != nil {
				log.Fatalf("could not connect to the database: %v", err)
		}
		defer db.Close()

    http.HandleFunc("/health", api.HealthCheckHandler)

    log.Println("Server is running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
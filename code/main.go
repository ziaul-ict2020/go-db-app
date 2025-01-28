// File: code/main.go

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Collect PostgreSQL connection parameters from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("One or more required environment variables are missing")
	}

	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		host, port, user, password, dbname)

	// Connect to the database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v", err)
	}
	defer db.Close()

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping the database: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL database!")

	// HTTP handler to test the database
	http.HandleFunc("/checkdb", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"status": "error", "message": "Database connection failed"})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "success", "message": "Database is working fine"})
	})

	// Start the server
	httpPort := os.Getenv("APP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	fmt.Printf("Starting server on port %s\n", httpPort)
	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}

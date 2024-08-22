package main

import (
	"database/sql"
	"embed"
	//"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

//go:embed vue_this/dist/*
var static embed.FS

func main() {
	// 1. Configuration and Error Handling:
	// Use environment variables for sensitive information (database credentials)
	// and provide defaults for development.
	// Centralize configuration for easier management.
	config := struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getIntEnv("DB_PORT", 5432),
		User:     getEnv("DB_USER", "jim"),
		Password: getEnv("DB_PASSWORD", "whatsimportantnow"),
		DBName:   getEnv("DB_NAME", "dateapp"),
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database:", err)
	}

	fmt.Println("Successfully connected to database!")

	// 2. API Routing:
	// Use a separate function to handle API routes, making the code more modular.
	// Handle errors explicitly within the API route handler.
	http.HandleFunc("/api/data", handleAPIData(db))

	// 3. Static File Serving:
	// Handle errors when setting up the file server.
	fsys, err := fs.Sub(static, "vue_this/dist")
	if err != nil {
		log.Fatal("Error setting up file server:", err)
	}
	http.Handle("/", http.FileServer(http.FS(fsys)))

	// 4. Server Startup:
	// Log the server startup message with the actual port number.
	port := getIntEnv("PORT", 8080)
	log.Printf("Server starting on http://localhost:%d", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

// Helper functions to get environment variables with defaults
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getIntEnv(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	// ... (add error handling for parsing the integer)
	return defaultValue // Replace with actual parsed value
}

// API route handler function
func handleAPIData(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ... (rest of your API route logic remains the same)
	}
}

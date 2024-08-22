package main

import (
	"database/sql"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

//go:embed vue_this/dist/*
var static embed.FS

const (
	host     = "localhost"
	port     = 5432
	user     = "jim"
	password = "whatsimportantnow"
	dbname   = "dateapp"
)

func main() {
	// Set up database connection
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to database!")

	// Set up file server for frontend
	fsys, err := fs.Sub(static, "vue_this/dist")
	if err != nil {
		log.Fatal(err)
	}

	// Set up API routes
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request to /api/data")
		rows, err := db.Query("SELECT id, name FROM users_main LIMIT 10")
		if err != nil {
			log.Printf("Error querying database: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var results []map[string]interface{}
		for rows.Next() {
			var id int
			var name string
			if err := rows.Scan(&id, &name); err != nil {
				log.Printf("Error scanning row: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			results = append(results, map[string]interface{}{
				"id":   id,
				"name": name,
			})
			log.Printf("Retrieved row: id=%d, name=%s", id, name)
		}

		log.Printf("Returning %d results", len(results))

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	})

	// Serve static files
	http.Handle("/", http.FileServer(http.FS(fsys)))

	// Start the server
	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

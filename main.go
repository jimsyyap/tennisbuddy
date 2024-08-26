package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//go:embed vue_this/dist/*
var static embed.FS

// Define your models
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

type Location struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Address string `gorm:"not null"`
	City    string `gorm:"not null"`
	State   string `gorm:"not null"`
	Country string `gorm:"not null"`
}

type Match struct {
	gorm.Model
	Player1ID  uint
	Player2ID  uint
	LocationID uint
	MatchDate  string `gorm:"not null"`
	Score      string
	WinnerID   uint
	Player1    User     `gorm:"foreignKey:Player1ID"`
	Player2    User     `gorm:"foreignKey:Player2ID"`
	Location   Location `gorm:"foreignKey:LocationID"`
	Winner     User     `gorm:"foreignKey:WinnerID"`
}

type Tournament struct {
	gorm.Model
	Name       string `gorm:"not null"`
	StartDate  string `gorm:"not null"`
	EndDate    string `gorm:"not null"`
	LocationID uint
	Location   Location `gorm:"foreignKey:LocationID"`
}

func main() {
	// Configuration
	config := struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
	}{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "jim"),
		Password: getEnv("DB_PASSWORD", "whatsimportantnow"),
		DBName:   getEnv("DB_NAME", "tennisbuddy"),
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Australia/Melbourne",
		config.Host, config.User, config.Password, config.DBName, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Controlled migration
	err = performMigration(db)
	if err != nil {
		log.Fatal("Failed to perform migration:", err)
	}

	fmt.Println("Successfully connected to database and migrated schemas!")

	// Set up CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	// API Routing
	mux := http.NewServeMux()
	mux.HandleFunc("/api/hello", handleHello(db))

	// Wrap mux with CORS middleware
	handler := c.Handler(mux)

	// Server Startup
	port := getEnv("PORT", "8080")
	log.Printf("Server starting on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

func performMigration(db *gorm.DB) error {
	// Check if the tables exist
	if db.Migrator().HasTable(&User{}) &&
		db.Migrator().HasTable(&Location{}) &&
		db.Migrator().HasTable(&Match{}) &&
		db.Migrator().HasTable(&Tournament{}) {
		fmt.Println("Tables already exist, skipping migration")
		return nil
	}

	// If tables don't exist, create them
	err := db.AutoMigrate(&User{}, &Location{}, &Match{}, &Tournament{})
	if err != nil {
		return err
	}

	// Add any additional migration steps here if needed
	// For example, adding indexes:
	err = db.Exec("CREATE INDEX IF NOT EXISTS idx_users_username ON users(username)").Error
	if err != nil {
		return err
	}

	err = db.Exec("CREATE INDEX IF NOT EXISTS idx_users_email ON users(email)").Error
	if err != nil {
		return err
	}

	return nil
}

// Helper function to get environment variables with defaults
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// Hello World API endpoint
func handleHello(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var count int64
		db.Model(&User{}).Count(&count)

		response := map[string]string{
			"message": fmt.Sprintf("Hello World! Connected to tennisbuddy database. There are %d users.", count),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

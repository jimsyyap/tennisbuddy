package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

//go:embed vue_this/dist/*
var static embed.FS

// User model (update to include json tags)
type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"` // "-" means this field is not marshalled to JSON
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

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
	r := mux.NewRouter()
	r.HandleFunc("/api/hello", handleHello(db)).Methods("GET")

	// CRUD routes
	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

	// Wrap router with CORS middleware
	handler := c.Handler(r)

	// Server Startup
	port := getEnv("PORT", "8080")
	log.Printf("Server starting on http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal("Error starting server:", err)
	}
}

/*
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

	// crud stuff
	r := mux.NewRouter()

	r.HandleFunc("/api/users", getUsers).Methods("GET")
	r.HandleFunc("/api/users", createUser).Methods("POST")
	r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")
}
*/

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user User
	db.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("User deleted successfully")
}

// ... other functions remain the same
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

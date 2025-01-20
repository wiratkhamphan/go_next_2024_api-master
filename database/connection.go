package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver for sql.DB
	"github.com/jmoiron/sqlx"
)

var DB *sql.DB

// DatabaseConfig initializes the database connection using sql.DB
func DatabaseConfig() {
	dbUser := "root"
	dbPass := ""
	dbName := "go_workshop_dmeo"
	dsn := dbUser + ":" + dbPass + "@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Create database connection using sql.Open
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Check if the connection is successful
	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	// Log success
	log.Println("Database connected successfully")

	// If you need to manually migrate or create tables (without AutoMigrate)
	// createTableQuery := `
	// 	CREATE TABLE IF NOT EXISTS users (
	// 		id INT AUTO_INCREMENT PRIMARY KEY,
	// 		name VARCHAR(100),
	// 		email VARCHAR(100) UNIQUE,
	// 		password VARCHAR(255)
	// 	);
	// `

	// _, err = DB.Exec(createTableQuery)
	// if err != nil {
	// 	log.Fatalf("Failed to create users table: %v", err)
	// }

	log.Println("Table created or already exists")
}

// DatabaseConfig initializes and returns a database connection
func DatabaseConfig1() (*sqlx.DB, error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "go_workshop_dmeo"

	// Create database connection
	db, err := sqlx.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return db, nil
}

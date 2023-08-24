package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	conn *sql.DB
}

var instance *Database

const maxRetries = 3
const retryInterval = 2 * time.Second

func GetInstance() (*Database, error) {
	if instance != nil {
		return instance, nil
	}

	envVars, err := LoadEnvVariables()
	if err != nil {
		return nil, fmt.Errorf("error loading environment variables: %w", err)
	}

	connectionString := createConnectionString(envVars)

	for i := 0; i < maxRetries; i++ {
		db, err := connectToDatabase(connectionString)
		if err == nil {
			instance = &Database{conn: db}
			return instance, nil
		}
		time.Sleep(retryInterval)
	}

	return nil, fmt.Errorf("failed to establish a connection after %d retries", maxRetries)
}

func createConnectionString(envVars map[string]string) string {
	user := envVars["DB_USER"]
	password := envVars["DB_PASS"]
	dbName := envVars["DB_NAME"]

	return fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s", user, dbName, password)
}

func connectToDatabase(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func (db *Database) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.conn.Query(query, args...)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, err
	}
	return rows, nil
}

func (db *Database) Execute(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.conn.Exec(query, args...)
	if err != nil {
		log.Printf("Error executing non-query statement: %v", err)
		return nil, err
	}
	return result, nil
}

func (db *Database) PrepareStatement(query string) (*sql.Stmt, error) {
	stmt, err := db.conn.Prepare(query)
	if err != nil {
		log.Printf("Error preparing statement: %v", err)
		return nil, err
	}
	return stmt, nil
}

func (db *Database) BeginTransaction() (*sql.Tx, error) {
	tx, err := db.conn.Begin()
	if err != nil {
		log.Printf("Error starting transaction: %v", err)
		return nil, err
	}
	return tx, nil
}

func (db *Database) Close() {
	if err := db.conn.Close(); err != nil {
		log.Printf("Error closing the database connection: %v", err)
	}
}

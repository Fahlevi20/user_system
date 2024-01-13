// postgress_handler.go
package Postgress

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

// InitDatabase melakukan inisialisasi koneksi ke database PostgreSQL
func InitDatabase() *pgxpool.Pool {
	connString := "postgres://username:password@localhost:5432/yourdatabase"

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatalf("Failed to parse connection string: %v", err)
	}

	// Atur konfigurasi pool database sesuai kebutuhan
	poolConfig.MaxConns = 5
	poolConfig.MinConns = 1

	// Buat instance pool database
	db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Connected to the database")

	return db
}

// GetDatabasePool mengembalikan instance pool database yang sudah diinisialisasi
func GetDatabasePool() *pgxpool.Pool {
	if db == nil {
		log.Fatal("Database pool is not initialized. Call InitDatabase first.")
	}
	return db
}

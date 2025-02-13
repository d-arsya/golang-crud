package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Error loading .env file")
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	dbpool, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Println(err)
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	DB = dbpool
	fmt.Println("Connected to PostgreSQL!")
}

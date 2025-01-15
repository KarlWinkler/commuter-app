package db

import (
	"fmt"
	"os"

	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func New() *sql.DB {
	envFile, _ := godotenv.Read(".env")

	database_info := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		envFile["DATABASE_HOST"],
		envFile["DATABASE_PORT"],
		envFile["DATABASE_USERNAME"],
		envFile["DATABASE_PASSWORD"],
		envFile["DATABASE_NAME"],
	)

	db, err := sql.Open("postgres", database_info)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	
	
	return db
}
package models

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kooparse/ameyokocho/utils"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func ConnectDB() {
	utils.ErrorCheck(godotenv.Load())

	config := fmt.Sprintf("port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := sql.Open("postgres", config)
	utils.ErrorCheck(err)

	DB = db
	log.Println("DB connected")
}

func SetupDB() {
	_, err := DB.Exec(`
   CREATE TABLE IF NOT EXISTS
    books(
      id UUID PRIMARY KEY,
      title TEXT,
      description TEXT,
      isbn_10 TEXT,
      isbn_13 TEXT,
      language TEXT,
      published_date TEXT
    );
  `)
	utils.ErrorCheck(err)

	_, err = DB.Exec(`
   CREATE TABLE IF NOT EXISTS
    authors(
      id UUID PRIMARY KEY,
      name TEXT
    );
  `)
	utils.ErrorCheck(err)
}

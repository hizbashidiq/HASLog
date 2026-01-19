package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hizbashidiq/HASLog/internal/api"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	DB_URL := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s",
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME"))	

	db, err := sql.Open("pgx", DB_URL)
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()

	timeout := 2 * time.Second

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	api.Setup(db, timeout, jwtSecret)

	log.Println("Starting web server...")
	http.ListenAndServe(os.Getenv("SERVER_ADDRESS"), nil)
}
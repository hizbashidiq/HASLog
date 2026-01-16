package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/hizbashidiq/HASLog/internal/api"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", DB_URL)
	if err!=nil{
		log.Fatal(err)
	}
	defer db.Close()

	timeout := 2 * time.Second

	api.Setup(db, timeout)

	log.Println("Starting web server...")
	http.ListenAndServe(":8080", nil)
}
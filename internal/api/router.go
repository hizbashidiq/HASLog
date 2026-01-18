package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/hizbashidiq/HASLog/internal/api/handler"
	"github.com/hizbashidiq/HASLog/internal/repository"
	"github.com/hizbashidiq/HASLog/internal/usecase"
)



func Setup(db *sql.DB, timeout time.Duration){

	// public APIs
	NewRegistrationRouter(db, timeout)

	// private APIs
	NewLoginRouter(db, timeout)
}

func NewRegistrationRouter(db *sql.DB, timeout time.Duration){
	ur := repository.NewUserRepository(db)
	ru := usecase.NewRegistrationUsecase(ur, timeout)
	rh := handler.NewRegistrationHandler(ru)

	http.HandleFunc("/registration", rh.Register)
}

func NewLoginRouter(db *sql.DB, timeout time.Duration){
	ur := repository.NewUserRepository(db)
	lu := usecase.NewLoginUsecase(ur, timeout)
	lh := handler.NewLoginHandler(lu)

	http.HandleFunc("/login", lh.Login)
}
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
	NewRegistrationRouter(db, timeout)
}

func NewRegistrationRouter(db *sql.DB, timeout time.Duration){
	ur := repository.NewUserRepository(db)
	ru := usecase.NewRegistrationUsecase(ur, timeout)
	rh := handler.NewRegistrationHandler(ru)

	http.HandleFunc("/registration", rh.Register)
}
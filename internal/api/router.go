package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/hizbashidiq/HASLog/internal/api/handler"
	"github.com/hizbashidiq/HASLog/internal/api/middleware"
	"github.com/hizbashidiq/HASLog/internal/repository"
	"github.com/hizbashidiq/HASLog/internal/usecase"
)



func Setup(db *sql.DB, timeout time.Duration, jwtSecret []byte){

	// public APIs
	NewRegistrationRouter(db, timeout)
	NewLoginRouter(db, timeout, jwtSecret)

	// private APIs
	NewProfileRouter(db, timeout, jwtSecret)
}

func NewRegistrationRouter(db *sql.DB, timeout time.Duration){
	ur := repository.NewUserRepository(db)
	ru := usecase.NewRegistrationUsecase(ur, timeout)
	rh := handler.NewRegistrationHandler(ru)

	http.HandleFunc("/registration", rh.Register)
}

func NewLoginRouter(db *sql.DB, timeout time.Duration, jwtSecret []byte){
	ur := repository.NewUserRepository(db)
	lu := usecase.NewLoginUsecase(ur, timeout, jwtSecret)
	lh := handler.NewLoginHandler(lu)

	http.HandleFunc("/login", lh.Login)
}

func NewProfileRouter(db *sql.DB, timeout time.Duration, jwtSecret []byte){
	ur := repository.NewUserRepository(db)
	pu := usecase.NewProfileUsecase(ur, timeout)
	ph := handler.NewProfileHandler(pu)
	jm := middleware.NewJWTMiddleware(jwtSecret)

	http.Handle("/profile", jm.JwtAuthMiddleware(http.HandlerFunc(ph.Profile)))
}
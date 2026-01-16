package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/hizbashidiq/HASLog/internal/domain"
	"github.com/hizbashidiq/HASLog/internal/usecase"
)

type RegistrationHandler struct {
	ru *usecase.RegistrationUsecase
}

func NewRegistrationHandler(ru *usecase.RegistrationUsecase) *RegistrationHandler{
	return &RegistrationHandler{
		ru: ru,
	}
}

func (rh *RegistrationHandler) Register(w http.ResponseWriter, r *http.Request){
	var req domain.RegistrationRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err!=nil{
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := rh.ru.Register(r.Context(), domain.RegistrationRequest{
		Username: req.Username,
		Email: req.Email,
		Password: req.Password,
	})

	if err!=nil{
		switch{
		case errors.Is(err, usecase.ErrEmailAlreadyExists):
			http.Error(w, err.Error(), http.StatusConflict)
		case errors.Is(err, usecase.ErrUsernameAlreadyExists):
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}
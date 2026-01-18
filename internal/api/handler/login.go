package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hizbashidiq/HASLog/internal/domain"
	"github.com/hizbashidiq/HASLog/internal/usecase"
)

type LoginHandler struct{
	lu *usecase.LoginUsecase
}

func NewLoginHandler(lu *usecase.LoginUsecase) *LoginHandler{
	return &LoginHandler{
		lu: lu,
	}
}

func (lh *LoginHandler)Login(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req domain.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err!=nil{
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	AccessToken, err := lh.lu.Login(r.Context(), domain.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})

	if err!=nil{
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(domain.LoginResponse{
		AccessToken: AccessToken,
	})
	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
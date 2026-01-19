package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hizbashidiq/HASLog/internal/contextkeys"
	"github.com/hizbashidiq/HASLog/internal/usecase"
)

type ProfileHandler struct{
	pu *usecase.ProfileUsecase
}

func NewProfileHandler(pu *usecase.ProfileUsecase)*ProfileHandler{
	return &ProfileHandler{
		pu:pu,
	}
}

func (ph *ProfileHandler)Profile(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.Context()

	userID,ok := ctx.Value(contextkeys.UserID).(int64)
	if !ok{
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	res,err := ph.pu.GetProfile(ctx, userID)
	if err!=nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(res)

	if err!=nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
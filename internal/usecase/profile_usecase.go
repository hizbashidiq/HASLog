package usecase

import (
	"context"
	"time"

	"github.com/hizbashidiq/HASLog/internal/domain"
	"github.com/hizbashidiq/HASLog/internal/repository"
)

type ProfileUsecase struct {
	ur *repository.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(ur *repository.UserRepository, contextTimeout time.Duration)*ProfileUsecase{
	return &ProfileUsecase{
		ur: ur,
		contextTimeout: contextTimeout,
	}
}

func (pu *ProfileUsecase)GetProfile(ctx context.Context, userID int64) (domain.ProfileResponse, error){
	ctx, cancel := context.WithTimeout(ctx, pu.contextTimeout)
	defer cancel()

	user, err := pu.ur.FindByID(ctx, userID)
	if err!=nil{
		return domain.ProfileResponse{}, err
	}

	return domain.ProfileResponse{
		Username: user.Username,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}
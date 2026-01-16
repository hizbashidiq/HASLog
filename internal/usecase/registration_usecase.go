package usecase

import (
	"context"
	"errors"
	"time"
	"fmt"

	"github.com/hizbashidiq/HASLog/internal/domain"
	"github.com/hizbashidiq/HASLog/internal/repository"
	"github.com/jackc/pgx/v5/pgconn"
	"golang.org/x/crypto/bcrypt"
)

var(
	ErrEmailAlreadyExists = errors.New("email already exist")
	ErrUsernameAlreadyExists = errors.New("username already exist")
	ErrUnknown = errors.New("unknown error")
)

type RegistrationUsecase struct{
	ur *repository.UserRepository
	contextTimeout time.Duration
}

func NewRegistrationUsecase(ur *repository.UserRepository, contextTimeout time.Duration) *RegistrationUsecase{
	return &RegistrationUsecase{
		ur: ur,
		contextTimeout: contextTimeout,
	}
}

func (ru *RegistrationUsecase) Register(ctx context.Context, input domain.RegistrationRequest) error{

	ctx, cancel := context.WithTimeout(ctx, ru.contextTimeout)
	defer cancel()

	// hash password
	hashedPassword, err := ru.hashPassword(input.Password)
	if err!=nil{
		return err
	}
	
	user := domain.User{
		Username: input.Username,
		Email: input.Email,
		PasswordHash: hashedPassword,
	}

	err = ru.ur.CreateUser(ctx, user)
	if err!=nil{
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505"{
			switch pgErr.ConstraintName{
			case "users_username_key":
				return ErrUsernameAlreadyExists
			case "users_email_key":
				return ErrEmailAlreadyExists
			}
		}
		return fmt.Errorf("%w: %v", ErrUnknown, err)
	}
	return nil
}

func (ru *RegistrationUsecase)hashPassword(password string) (string, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	
	return string(hashedPassword), err
}

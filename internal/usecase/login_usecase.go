package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hizbashidiq/HASLog/internal/domain"
	"github.com/hizbashidiq/HASLog/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	ur *repository.UserRepository
	contextTimeout time.Duration
	jwtSecret []byte
}

func NewLoginUsecase(ur *repository.UserRepository, contextTimeout time.Duration) *LoginUsecase {
	return &LoginUsecase{
		ur: ur,
		contextTimeout: contextTimeout,
	}
}

func (lu *LoginUsecase)generateAccessToken(userID int64, secretKey []byte)(string, error){
	claims := jwt.RegisteredClaims{
		Issuer: "haslog",
		Subject: strconv.FormatInt(userID, 10),
		IssuedAt: jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken,err := token.SignedString(secretKey)
	if err!=nil{
		return "", err
	}

	return signedToken, err
}



func (lu *LoginUsecase)Login(ctx context.Context, lr domain.LoginRequest) (string,error){
	ctx, cancel := context.WithTimeout(ctx, lu.contextTimeout)
	defer cancel()
	
	user, err:=lu.ur.FindByUsername(ctx, lr.Username)
	if err!=nil{
		return "",ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(lr.Password))
	if err!=nil{
		return "",ErrInvalidCredentials
	}

	accessToken, err := lu.generateAccessToken(user.ID, lu.jwtSecret)
	if err!=nil{
		return "",err
	}

	return accessToken,nil
}
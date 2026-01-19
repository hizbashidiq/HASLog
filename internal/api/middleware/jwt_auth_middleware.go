package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hizbashidiq/HASLog/internal/contextkeys"
)

type JWTMiddleware struct{
	secret []byte
}

func NewJWTMiddleware(secret []byte)*JWTMiddleware{
	return &JWTMiddleware{
		secret: secret,
	}
}

func (jm *JWTMiddleware)JwtAuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer "){
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error){
			if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
				return nil, fmt.Errorf("invalid signing method")
			}
			return jm.secret, nil
		})
		if err!=nil || !token.Valid{
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		userID, err := strconv.ParseInt(claims.Subject,10,64)
		if err!=nil{
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), contextkeys.UserID, userID)

		next.ServeHTTP(w,r.WithContext(ctx))
	})
}
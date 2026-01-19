package repository

import (
	"database/sql"

	"github.com/hizbashidiq/HASLog/internal/domain"
	_ "github.com/jackc/pgx/v5"

	"context"
)


type UserRepository struct{
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository{
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user domain.User) error{
	query := `
		INSERT INTO users(username, email, password_hash)
		VALUES($1, $2, $3)
	`
	_, err := ur.db.ExecContext(ctx, query, user.Username, user.Email, user.PasswordHash)
	return err
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error){
	var user domain.User

	query := `
		SELECT id, username, email, password_hash, created_at
		FROM users
		WHERE email=$1
	`
	err := ur.db.QueryRowContext(ctx, query, email).
		Scan(
			&user.ID, 
			&user.Username, 
			&user.Email, 
			&user.PasswordHash, 
			&user.CreatedAt,
		)

	return user, err
}

func (ur *UserRepository) FindByUsername(ctx context.Context, username string)(domain.User, error){
	var user domain.User
	query := `
		SELECT id, username, email, password_hash, created_at
		FROM users
		WHERE username=$1
	`
	err := ur.db.QueryRowContext(ctx, query, username).
		Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
		)
		
	return user, err
}

func (ur *UserRepository)FindByID(ctx context.Context, userID int64)(domain.User, error){
	var user domain.User
	query := `
		SELECT id, username, email, password_hash, created_at
		FROM users
		WHERE id=$1
	`
	err := ur.db.QueryRowContext(ctx, query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	return user, err
}
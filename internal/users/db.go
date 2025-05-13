package users

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type DBStore struct {
	db *sqlx.DB
}

func NewDBStore(db *sqlx.DB) *DBStore {
	return &DBStore{db: db}
}

func (s *DBStore) CreateUser(ctx context.Context, username, password string) (User, error) {
	if username == "" || password == "" {
		return User{}, errors.New("username and password cannot be empty")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	u := User{
		Username:     username,
		PasswordHash: string(hash),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	query := `INSERT INTO users (username, password_hash, created_at, updated_at) 
              VALUES (:username, :password_hash, :created_at, :updated_at) 
              RETURNING *`
	rows, err := s.db.NamedQueryContext(ctx, query, u)
	if err != nil {
		return User{}, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Failed to close rows in CreateUser: %v", err)
		}
	}()
	if rows.Next() {
		if err := rows.StructScan(&u); err != nil {
			return User{}, err
		}
	} else {
		return User{}, errors.New("failed to retrieve inserted user")
	}
	return u, nil
}

func (s *DBStore) GetUserByUsername(ctx context.Context, username string) (User, error) {
	var u User
	err := s.db.GetContext(ctx, &u, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return User{}, errors.New("user not found")
		}
		return User{}, err
	}
	return u, nil
}

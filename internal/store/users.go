package store

import (
	"context"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  password  `json:"-"`
	CreatedAt string   `json:"created_at"`
	
}

type password struct {
	text *string
	hash []byte
}

func (p *password) Set(text string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	p.text = &text
	p.hash = hash

	return nil
}

func (p *password) Compare(text string) error {
	return bcrypt.CompareHashAndPassword(p.hash, []byte(text))
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO posts (username,  password, email) VALUES (&1,&2, &3, &4) RETURNING id, created_at, updated_at
	`

	// ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	// defer cancel()

	


	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		
			return err
		}
	

	return nil
}

func (s *UsersStore) GetByID(ctx context.Context, userID int64)(*User, error){
	query := `
		SELECT users.id, username, email, password, created_at,
		FROM users
		
		WHERE users.id = $1 
	`

	// ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	// defer cancel()

	user := &User{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}

func (s *UsersStore) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, username, email, password, created_at FROM users
		WHERE email = $1 
	`

	// ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	// defer cancel()

	user := &User{}
	err := s.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password.hash,
		&user.CreatedAt,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil
}
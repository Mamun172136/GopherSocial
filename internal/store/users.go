package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  string `json:"-"`
	CreatedAt string   `json:"created_at"`
	
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
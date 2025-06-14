package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"createdAt"`
}
type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (userName,email,password) VALUES($1,$2,$3) RETURNING id,createdAt`

	err := s.db.QueryRowContext(ctx, query, user.UserName, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

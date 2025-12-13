package data

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Version   int       `json:"-"`
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Get(id int64) (*User, error) {
	query := `
	SELECT id, created_at, name, email, version
	FROM users
	WHERE id = $1`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Version,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

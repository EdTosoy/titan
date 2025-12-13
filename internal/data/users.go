package data

import (
	"context"
	"database/sql"
	"errors"
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

func (m UserModel) Insert(user *User) error {
	query := `
	INSERT INTO users (name, email)
	VALUES ($1, $2)
	RETURNING id, created_at, version`

	args := []any{user.Name, user.Email}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Version,
	)

}

func (m UserModel) Update(user *User) error {
	query := `
	UPDATE users 
	SET name = $1, email = $2, version = version + 1
	WHERE id = $3 AND version = $4
	RETURNING version`

	args := []any{user.Name, user.Email, user.ID, user.Version}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.Version,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("edit conflict")
		}
		return err
	}

	return nil
}

func (m UserModel) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// Check if any row was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

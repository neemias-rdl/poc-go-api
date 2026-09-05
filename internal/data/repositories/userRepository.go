package repositories

import (
	"api/internal/data"
	"api/internal/data/entities"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type UserRepository struct {
	db data.DB
}

func NewUserRepository(db data.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, username, email, hashed_password
		FROM users
		WHERE id = $1
	`, id)

	if err != nil {
		return nil, fmt.Errorf("Error querying user by id: " + err.Error())
	}

	defer rows.Close()

	var user entities.User

	if err := rows.Scan(
		&user.UUID,
		&user.Username,
		&user.Email,
		&user.HashedPassword,
	); err != nil {
		return nil, fmt.Errorf("Error parsing user: " + err.Error())
	}

	return &user, nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*entities.User, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, username, email, hashed_password
		FROM users
		WHERE username = $1
	`, username)

	if err != nil {
		return nil, fmt.Errorf("Error querying user by username: " + err.Error())
	}

	defer rows.Close()

	var user entities.User

	if err := rows.Scan(
		&user.UUID,
		&user.Username,
		&user.Email,
		&user.HashedPassword,
	); err != nil {
		return nil, fmt.Errorf("Error parsing user: " + err.Error())
	}

	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user entities.User) (*entities.User, error) {
	err := r.db.QueryRow(ctx, `
        INSERT INTO users (id, username, email, hashed_password)
        VALUES ($1, $2, $3, $4)
        RETURNING id, username, email, hashed_password
    `,
		user.UUID,
		user.Username,
		user.Email,
		user.HashedPassword,
	).Scan(
		&user.UUID,
		&user.Username,
		&user.Email,
		&user.HashedPassword,
	)

	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}

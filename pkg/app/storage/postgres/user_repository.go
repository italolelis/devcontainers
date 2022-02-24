package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/italolelis/devcontainers/pkg/app/user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*user.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("could not query users table: %w", err)
	}

	defer rows.Close()

	var users []*user.User
	for rows.Next() {
		u := user.User{}
		if err := rows.Scan(&u.ID, &u.Username, &u.CreatedAt); err != nil {
			return nil, fmt.Errorf("could not scan users table: %w", err)
		}
		users = append(users, &u)
	}

	return users, nil
}

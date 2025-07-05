package repositories

import (
	"authjwt/internal/models"
	"database/sql"
	"fmt"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *models.User) (*models.User, error) {
	err := r.db.QueryRow(
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Name, user.Email, user.Password,
	).Scan(&user.ID)

	if err != nil {
		return nil, fmt.Errorf("s.repo.Create: %v", err)
	}
	return user, nil
}

func (r *UserRepo) GetByEmail(user *models.User) (*models.User, error) {
	res, err := r.db.QueryRow(
		"SELECT id, name, email, password FROM users WHERE email = $1",
		user.Email,
	)

	if err != nil {
		return nil, fmt.Errorf("s.repo.Create: %v", err)
	}
	return user, nil
}

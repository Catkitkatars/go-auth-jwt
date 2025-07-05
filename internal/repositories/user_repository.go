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

func (r *UserRepo) SaveTokenByUser(user *models.User, token string) (bool, error) {
	_, err := r.db.Exec(
		"UPDATE users SET refresh_token = $1 WHERE id = $2",
		token, user.ID,
	)

	if err != nil {
		return false, fmt.Errorf("s.repo.SaveTokenByUser: %v", err)
	}
	return true, nil
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	userRepo := &models.User{}

	err := r.db.QueryRow(
		"SELECT id, name, email, password FROM users WHERE email = $1",
		email,
	).Scan(&userRepo.ID, &userRepo.Name, &userRepo.Email, &userRepo.Password)

	if err != nil {
		return nil, fmt.Errorf("s.repo.GetByEmail: %v", err)
	}
	return userRepo, nil
}

package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type IUserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}

type UserRepository struct {
	users []*models.User
}

func NewUserRepository() IUserRepository {
	p1, _ := bcrypt.GenerateFromPassword([]byte("11111111"), bcrypt.DefaultCost)
	p2, _ := bcrypt.GenerateFromPassword([]byte("22222222"), bcrypt.DefaultCost)

	users := []*models.User{
		{
			Id:           1,
			Email:        "alex@example.com",
			Name:         "Alex",
			PasswordHash: string(p1),
		},
		{
			Id:           2,
			Email:        "mary@example.com",
			Name:         "Mary",
			PasswordHash: string(p2),
		},
	}

	return &UserRepository{users: users}
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	for _, user := range r.users {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

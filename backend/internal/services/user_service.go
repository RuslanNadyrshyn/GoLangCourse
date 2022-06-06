package services

import (
	"Delivery/backend/internal/repositories"
	"Delivery/backend/internal/repositories/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store *repositories.Store
}

func NewUserService(store *repositories.Store) *UserService {
	return &UserService{
		store: store,
	}
}

func (r *UserService) Insert(u *models.User) (id int64, err error) {
	if u.Email == "" {
		id, err = r.store.UserRepo.InsertName(u)
		if err != nil {
			return 0, err
		}
		return id, nil
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	u.PasswordHash = string(PasswordHash)

	id, err = r.store.UserRepo.Insert(u)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return id, nil
}

func (r *UserService) GetByEmail(email string) (user *models.User, err error) {
	user, err = r.store.UserRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserService) GetById(id int64) (user *models.User, err error) {
	user, err = r.store.UserRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserService) Update(u *models.User) (id int64, err error) {
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	u.PasswordHash = string(PasswordHash)

	if r.store.UserRepo.TX != nil {
		id, err = r.store.UserRepo.Update(u)
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	id, err = r.store.UserRepo.Update(u)
	if err != nil {
		return 0, err
	}

	return id, nil
}

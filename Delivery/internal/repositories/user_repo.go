package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IUserRepository interface {
	Insert(u *models.User) (int64, error)
	GetByEmail(email string) (models.User, error)
	GetById(id int64) (models.User, error)
	Update(u models.User, email string) (int64, error)
}

type UserRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewUserRepository(conn *sql.DB) IUserRepository {
	return &UserRepository{
		DB: conn,
	}
}

func (r *UserRepository) Insert(u *models.User) (id int64, err error) {
	if u.Email == "" {
		result, err := r.DB.Exec("INSERT users(name) VALUES(?)", u.Name)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
		id, err = result.LastInsertId()
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
		return id, nil
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	u.PasswordHash = string(PasswordHash)

	result, err := r.DB.Exec("INSERT users(name, email, password_hash) VALUES(?, ?, ?)", u.Name, u.Email, u.PasswordHash)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) GetByEmail(email string) (user models.User, err error) {
	err = r.DB.QueryRow("SELECT id, email, password_hash, name FROM users WHERE email = ?", email).
		Scan(&user.Id, &user.Email, &user.PasswordHash, &user.Name)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

func (r *UserRepository) GetById(id int64) (user models.User, err error) {
	err = r.DB.QueryRow("SELECT id, ifnull(email, ''), name FROM users WHERE id = ?", id).Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func (r *UserRepository) Update(u models.User, email string) (id int64, err error) {
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	u.PasswordHash = string(PasswordHash)

	if r.TX != nil {
		_, err = r.TX.Exec("UPDATE users SET name=?, email=?, password_hash=? WHERE email = ?", u.Name, u.Email, u.PasswordHash, email)
		if err != nil {
			_ = r.TX.Rollback()
		}
		err = r.TX.Commit()
		if err != nil {
			_ = r.TX.Rollback()
		}
		return id, nil
	}

	_, err = r.DB.Exec("UPDATE users SET name=?, email=?, password_hash=? WHERE email = ?", u.Name, u.Email, u.PasswordHash, email)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}

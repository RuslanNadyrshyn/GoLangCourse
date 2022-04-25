package database

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type UserDBRepository struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewUserRepository(conn *sql.DB) UserDBRepository {
	return UserDBRepository{
		DB: conn,
	}
}

func (r UserDBRepository) Insert(u *models.User) (int, error) {
	var id int64

	if u.Email == "" {
		result, err := r.DB.Exec("INSERT users(name) VALUES(?)", u.Name)
		if err != nil {
			fmt.Println(err)
			return int(id), err
		}
		id, err = result.LastInsertId()
		if err != nil {
			fmt.Println(err)
			return int(id), err
		}
		return int(id), err
	}

	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	u.PasswordHash = string(PasswordHash)
	if r.TX != nil {
		err := r.TX.QueryRow("INSERT users(name, email, password_hash) VALUES(?, ?, ?) RETURNING id", u.Name, u.Email, u.PasswordHash).Scan(&id)
		if err != nil {
			_ = r.TX.Rollback()
		}
		err = r.TX.Commit()
		if err != nil {
			_ = r.TX.Rollback()
		}
		return int(id), err
	}

	result, err := r.DB.Exec("INSERT users(name, email, password_hash) VALUES(?, ?, ?)", u.Name, u.Email, u.PasswordHash)
	if err != nil {
		fmt.Println(err)
		return int(id), err
	}
	id, err = result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return int(id), err
	}

	return int(id), err
}

func (r UserDBRepository) GetByEmail(email string) (models.User, error) {
	var user models.User

	err := r.DB.QueryRow("SELECT id, email, password_hash, name FROM users WHERE email = ?", email).
		Scan(&user.Id, &user.Email, &user.PasswordHash, &user.Name)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r UserDBRepository) GetById(id int) (models.User, error) {
	var user models.User

	err := r.DB.QueryRow("SELECT id, email, name FROM users WHERE id = ?", id).Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r UserDBRepository) Update(u models.User, email string) (int, error) {
	var id int
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	u.PasswordHash = string(PasswordHash)

	if r.TX != nil {
		err := r.TX.QueryRow("UPDATE users SET name=?, email=?, password_hash=? WHERE email = ? RETURNING id", u.Name, u.Email, u.PasswordHash, email).Scan(&id)
		if err != nil {
			_ = r.TX.Rollback()
		}
		err = r.TX.Commit()
		if err != nil {
			_ = r.TX.Rollback()
		}
		return id, err
	}

	_, err = r.DB.Exec("UPDATE users SET name=?, email=?, password_hash=? WHERE email = ?", u.Name, u.Email, u.PasswordHash, email)

	if err != nil {
		fmt.Println(err)
		return id, err
	}

	return id, err
}

func (r UserDBRepository) Delete(email string) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE email = ?", email)
	if err != nil {
		return err
	}

	return nil
}

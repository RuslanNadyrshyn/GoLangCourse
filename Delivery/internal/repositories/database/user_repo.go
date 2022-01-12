package database

import (
	"awesomeProject/internal/repositories/models"
	"database/sql"
	"fmt"
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

func (r UserDBRepository) Insert(u models.User) (int, error) {
	var id int

	if r.TX != nil {
		err := r.TX.QueryRow("INSERT users(name, email, password_hash) VALUES(?, ?, ?) RETURNING id", u.Name, u.Email, u.PasswordHash).Scan(&id)
		if err != nil {
			_ = r.TX.Rollback()
		}
		err = r.TX.Commit()
		if err != nil {
			_ = r.TX.Rollback()
		}
		return id, err
	}
	_, err := r.DB.Exec("INSERT users(name, email, password_hash) VALUES(?, ?, ?)", u.Name, u.Email, u.PasswordHash)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, err
}

func (r UserDBRepository) GetByEmail(email string) (models.User, error) {
	var user models.User

	err := r.DB.QueryRow("SELECT id, email, name FROM users WHERE email = ?", email).Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r UserDBRepository) Delete(email string) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE email = ?", email)
	if err != nil {
		return err
	}

	return nil
}

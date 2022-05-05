package repositories

import (
	"Delivery/Delivery/internal/repositories/models"
	"database/sql"
)

type UserRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewUserRepo(conn *sql.DB) *UserRepo {
	return &UserRepo{
		DB: conn,
	}
}

func (r *UserRepo) Insert(u *models.User) (id int64, err error) {
	if r.TX != nil {
		result, err := r.TX.Exec("INSERT users(name, email, password_hash) VALUES(?, ?, ?)",
			u.Name, u.Email, u.PasswordHash)
		if err != nil {
			return 0, err
		}
		id, err = result.LastInsertId()
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	result, err := r.DB.Exec("INSERT users(name, email, password_hash) VALUES(?, ?, ?)",
		u.Name, u.Email, u.PasswordHash)
	if err != nil {
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepo) InsertName(u *models.User) (id int64, err error) {
	if r.TX != nil {
		result, err := r.DB.Exec("INSERT users(name) VALUES(?)", u.Name)
		if err != nil {
			return 0, err
		}
		id, err = result.LastInsertId()
		if err != nil {
			return 0, err
		}
		return id, nil
	}
	result, err := r.DB.Exec("INSERT users(name) VALUES(?)", u.Name)
	if err != nil {
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, email, password_hash, name FROM users WHERE email = ?", email).
		Scan(&user.Id, &user.Email, &user.PasswordHash, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepo) GetById(id int64) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT id, ifnull(email, ''), name FROM users WHERE id = ?", id).
		Scan(&user.Id, &user.Email, &user.Name)
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func (r *UserRepo) Update(u *models.User) (id int64, err error) {
	_, err = r.DB.Exec("UPDATE users SET name=?, email=?, password_hash=? WHERE id = ?",
		&u.Name, &u.Email, &u.PasswordHash, &u.Id)
	if err != nil {
		return 0, err
	}
	return u.Id, nil
}

func (r *UserRepo) BeginTx() error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	r.TX = tx
	return nil
}

func (r *UserRepo) CommitTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.CommitTx()
	}
	return nil
}

func (r *UserRepo) RollbackTx() error {
	defer func() {
		r.TX = nil
	}()
	if r.TX != nil {
		return r.RollbackTx()
	}
	return nil
}

func (r *UserRepo) GetTx() *sql.Tx {
	return r.TX
}

func (r *UserRepo) SetTx(tx *sql.Tx) {
	r.TX = tx
}

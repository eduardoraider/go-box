package repositories

import (
	"database/sql"
	"github.com/eduardoraider/go-box/internal/users"
	"time"
)

type UserReadRepository interface {
	Login(string) *sql.Row
	Get(int64) *sql.Row
	List() (*sql.Rows, error)
}

type UserWriteRepository interface {
	Create(*users.User) (int64, error)
	Update(int64, *users.User) error
	Delete(int64) error
}

type UserReadWriteRepository interface {
	UserReadRepository
	UserWriteRepository
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) Create(u *users.User) (id int64, err error) {
	stmt := `INSERT INTO users(name, login, password, modified_at) values($1, $2, $3, $4) RETURNING id`
	err = ur.db.QueryRow(stmt, u.Name, u.Login, u.Password, u.ModifiedAt).Scan(&id)
	if err != nil {
		return -1, err
	}

	return
}

func (ur *UserRepository) Update(id int64, u *users.User) error {
	u.ModifiedAt = time.Now()

	stmt := `UPDATE users SET name=$1, modified_at=$2, last_login=$3 WHERE id=$4`
	_, err := ur.db.Exec(stmt, u.Name, u.ModifiedAt, u.LastLogin, id)
	return err
}

func (ur *UserRepository) Delete(id int64) error {
	stmt := `UPDATE users SET modified_at=$1, deleted=true WHERE id=$2`
	_, err := ur.db.Exec(stmt, time.Now(), id)
	return err
}

func (ur *UserRepository) Get(id int64) *sql.Row {
	stmt := `SELECT * FROM users WHERE id=$1`
	return ur.db.QueryRow(stmt, id)
}

func (ur *UserRepository) List() (*sql.Rows, error) {
	stmt := `SELECT * FROM users WHERE deleted=false`
	return ur.db.Query(stmt)
}

func (ur *UserRepository) Login(login string) *sql.Row {
	stmt := `SELECT * FROM users WHERE login=$1`
	return ur.db.QueryRow(stmt, login)
}

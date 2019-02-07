package mysql

import (
	"database/sql"

	"github.com/tayusa/notugly_backend/app/domain"
)

type userRepository struct {
	db *sql.DB
}

func (u *userRepository) FindById(uid string) (domain.User, error) {
	var user domain.User
	var sex string
	err := u.db.QueryRow(
		"SELECT id, name, sex, age FROM users WHERE id = ?", uid).
		Scan(&user.Id, &user.Name, &sex, &user.Age)

	switch sex {
	case "\x00":
		user.Sex = 0
	case "\x01":
		user.Sex = 1
	}

	return user, err
}

func (u *userRepository) Store(user domain.User) (err error) {
	_, err = u.db.Exec(
		"INSERT INTO users(id, name, sex, age) VALUES(?, ?, ?, ?)",
		user.Id, user.Name, user.Sex, user.Age)
	return
}

func (u *userRepository) Update(user domain.User) (err error) {
	_, err = u.db.Exec(
		"UPDATE users SET name = ?, sex = ?, age = ? WHERE id = ?",
		user.Name, user.Sex, user.Age, user.Id)
	return
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

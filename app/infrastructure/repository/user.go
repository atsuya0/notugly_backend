package repository

import (
	"database/sql"

	"github.com/tayusa/notugly_backend/app/domain"
)

type userRepository struct {
	conn *sql.DB
}

func (u *userRepository) FindById(uid string) (domain.User, error) {
	var user domain.User
	err := u.conn.QueryRow(
		"SELECT id, name, sex, age FROM users WHERE id = ?", uid).
		Scan(&user.Id, &user.Name, &user.Sex, &user.Age)

	return user, err
}

func (u *userRepository) Store(user domain.User) (err error) {
	_, err = u.conn.Exec(
		"INSERT INTO users(id, name, sex, age) VALUES(?, ?, ?, ?)",
		user.Id, user.Name, user.Sex, user.Age)
	return
}

func (u *userRepository) Update(user domain.User) (err error) {
	_, err = u.conn.Exec(
		"UPDATE users SET name = ?, sex = ?, age = ? WHERE id = ?",
		user.Name, user.Sex, user.Age, user.Id)
	return
}

func NewUserRepository(conn *sql.DB) *userRepository {
	return &userRepository{conn}
}

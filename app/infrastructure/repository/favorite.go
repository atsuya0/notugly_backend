package repository

import (
	"database/sql"

	"github.com/tayusa/notugly_backend/app/domain"
)

type favoriteRepository struct {
	conn *sql.DB
}

func (f *favoriteRepository) Store(favorite domain.Favorite) (err error) {
	_, err = f.conn.Exec(
		"INSERT INTO favorites(coordinate_id, user_id) VALUES(?, ?)",
		favorite.CoordinateId, favorite.UserId)
	return
}

func (f *favoriteRepository) Delete(favorite domain.Favorite) (err error) {
	_, err = f.conn.Exec(
		"DELETE FROM favorite WHERE coordinate_id = ? and user_id = ?",
		favorite.CoordinateId, favorite.UserId)
	return
}

func NewFavoriteRepository(conn *sql.DB) *favoriteRepository {
	return &favoriteRepository{conn}
}

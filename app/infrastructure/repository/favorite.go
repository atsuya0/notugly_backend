package repository

import (
	"database/sql"

	"github.com/tayusa/notugly_backend/app/domain"
)

type favoriteRepository struct {
	db *sql.DB
}

func (f *favoriteRepository) Store(favorite domain.Favorite) (err error) {
	_, err = f.db.Exec(
		"INSERT INTO favorites(coordinate_id, user_id) VALUES(?, ?)",
		favorite.CoordinateId, favorite.UserId)
	return
}

func (f *favoriteRepository) Delete(favorite domain.Favorite) (err error) {
	_, err = f.db.Exec(
		"DELETE FROM favorites WHERE coordinate_id = ? AND user_id = ?",
		favorite.CoordinateId, favorite.UserId)
	return
}

func NewFavoriteRepository(db *sql.DB) *favoriteRepository {
	return &favoriteRepository{db}
}

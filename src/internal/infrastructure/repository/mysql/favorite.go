package mysql

import (
	"context"
	"database/sql"

	"github.com/tayusa/notugly_backend/internal/domain"
)

type favoriteRepository struct {
	db *sql.DB
}

func (f *favoriteRepository) Store(ctx context.Context, favorite domain.Favorite) (err error) {
	_, err = f.db.ExecContext(ctx,
		"INSERT INTO favorites(coordinate_id, user_id) VALUES(?, ?)",
		favorite.CoordinateId, favorite.UserId)
	return
}

func (f *favoriteRepository) Delete(ctx context.Context, favorite domain.Favorite) (err error) {
	_, err = f.db.ExecContext(ctx,
		"DELETE FROM favorites WHERE coordinate_id = ? AND user_id = ?",
		favorite.CoordinateId, favorite.UserId)
	return
}

func NewFavoriteRepository(db *sql.DB) *favoriteRepository {
	return &favoriteRepository{db}
}

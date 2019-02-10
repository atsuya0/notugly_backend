package repository

import (
	"github.com/tayusa/notugly_backend/domain"
)

type FavoriteRepository interface {
	Store(domain.Favorite) error
	Delete(domain.Favorite) error
}

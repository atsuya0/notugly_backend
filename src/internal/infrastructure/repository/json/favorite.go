package json

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/tayusa/notugly_backend/internal/domain"
)

type favoriteRepository struct {
}

func (f *favoriteRepository) Store(_ context.Context, _ domain.Favorite) error {
	return nil
}

func (f *favoriteRepository) Delete(_ context.Context, _ domain.Favorite) error {
	return nil
}

func NewFavoriteRepository() *favoriteRepository {
	return &favoriteRepository{}
}

func GetFavorites(method int) ([]domain.Favorite, error) {
	bytes, err := ioutil.ReadFile(
		filepath.Join(jsonPath, "favorites", methods[method]+".json"))
	if err != nil {
		return []domain.Favorite{}, err
	}

	var favorites []domain.Favorite
	if err = json.Unmarshal(bytes, &favorites); err != nil {
		return []domain.Favorite{}, err
	}
	return favorites, nil
}

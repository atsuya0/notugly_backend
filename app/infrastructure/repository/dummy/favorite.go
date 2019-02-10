package dummy

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/tayusa/notugly_backend/domain"
)

type dummyFavoriteRepository struct {
}

func (f *dummyFavoriteRepository) Store(_ domain.Favorite) error {
	return nil
}

func (f *dummyFavoriteRepository) Delete(_ domain.Favorite) error {
	return nil
}

func NewDummyFavoriteRepository() *dummyFavoriteRepository {
	return &dummyFavoriteRepository{}
}

func GetFavorites(method int) ([]domain.Favorite, error) {
	bytes, err := ioutil.ReadFile(filepath.Join("testdata/json/favorites", methods[method]+".json"))
	if err != nil {
		return []domain.Favorite{}, err
	}

	var favorites []domain.Favorite
	if err = json.Unmarshal(bytes, &favorites); err != nil {
		return []domain.Favorite{}, err
	}
	return favorites, nil
}

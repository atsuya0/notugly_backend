package dummy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/tayusa/notugly_backend/domain"
)

type dummyUserRepository struct {
}

func (u *dummyUserRepository) FindById(id string) (domain.User, error) {
	users, err := GetUsers(GET)
	if err != nil {
		return domain.User{}, err
	}

	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}

	return domain.User{}, fmt.Errorf("err no rows")
}

func (u *dummyUserRepository) Store(_ domain.User) error {
	return nil
}

func (u *dummyUserRepository) Update(_ domain.User) error {
	return nil
}

func NewDummyUserRepository() *dummyUserRepository {
	return &dummyUserRepository{}
}

func GetUsers(method int) ([]domain.User, error) {
	bytes, err := ioutil.ReadFile(filepath.Join("testdata/json/users", methods[method]+".json"))
	if err != nil {
		return []domain.User{}, err
	}

	var users []domain.User
	if err = json.Unmarshal(bytes, &users); err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

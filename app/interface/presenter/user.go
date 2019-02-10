package presenter

import (
	"encoding/json"

	"github.com/tayusa/notugly_backend/domain"
)

type userPresenter struct {
}

func (u *userPresenter) ResponseUser(user domain.User) ([]byte, error) {
	output, err := json.Marshal(&user)
	if err != nil {
		return []byte{}, err
	}
	return output, nil
}

func NewUserPresenter() *userPresenter {
	return &userPresenter{}
}

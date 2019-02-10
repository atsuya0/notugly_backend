package presenter

import (
	"github.com/tayusa/notugly_backend/domain"
)

type UserPresenter interface {
	ResponseUser(domain.User) ([]byte, error)
}

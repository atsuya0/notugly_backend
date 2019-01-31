package presenter

import (
	"github.com/tayusa/notugly_backend/app/domain"
)

type UserPresenter interface {
	ResponseUser(domain.User) ([]byte, error)
}

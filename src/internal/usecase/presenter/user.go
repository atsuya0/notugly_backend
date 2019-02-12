package presenter

import (
	"github.com/tayusa/notugly_backend/internal/domain"
)

type UserPresenter interface {
	ResponseUser(domain.User) ([]byte, error)
}

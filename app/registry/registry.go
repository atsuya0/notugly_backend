package registry

import (
	"github.com/tayusa/notugly_backend/infrastructure/api/handler"
)

type Iteractor interface {
	NewAppHandler() handler.AppHandler
}

package registry

import (
	"database/sql"

	"github.com/tayusa/notugly_backend/app/infrastructure/api/handler"
	"github.com/tayusa/notugly_backend/app/infrastructure/repository"
	"github.com/tayusa/notugly_backend/app/interface/controller"
	"github.com/tayusa/notugly_backend/app/interface/presenter"
	"github.com/tayusa/notugly_backend/app/usecase/service"
)

type interactor struct {
	db *sql.DB
}

func NewInteractor(db *sql.DB) Iteractor {
	return &interactor{db}
}

type Iteractor interface {
	NewAppHandler() handler.AppHandler
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return handler.NewAppHandler(
		i.NewUserHandler(), i.NewCoordinateHandler(), i.NewFavoriteHandler())
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(
		controller.NewUserController(
			service.NewUserService(
				repository.NewUserRepository(i.db),
				presenter.NewUserPresenter())))
}

func (i *interactor) NewCoordinateHandler() handler.CoordinateHandler {
	return handler.NewCoordinateHandler(
		controller.NewCoordinateController(
			service.NewCoordinateService(
				repository.NewCoordinateRepository(i.db),
				presenter.NewCoordinatePresenter())))
}

func (i *interactor) NewFavoriteHandler() handler.FavoriteHandler {
	return handler.NewFavoriteHandler(
		controller.NewFavoriteController(
			service.NewFavoriteService(
				repository.NewFavoriteRepository(i.db))))
}

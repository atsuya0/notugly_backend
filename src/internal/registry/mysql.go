package registry

import (
	"database/sql"

	"github.com/tayusa/notugly_backend/internal/infrastructure/api/handler"
	"github.com/tayusa/notugly_backend/internal/infrastructure/repository/mysql"
	"github.com/tayusa/notugly_backend/internal/interface/controller"
	"github.com/tayusa/notugly_backend/internal/interface/presenter"
	"github.com/tayusa/notugly_backend/internal/usecase/service"
)

type interactor struct {
	db *sql.DB
}

func NewInteractor(db *sql.DB) Iteractor {
	return &interactor{db}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return handler.NewAppHandler(
		i.NewUserHandler(), i.NewCoordinateHandler(), i.NewFavoriteHandler())
}

func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(
		controller.NewUserController(
			service.NewUserService(
				mysql.NewUserRepository(i.db),
				presenter.NewUserPresenter())))
}

func (i *interactor) NewCoordinateHandler() handler.CoordinateHandler {
	return handler.NewCoordinateHandler(
		controller.NewCoordinateController(
			service.NewCoordinateService(
				"images",
				mysql.NewCoordinateRepository(i.db),
				presenter.NewCoordinatePresenter())))
}

func (i *interactor) NewFavoriteHandler() handler.FavoriteHandler {
	return handler.NewFavoriteHandler(
		controller.NewFavoriteController(
			service.NewFavoriteService(
				mysql.NewFavoriteRepository(i.db))))
}

package registry

import (
	"github.com/tayusa/notugly_backend/internal/infrastructure/api/handler"
	"github.com/tayusa/notugly_backend/internal/infrastructure/repository/json"
	"github.com/tayusa/notugly_backend/internal/interface/controller"
	"github.com/tayusa/notugly_backend/internal/interface/presenter"
	"github.com/tayusa/notugly_backend/internal/usecase/service"
)

type dummyInteractor struct {
	imagePath string
}

func NewDummyInteractor(imagePath string) Iteractor {
	return &dummyInteractor{imagePath}
}

func (i *dummyInteractor) NewAppHandler() handler.AppHandler {
	return handler.NewAppHandler(
		i.NewUserHandler(), i.NewCoordinateHandler(), i.NewFavoriteHandler())
}

func (i *dummyInteractor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(
		controller.NewUserController(
			service.NewUserService(
				json.NewUserRepository(),
				presenter.NewUserPresenter())))
}

func (i *dummyInteractor) NewCoordinateHandler() handler.CoordinateHandler {
	return handler.NewCoordinateHandler(
		controller.NewCoordinateController(
			service.NewCoordinateService(
				i.imagePath,
				json.NewCoordinateRepository(),
				presenter.NewCoordinatePresenter())))
}

func (i *dummyInteractor) NewFavoriteHandler() handler.FavoriteHandler {
	return handler.NewFavoriteHandler(
		controller.NewFavoriteController(
			service.NewFavoriteService(
				json.NewFavoriteRepository())))
}

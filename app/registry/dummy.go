package registry

import (
	"github.com/tayusa/notugly_backend/app/infrastructure/api/handler"
	"github.com/tayusa/notugly_backend/app/infrastructure/repository/dummy"
	"github.com/tayusa/notugly_backend/app/interface/controller"
	"github.com/tayusa/notugly_backend/app/interface/presenter"
	"github.com/tayusa/notugly_backend/app/usecase/service"
)

type dummyInteractor struct {
}

func NewDummyInteractor() Iteractor {
	return &dummyInteractor{}
}

func (i *dummyInteractor) NewAppHandler() handler.AppHandler {
	return handler.NewAppHandler(
		i.NewUserHandler(), i.NewCoordinateHandler(), i.NewFavoriteHandler())
}

func (i *dummyInteractor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(
		controller.NewUserController(
			service.NewUserService(
				dummy.NewDummyUserRepository(),
				presenter.NewUserPresenter())))
}

func (i *dummyInteractor) NewCoordinateHandler() handler.CoordinateHandler {
	return handler.NewCoordinateHandler(
		controller.NewCoordinateController(
			service.NewCoordinateService(
				"testdata/images",
				dummy.NewDummyCoordinateRepository(),
				presenter.NewCoordinatePresenter())))
}

func (i *dummyInteractor) NewFavoriteHandler() handler.FavoriteHandler {
	return handler.NewFavoriteHandler(
		controller.NewFavoriteController(
			service.NewFavoriteService(
				dummy.NewDummyFavoriteRepository())))
}

package registry

import (
	"database/sql"

	"github.com/tayusa/notugly_backend/app/infrastructure/api/handler"
	"github.com/tayusa/notugly_backend/app/infrastructure/repository"
	"github.com/tayusa/notugly_backend/app/interface/controller"
	"github.com/tayusa/notugly_backend/app/interface/presenter"
	output "github.com/tayusa/notugly_backend/app/usecase/presenter"
	input "github.com/tayusa/notugly_backend/app/usecase/repository"
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

// user
func (i *interactor) NewUserHandler() handler.UserHandler {
	return handler.NewUserHandler(i.NewUserController())
}

func (i *interactor) NewUserController() controller.UserController {
	return controller.NewUserController(i.NewUserService())
}

func (i *interactor) NewUserService() service.UserService {
	return service.NewUserService(i.NewUserRepository(), i.NewUserPresenter())
}

func (i *interactor) NewUserRepository() input.UserRepository {
	return repository.NewUserRepository(i.db)
}

func (i *interactor) NewUserPresenter() output.UserPresenter {
	return presenter.NewUserPresenter()
}

// coordinate
func (i *interactor) NewCoordinateHandler() handler.CoordinateHandler {
	return handler.NewCoordinateHandler(i.NewCoordinateController())
}

func (i *interactor) NewCoordinateController() controller.CoordinateController {
	return controller.NewCoordinateController(i.NewCoordinateService())
}

func (i *interactor) NewCoordinateService() service.CoordinateService {
	return service.NewCoordinateService(i.NewCoordinateRepository(), i.NewCoordinatePresenter())
}

func (i *interactor) NewCoordinateRepository() input.CoordinateRepository {
	return repository.NewCoordinateRepository(i.db)
}

func (i *interactor) NewCoordinatePresenter() output.CoordinatePresenter {
	return presenter.NewCoordinatePresenter()
}

// favorite
func (i *interactor) NewFavoriteHandler() handler.FavoriteHandler {
	return handler.NewFavoriteHandler(i.NewFavoriteController())
}

func (i *interactor) NewFavoriteController() controller.FavoriteController {
	return controller.NewFavoriteController(i.NewFavoriteService())
}

func (i *interactor) NewFavoriteService() service.FavoriteService {
	return service.NewFavoriteService(i.NewFavoriteRepository())
}

func (i *interactor) NewFavoriteRepository() input.FavoriteRepository {
	return repository.NewFavoriteRepository(i.db)
}

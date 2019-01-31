package handler

type appHandler struct {
	UserHandler
	CoordinateHandler
	FavoriteHandler
}

type AppHandler interface {
	UserHandler
	CoordinateHandler
	FavoriteHandler
}

func NewAppHandler(
	user UserHandler,
	coordinate CoordinateHandler,
	favorite FavoriteHandler) AppHandler {

	return &appHandler{user, coordinate, favorite}
}

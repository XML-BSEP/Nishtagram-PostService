package handler

import (
	"github.com/gin-gonic/gin"
	"post-service/usecase"
)

type FavoriteHandler interface {
	AddPostToFavorite(context *gin.Context)
}

type favoriteHandler struct {
	favoriteUseCase usecase.FavoriteUseCase
}

func (f favoriteHandler) AddPostToFavorite(context *gin.Context) {
	panic("implement me")
}

func NewFavoriteHandler(favoriteUseCase usecase.FavoriteUseCase) FavoriteHandler {
	return &favoriteHandler{favoriteUseCase: favoriteUseCase}
}

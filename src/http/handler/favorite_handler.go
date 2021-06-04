package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"post-service/dto"
	"post-service/usecase"
)

type FavoriteHandler interface {
	AddPostToFavorite(context *gin.Context)
	RemovePostFromFavorites(context *gin.Context)
}

type favoriteHandler struct {
	favoriteUseCase usecase.FavoriteUseCase
}

func (f favoriteHandler) RemovePostFromFavorites(context *gin.Context) {
	var favoriteDTO dto.FavoriteDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&favoriteDTO); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	err := f.favoriteUseCase.RemovePostFromFavorites(favoriteDTO, context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, "ok")
}

func (f favoriteHandler) AddPostToFavorite(context *gin.Context) {
	var favoriteDTO dto.FavoriteDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&favoriteDTO); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	err := f.favoriteUseCase.AddPostToFavorites(favoriteDTO, context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, "ok")
}

func NewFavoriteHandler(favoriteUseCase usecase.FavoriteUseCase) FavoriteHandler {
	return &favoriteHandler{favoriteUseCase: favoriteUseCase}
}

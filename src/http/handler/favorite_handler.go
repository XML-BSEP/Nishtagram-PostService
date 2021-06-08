package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"post-service/dto"
	"post-service/http/middleware"
	"post-service/usecase"
)

type FavoriteHandler interface {
	AddPostToFavorite(context *gin.Context)
	RemovePostFromFavorites(context *gin.Context)
	GetFavorites(ctx *gin.Context)
}

type favoriteHandler struct {
	favoriteUseCase usecase.FavoriteUseCase
}

func (f favoriteHandler) GetFavorites(ctx *gin.Context) {
	userId, _ := middleware.ExtractUserId(ctx.Request)
	favorite, err := f.favoriteUseCase.GetFavoritesForUser(userId, context.Background())

	if err != nil {
		ctx.JSON(500, "server error")
		ctx.Abort()
		return
	}

	ctx.JSON(200, favorite)
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
		context.Abort()
		return
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

	favoriteDTO.UserId, _ = middleware.ExtractUserId(context.Request)
	err := f.favoriteUseCase.AddPostToFavorites(favoriteDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}

func NewFavoriteHandler(favoriteUseCase usecase.FavoriteUseCase) FavoriteHandler {
	return &favoriteHandler{favoriteUseCase: favoriteUseCase}
}

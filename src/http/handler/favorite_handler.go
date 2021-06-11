package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
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
	logger *logger.Logger
}

func (f favoriteHandler) GetFavorites(ctx *gin.Context) {
	f.logger.Logger.Println("Handling GETTING FAVORITES")
	userId, _ := middleware.ExtractUserId(ctx.Request, f.logger)
	favorite, err := f.favoriteUseCase.GetFavoritesForUser(userId, context.Background())

	if err != nil {
		ctx.JSON(500, "server error")
		ctx.Abort()
		return
	}

	ctx.JSON(200, favorite)
}

func (f favoriteHandler) RemovePostFromFavorites(context *gin.Context) {
	f.logger.Logger.Println("Handling REMOVING POST FROM FAVORITES")
	var favoriteDTO dto.FavoriteDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&favoriteDTO); err != nil {
		f.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}
	favoriteDTO.UserId, _ = middleware.ExtractUserId(context.Request, f.logger)
	err := f.favoriteUseCase.RemovePostFromFavorites(favoriteDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}

func (f favoriteHandler) AddPostToFavorite(context *gin.Context) {
	f.logger.Logger.Println("Handling ADDING POST TO FAVORITES")
	var favoriteDTO dto.FavoriteDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&favoriteDTO); err != nil {
		f.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	favoriteDTO.UserId, _ = middleware.ExtractUserId(context.Request, f.logger)
	err := f.favoriteUseCase.AddPostToFavorites(favoriteDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}

func NewFavoriteHandler(favoriteUseCase usecase.FavoriteUseCase, logger *logger.Logger) FavoriteHandler {
	return &favoriteHandler{favoriteUseCase: favoriteUseCase, logger: logger}
}

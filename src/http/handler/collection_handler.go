package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"post-service/dto"
	"post-service/http/middleware"
	"post-service/usecase"
)

type CollectionHandler interface {
	CreateCollection(context *gin.Context)
	AddPostToCollection(context *gin.Context)
	RemovePostFromCollection(context *gin.Context)
	DeleteCollection(context *gin.Context)
	GetAllCollections(ctx *gin.Context)
	GetCollection(ctx *gin.Context)
}

type collectionHandler struct {
	collectionUseCase usecase.CollectionUseCase
}

func (c collectionHandler) GetCollection(ctx *gin.Context) {
	var collectionDTO dto.CollectionDTO

	decoder := json.NewDecoder(ctx.Request.Body)

	if err := decoder.Decode(&collectionDTO); err != nil {
		ctx.JSON(400, "invalid request")
		ctx.Abort()
		return
	}
	userId, _ := middleware.ExtractUserId(ctx.Request)
	collection, err := c.collectionUseCase.GetCollection(userId, collectionDTO.CollectionName, context.Background())
	if err != nil {
		ctx.JSON(500, "server error")
		ctx.Abort()
		return
	}

	ctx.JSON(200, collection)

}

func (c collectionHandler) GetAllCollections(ctx *gin.Context) {

	userId, _ := middleware.ExtractUserId(ctx.Request)

	collections, err := c.collectionUseCase.GetAllCollectionsPerUser(userId, ctx)

	if err != nil {
		ctx.JSON(500, "server error")
		ctx.Abort()
		return
	}

	ctx.JSON(200, collections)
}

func (c collectionHandler) CreateCollection(context *gin.Context) {
	var collectionDTO dto.CollectionDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&collectionDTO); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	err := c.collectionUseCase.CreateCollection(collectionDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")

}

func (c collectionHandler) RemovePostFromCollection(context *gin.Context) {
	var collectionDTO dto.CollectionDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&collectionDTO); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	err := c.collectionUseCase.RemovePostFromCollection(collectionDTO, context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, "ok")
}

func (c collectionHandler) DeleteCollection(context *gin.Context) {
	var collectionDTO dto.CollectionDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&collectionDTO); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	err := c.collectionUseCase.DeleteCollection(collectionDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}

func (c collectionHandler) AddPostToCollection(context *gin.Context) {
	var collectionDTO dto.CollectionDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&collectionDTO); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}
	collectionDTO.UserId, _ = middleware.ExtractUserId(context.Request)
	err := c.collectionUseCase.AddPostToCollection(collectionDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}

func NewCollectionHandler(collectionUseCase usecase.CollectionUseCase) CollectionHandler {
	return &collectionHandler{collectionUseCase: collectionUseCase}
}
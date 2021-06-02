package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"post-service/dto"
	"post-service/usecase"
)

type CollectionHandler interface {
	CreateCollection(context *gin.Context)
	AddPostToCollection(context *gin.Context)
	RemovePostFromCollection(context *gin.Context)
	DeleteCollection(context *gin.Context)
}

type collectionHandler struct {
	collectionUseCase usecase.CollectionUseCase
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

	err := c.collectionUseCase.AddPostToCollection(collectionDTO, context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, "ok")
}

func NewCollectionHandler(collectionUseCase usecase.CollectionUseCase) CollectionHandler {
	return &collectionHandler{collectionUseCase: collectionUseCase}
}
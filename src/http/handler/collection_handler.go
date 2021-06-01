package handler

import (
	"github.com/gin-gonic/gin"
	"post-service/usecase"
)

type CollectionHandler interface {
	AddPostToCollection(context *gin.Context)
}

type collectionHandler struct {
	collectionUseCase usecase.CollectionUseCase
}

func (c collectionHandler) AddPostToCollection(context *gin.Context) {
	panic("implement me")
}

func NewCollectionHandler(collectionUseCase usecase.CollectionUseCase) CollectionHandler {
	return &collectionHandler{collectionUseCase: collectionUseCase}
}
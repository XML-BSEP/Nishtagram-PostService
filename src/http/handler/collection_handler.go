package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
	"github.com/microcosm-cc/bluemonday"
	"post-service/dto"
	"post-service/http/middleware"
	"post-service/usecase"
	"strings"
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
	logger *logger.Logger
}

func (c collectionHandler) GetCollection(ctx *gin.Context) {
	c.logger.Logger.Println("Handling GETTING COLLECTION")
	var collectionDTO dto.CollectionDTO

	decoder := json.NewDecoder(ctx.Request.Body)

	if err := decoder.Decode(&collectionDTO); err != nil {
		c.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		ctx.JSON(400, "invalid request")
		ctx.Abort()
		return
	}
	userId, _ := middleware.ExtractUserId(ctx.Request, c.logger)
	collection, err := c.collectionUseCase.GetCollection(userId, collectionDTO.CollectionName, context.Background())
	if err != nil {
		ctx.JSON(500, "server error")
		ctx.Abort()
		return
	}

	ctx.JSON(200, collection)

}

func (c collectionHandler) GetAllCollections(ctx *gin.Context) {
	c.logger.Logger.Println("Handling GETTING ALL COLLECTIONS")
	userId, _ := middleware.ExtractUserId(ctx.Request, c.logger)
	collections, err := c.collectionUseCase.GetAllCollectionsPerUser(userId, ctx)
	if err != nil {
		ctx.JSON(500, "server error")
		ctx.Abort()
		return
	}
	ctx.JSON(200, collections)
}

func (c collectionHandler) CreateCollection(context *gin.Context) {
	c.logger.Logger.Println("HANDLING CREATING COLLECTION")
	var collectionDTO dto.CollectionDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&collectionDTO); err != nil {
		c.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
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
	c.logger.Logger.Println("Handling REMOVING POST FROM COLLECTION")
	var collectionDTO dto.CollectionDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&collectionDTO); err != nil {
		c.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
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
	c.logger.Logger.Println("Handling DELETING COLLECTION")
	var collectionDTO dto.CollectionDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&collectionDTO); err != nil {
		c.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
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
	c.logger.Logger.Println("Handling ADDING POST TO COLLECTION")
	var collectionDTO dto.CollectionDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&collectionDTO); err != nil {
		c.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	policy := bluemonday.UGCPolicy()
	collectionDTO.PostBy = strings.TrimSpace(policy.Sanitize(collectionDTO.PostBy))
	collectionDTO.CollectionName = strings.TrimSpace(policy.Sanitize(collectionDTO.CollectionName))
	collectionDTO.PostId = strings.TrimSpace(policy.Sanitize(collectionDTO.PostId))

	if collectionDTO.PostBy == "" || collectionDTO.CollectionName == "" || collectionDTO.PostId ==  ""{
		c.logger.Logger.Errorf("error while verifying and validating collection fields\n")
		c.logger.Logger.Warnf("possible xss attack from IP address: %v\n", context.Request.Referer())
		context.JSON(400, gin.H{"message" : "Fields are empty or xss attack happened"})
		return
	}


	collectionDTO.UserId, _ = middleware.ExtractUserId(context.Request, c.logger)
	err := c.collectionUseCase.AddPostToCollection(collectionDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}

func NewCollectionHandler(collectionUseCase usecase.CollectionUseCase, logger *logger.Logger) CollectionHandler {
	return &collectionHandler{collectionUseCase: collectionUseCase, logger: logger}
}
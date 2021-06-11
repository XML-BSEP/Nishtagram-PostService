package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
	"post-service/dto"
	"post-service/http/middleware"
	"post-service/usecase"
)

type LikeHandler interface {
	LikePost(context *gin.Context)
	DislikePost(context *gin.Context)
	RemoveLike(context *gin.Context)
	RemoveDislike(ctx *gin.Context)
}

type likeHandler struct {
	likeUseCase usecase.LikeUseCase
	logger *logger.Logger
}

func (l likeHandler) DislikePost(context *gin.Context) {
	l.logger.Logger.Println("Handling DISLIKING POST")
	var dislikeDTO dto.LikeDislikeDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&dislikeDTO); err != nil {
		l.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}
	dislikeDTO.UserId, _ = middleware.ExtractUserId(context.Request, l.logger)

	err := l.likeUseCase.DislikePost(dislikeDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}

func (l likeHandler) RemoveLike(context *gin.Context) {
	l.logger.Logger.Println("Handling LIKING POSTS")
	var likeDTO dto.LikeDislikeDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&likeDTO); err != nil {
		l.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	likeDTO.UserId, _ = middleware.ExtractUserId(context.Request, l.logger)

	err := l.likeUseCase.RemoveLike(likeDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}

func (l likeHandler) RemoveDislike(context *gin.Context) {
	l.logger.Logger.Println("Handling REMOVING DISLIKE")
	var dislikeDTO dto.LikeDislikeDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&dislikeDTO); err != nil {
		l.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	dislikeDTO.UserId, _ = middleware.ExtractUserId(context.Request, l.logger)
	err := l.likeUseCase.RemoveDislike(dislikeDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}

func (l likeHandler) LikePost(context *gin.Context) {
	l.logger.Logger.Println("Handling LIKING POST")
	var likeDTO dto.LikeDislikeDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&likeDTO); err != nil {
		l.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	likeDTO.UserId, _ = middleware.ExtractUserId(context.Request, l.logger)

	err := l.likeUseCase.LikePost(likeDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}

func NewLikeHandler(likeUseCase usecase.LikeUseCase, logger *logger.Logger) LikeHandler {
	return &likeHandler{likeUseCase: likeUseCase, logger: logger}
}

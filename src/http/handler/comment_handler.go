package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"post-service/dto"
	"post-service/usecase"
)

type CommentHandler interface {
	AddComment(context *gin.Context)
	DeleteComment(context *gin.Context)
}

type commentHandler struct {
	commentUseCase usecase.CommentUseCase
}

func (c commentHandler) DeleteComment(context *gin.Context) {
	var commentDTO dto.CommentDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&commentDTO); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	err := c.commentUseCase.DeleteComment(commentDTO, context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, "ok")
}

func (c commentHandler) AddComment(context *gin.Context) {
	var commentDTO dto.CommentDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&commentDTO); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	err := c.commentUseCase.AddComment(commentDTO, context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, "ok")
}

func NewCommentHandler(commentUseCase usecase.CommentUseCase) CommentHandler {
	return &commentHandler{commentUseCase: commentUseCase}
}
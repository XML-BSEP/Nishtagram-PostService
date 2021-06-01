package handler

import (
	"github.com/gin-gonic/gin"
	"post-service/usecase"
)

type CommentHandler interface {
	AddComment(context *gin.Context)
}

type commentHandler struct {
	commentUseCase usecase.CommentUseCase
}

func (c commentHandler) AddComment(context *gin.Context) {
	panic("implement me")
}

func NewCommentHandler(commentUseCase usecase.CommentUseCase) CommentHandler {
	return &commentHandler{commentUseCase: commentUseCase}
}
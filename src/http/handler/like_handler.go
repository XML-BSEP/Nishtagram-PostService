package handler

import (
	"github.com/gin-gonic/gin"
	"post-service/usecase"
)

type LikeHandler interface {
	LikePost(context *gin.Context)
}

type likeHandler struct {
	likeUseCase usecase.LikeUseCase
}

func (l likeHandler) LikePost(context *gin.Context) {
	panic("implement me")
}

func NewLikeHandler(likeUseCase usecase.LikeUseCase) LikeHandler {
	return &likeHandler{likeUseCase: likeUseCase}
}

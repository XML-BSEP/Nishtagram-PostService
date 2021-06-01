package handler

import (
	"github.com/gin-gonic/gin"
	"post-service/usecase"
)

type PostHandler interface {
	AddPost(context *gin.Context)
}

type postHandler struct {
	postUseCase usecase.PostUseCase
}

func (p postHandler) AddPost(context *gin.Context) {
	panic("implement me")
}


func NewPostHandler(postUseCase usecase.PostUseCase) PostHandler {
	return &postHandler{postUseCase: postUseCase}
}

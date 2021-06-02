package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"post-service/dto"
	"post-service/usecase"
)

type PostHandler interface {
	AddPost(context *gin.Context)
	EditPost(context *gin.Context)
	DeletePost(context *gin.Context)
	GetPostsByUser(context *gin.Context)
}

type postHandler struct {
	postUseCase usecase.PostUseCase
}

func (p postHandler) EditPost(context *gin.Context) {
	var updatePostDTO dto.UpdatePostDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&updatePostDTO); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	err := p.postUseCase.EditPost(updatePostDTO, context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, "ok")
}

func (p postHandler) DeletePost(context *gin.Context) {
	var deletePostDTO dto.DeletePostDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&deletePostDTO); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	err := p.postUseCase.DeletePost(deletePostDTO, context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, "ok")
}

func (p postHandler) GetPostsByUser(context *gin.Context) {
	var userId string

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&userId); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	posts, err := p.postUseCase.GetPostsByUser(userId, context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, posts)
}

func (p postHandler) AddPost(context *gin.Context) {
	var createPostDTO dto.CreatePostDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&createPostDTO); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	err := p.postUseCase.AddPost(createPostDTO, context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, "ok")
}


func NewPostHandler(postUseCase usecase.PostUseCase) PostHandler {
	return &postHandler{postUseCase: postUseCase}
}
package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"post-service/domain"
	"post-service/dto"
	"post-service/http/middleware"
	"post-service/usecase"
)

type PostHandler interface {
	AddPost(context *gin.Context)
	EditPost(context *gin.Context)
	DeletePost(context *gin.Context)
	GetPostsByUser(context *gin.Context)
	GenerateUserFeed(context *gin.Context)
	GetPostsOnProfile(context *gin.Context)
	GetPostById(ctx *gin.Context)
}

type postHandler struct {
	postUseCase usecase.PostUseCase
}

func (p postHandler) GetPostById(ctx *gin.Context) {
	var postDTO dto.GetPostDTO
	decoder := json.NewDecoder(ctx.Request.Body)

	if err := decoder.Decode(&postDTO); err != nil {
		ctx.JSON(400, "invalid request")
		ctx.Abort()
		return
	}
	userRequested, _ := middleware.ExtractUserId(ctx.Request)

	post, err := p.postUseCase.GetPost(postDTO.PostId, postDTO.UserId, userRequested, context.Background())
	if err != nil {
		ctx.JSON(500, gin.H{"message":"server error"})
		ctx.Abort()
		return
	}

	ctx.JSON(200, post)


}

func (p postHandler) GetPostsOnProfile(ctx *gin.Context) {
	var userDTO domain.Profile
	decoder := json.NewDecoder(ctx.Request.Body)

	if err := decoder.Decode(&userDTO); err != nil {
		ctx.JSON(400, "invalid request")
		ctx.Abort()
		return
	}

	userRequested, _ := middleware.ExtractUserId(ctx.Request)

	posts, err := p.postUseCase.GetPostsOnProfile(userDTO.Id, userRequested, context.Background())

	if err != nil {
		ctx.JSON(500, gin.H{"message":"server error"})
		ctx.Abort()
		return
	}

	ctx.JSON(200, posts)

}

func (p postHandler) GenerateUserFeed(context *gin.Context) {
	userId, err := middleware.ExtractUserId(context.Request)
	if err != nil {
		context.JSON(500, gin.H{"message":"server error"})
		context.Abort()
		return
	}

	posts, err := p.postUseCase.GenerateUserFeed(userId, userId, context)
	if err != nil {
		context.JSON(500, gin.H{"message":"server error"})
		context.Abort()
		return
	}

	context.JSON(200, posts)

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
		context.Abort()
		return
	}

	context.JSON(200, gin.H{"message" : "ok"})
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
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}

func (p postHandler) GetPostsByUser(context *gin.Context) {
	userRequested, err := middleware.ExtractUserId(context.Request)
	var userId dto.UserTag

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&userId); err != nil {
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	posts, err := p.postUseCase.GetPostsByUser(userId.UserId, userRequested, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
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
	createPostDTO.UserId.UserId, _ = middleware.ExtractUserId(context.Request)
	err := p.postUseCase.AddPost(createPostDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}


func NewPostHandler(postUseCase usecase.PostUseCase) PostHandler {
	return &postHandler{postUseCase: postUseCase}
}

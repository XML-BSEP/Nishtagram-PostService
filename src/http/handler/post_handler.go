package handler

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
	"github.com/microcosm-cc/bluemonday"
	"post-service/domain"
	"post-service/dto"
	"post-service/http/middleware"
	"post-service/usecase"
	"strings"
)

type PostHandler interface {
	AddPost(context *gin.Context)
	EditPost(context *gin.Context)
	DeletePost(context *gin.Context)
	GetPostsByUser(context *gin.Context)
	GenerateUserFeed(context *gin.Context)
	GetPostsOnProfile(context *gin.Context)
	GetPostById(ctx *gin.Context)
	GetLikedMedia(ctx *gin.Context)
	GetDislikedMedia(ctx *gin.Context)
}

type postHandler struct {
	postUseCase usecase.PostUseCase
	logger *logger.Logger
}

func (p postHandler) GetLikedMedia(ctx *gin.Context) {
	userId, _ := middleware.ExtractUserId(ctx.Request, p.logger)

	posts, err := p.postUseCase.GetAllLikedMedia(userId, ctx)

	if err != nil {
		ctx.JSON(500, gin.H{"message":"server error"})
		ctx.Abort()
		return
	}

	ctx.JSON(200, posts)
}

func (p postHandler) GetDislikedMedia(ctx *gin.Context) {
	userId, _ := middleware.ExtractUserId(ctx.Request, p.logger)

	posts, err := p.postUseCase.GetAllDislikedMedia(userId, ctx)

	if err != nil {
		ctx.JSON(500, gin.H{"message":"server error"})
		ctx.Abort()
		return
	}

	ctx.JSON(200, posts)
}

func (p postHandler) GetPostById(ctx *gin.Context) {
	p.logger.Logger.Println("Handling GETTING POST BY ID")
	var postDTO dto.GetPostDTO
	decoder := json.NewDecoder(ctx.Request.Body)

	if err := decoder.Decode(&postDTO); err != nil {
		p.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		ctx.JSON(400, "invalid request")
		ctx.Abort()
		return
	}
	userRequested, _ := middleware.ExtractUserId(ctx.Request, p.logger)

	post, err := p.postUseCase.GetPost(postDTO.PostId, postDTO.UserId, userRequested, context.Background())
	if err != nil {
		ctx.JSON(500, gin.H{"message":"server error"})
		ctx.Abort()
		return
	}

	ctx.JSON(200, post)


}

func (p postHandler) GetPostsOnProfile(ctx *gin.Context) {
	p.logger.Logger.Println("Handling GETTING POSTS ON PROFILE")
	var userDTO domain.Profile
	decoder := json.NewDecoder(ctx.Request.Body)

	if err := decoder.Decode(&userDTO); err != nil {
		p.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		ctx.JSON(400, "invalid request")
		ctx.Abort()
		return
	}

	userRequested, _ := middleware.ExtractUserId(ctx.Request, p.logger)

	posts, err := p.postUseCase.GetPostsOnProfile(userDTO.Id, userRequested, context.Background())

	if err != nil {
		ctx.JSON(500, gin.H{"message":"server error"})
		ctx.Abort()
		return
	}

	ctx.JSON(200, posts)

}

func (p postHandler) GenerateUserFeed(context *gin.Context) {
	p.logger.Logger.Println("Handling GENERATING USER FEED")
	userId, err := middleware.ExtractUserId(context.Request, p.logger)
	if err != nil {
		p.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
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
	p.logger.Logger.Println("Handling EDITING POST")
	var updatePostDTO dto.UpdatePostDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&updatePostDTO); err != nil {
		p.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
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
	p.logger.Logger.Println("Handling DELETING POST")
	var deletePostDTO dto.DeletePostDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&deletePostDTO); err != nil {
		p.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
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
	p.logger.Logger.Println("Handling GETTING POSTS BY USER")
	userRequested, err := middleware.ExtractUserId(context.Request, p.logger)
	var userId dto.UserTag

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&userId); err != nil {
		p.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
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
	p.logger.Logger.Println("Handling ADDING POST")
	var createPostDTO dto.CreatePostDTO

	decoder := json.NewDecoder(context.Request.Body)
	if err := decoder.Decode(&createPostDTO); err != nil {
		p.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	policy := bluemonday.UGCPolicy()

	createPostDTO.Location =  strings.TrimSpace(policy.Sanitize(createPostDTO.Location))
	createPostDTO.Caption =  strings.TrimSpace(policy.Sanitize(createPostDTO.Caption))
	for i,_ := range createPostDTO.Hashtags{
		createPostDTO.Hashtags[i] =  strings.TrimSpace(policy.Sanitize(createPostDTO.Hashtags[i]))
		if createPostDTO.Hashtags[i] == "" {
			p.logger.Logger.Errorf("fields are empty or xss attack happened")
			context.JSON(400, gin.H{"message" : "Fields are empty or xss attack happened"})
			return
		}
	}
	for i,_ := range createPostDTO.Album{
		createPostDTO.Album[i] =  strings.TrimSpace(policy.Sanitize(createPostDTO.Album[i]))
		if createPostDTO.Album[i] == "" {
			p.logger.Logger.Errorf("fields are empty or xss attack happened")
			context.JSON(400, gin.H{"message" : "Fields are empty or xss attack happened"})
			return
		}
	}

	createPostDTO.Image =  strings.TrimSpace(policy.Sanitize(createPostDTO.Image))
	createPostDTO.Video =  strings.TrimSpace(policy.Sanitize(createPostDTO.Video))

	if createPostDTO.Location == "" || createPostDTO.Caption == ""  {
		p.logger.Logger.Errorf("error while verifying and validating createPostDTO fields\n")
		p.logger.Logger.Warnf("possible xss attack from IP address: %v\n", context.Request.Referer())
		context.JSON(400, gin.H{"message" : "Fields are empty or xss attack happened"})
		return
	}



	createPostDTO.UserId.UserId, _ = middleware.ExtractUserId(context.Request, p.logger)
	err := p.postUseCase.AddPost(createPostDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}


func NewPostHandler(postUseCase usecase.PostUseCase, logger *logger.Logger) PostHandler {
	return &postHandler{postUseCase: postUseCase, logger: logger}
}

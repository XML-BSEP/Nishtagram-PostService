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

type CommentHandler interface {
	AddComment(context *gin.Context)
	DeleteComment(context *gin.Context)
	GetComments(ctx *gin.Context)
}

type commentHandler struct {
	commentUseCase usecase.CommentUseCase
	logger *logger.Logger
}

func (c commentHandler) GetComments(ctx *gin.Context) {
	c.logger.Logger.Println("Handling GETTING COMMENTS")
	var postDTO dto.PostDTO

	decoder := json.NewDecoder(ctx.Request.Body)

	if err := decoder.Decode(&postDTO); err != nil {
		c.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		ctx.JSON(400, "invalid request")
		ctx.Abort()
		return
	}

	comments, err := c.commentUseCase.GetAllCommentsByPost(postDTO.Id, context.Background())

	if err != nil {
		ctx.JSON(500, "server error")
		ctx.Abort()
		return
	}

	ctx.JSON(200, comments)
}

func (c commentHandler) DeleteComment(context *gin.Context) {
	c.logger.Logger.Println("Handling DELETING COMMENTS")
	var commentDTO dto.CommentDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&commentDTO); err != nil {
		c.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	err := c.commentUseCase.DeleteComment(commentDTO, context)

	if err != nil {
		context.JSON(500, "server error")
		context.Abort()
		return
	}

	context.JSON(200, "ok")
}

func (c commentHandler) AddComment(context *gin.Context) {
	c.logger.Logger.Println("Handling ADDING COMMENTS")
	var commentDTO dto.CommentDTO

	decoder := json.NewDecoder(context.Request.Body)

	if err := decoder.Decode(&commentDTO); err != nil {
		c.logger.Logger.Errorf("error while decoding json, error: %v\n", err)
		context.JSON(400, "invalid request")
		context.Abort()
		return
	}

	policy := bluemonday.UGCPolicy()
	commentDTO.Comment =  strings.TrimSpace(policy.Sanitize(commentDTO.Comment))
	commentDTO.PostId =  strings.TrimSpace(policy.Sanitize(commentDTO.PostId))
	commentDTO.PostBy =  strings.TrimSpace(policy.Sanitize(commentDTO.PostBy))

	if commentDTO.Comment == "" || commentDTO.PostId == "" || commentDTO.PostBy ==  ""{
		c.logger.Logger.Errorf("error while verifying and validating comment fields\n")
		c.logger.Logger.Warnf("possible xss attack from IP address: %v\n", context.Request.Referer())
		context.JSON(400, gin.H{"message" : "Fields are empty or xss attack happened"})
		return
	}


	commentDTO.CommentBy.Id, _ = middleware.ExtractUserId(context.Request, c.logger)
	err := c.commentUseCase.AddComment(commentDTO, context)

	if err != nil {
		context.JSON(500, "server error")
	}

	context.JSON(200, "ok")
}

func NewCommentHandler(commentUseCase usecase.CommentUseCase, logger *logger.Logger) CommentHandler {
	return &commentHandler{commentUseCase: commentUseCase, logger: logger}
}
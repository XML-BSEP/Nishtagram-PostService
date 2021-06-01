package usecase

import (
	"context"
	"post-service/dto"
	"post-service/repository"
)

type CommentUseCase interface {
	AddComment(comment dto.CommentDTO, ctx context.Context) error
	DeleteComment(comment dto.CommentDTO, ctx context.Context) error
}

type commentUseCase struct {
	commentRepository repository.CommentRepo
}

func (c commentUseCase) AddComment(comment dto.CommentDTO, ctx context.Context) error {
	panic("implement me")
}

func (c commentUseCase) DeleteComment(comment dto.CommentDTO, ctx context.Context) error {
	panic("implement me")
}

func NewCommentUseCase(commentRepository repository.CommentRepo) CommentUseCase {
	return &commentUseCase{
		commentRepository: commentRepository,
	}
}
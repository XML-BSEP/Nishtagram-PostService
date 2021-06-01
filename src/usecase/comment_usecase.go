package usecase

import (
	"post-service/dto"
	"post-service/repository"
)

type CommentUseCase interface {
	AddComment(comment dto.CommentDTO) error
	DeleteComment(comment dto.CommentDTO) error
}

type commentUseCase struct {
	commentRepository repository.CommentRepo
}

func (c commentUseCase) AddComment(comment dto.CommentDTO) error {
	panic("implement me")
}

func (c commentUseCase) DeleteComment(comment dto.CommentDTO) error {
	panic("implement me")
}

func NewCommentUseCase(commentRepository repository.CommentRepo) CommentUseCase {
	return &commentUseCase{
		commentRepository: commentRepository,
	}
}
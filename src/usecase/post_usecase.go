package usecase

import (
	"context"
	"post-service/dto"
	"post-service/repository"
)

type PostUseCase interface {
	AddPost(postDTO dto.PostDTO, ctx context.Context)
	DeletePost(postDTO dto.PostDTO, ctx context.Context)
	EditPost(postDTO dto.PostDTO, ctx context.Context)
	GetPostsForUserFeed(userId uint, ctx context.Context)
	GetPost(postId uint, userId uint, ctx context.Context)
}

type postUseCase struct {
	postRepository repository.PostRepo
}

func (p postUseCase) AddPost(postDTO dto.PostDTO, ctx context.Context) {
	panic("implement me")
}

func (p postUseCase) DeletePost(postDTO dto.PostDTO, ctx context.Context) {
	panic("implement me")
}

func (p postUseCase) EditPost(postDTO dto.PostDTO, ctx context.Context) {
	panic("implement me")
}

func (p postUseCase) GetPostsForUserFeed(userId uint, ctx context.Context) {
	panic("implement me")
}

func (p postUseCase) GetPost(postId uint, userId uint, ctx context.Context) {
	panic("implement me")
}

func NewPostUseCase(postRepository repository.PostRepo) PostUseCase {
	return &postUseCase{postRepository: postRepository}
}

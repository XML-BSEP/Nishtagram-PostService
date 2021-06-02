package usecase

import (
	"context"
	"post-service/dto"
	"post-service/repository"
)

type PostUseCase interface {
	AddPost(postDTO dto.CreatePostDTO, ctx context.Context) error
	DeletePost(postDTO dto.DeletePostDTO, ctx context.Context) error
	EditPost(postDTO dto.UpdatePostDTO, ctx context.Context) error
	GetPostsByUser(userId string, ctx context.Context) ([]dto.PostDTO, error)
	GetPost(postId string, userId string, ctx context.Context) (dto.PostDTO, error)
}

type postUseCase struct {
	postRepository repository.PostRepo
}

func (p postUseCase) AddPost(postDTO dto.CreatePostDTO, ctx context.Context) error {
	return p.postRepository.CreatePost(postDTO, context.Background())
}

func (p postUseCase) DeletePost(postDTO dto.DeletePostDTO, ctx context.Context) error {
	return p.postRepository.DeletePost(postDTO, context.Background())
}

func (p postUseCase) EditPost(postDTO dto.UpdatePostDTO, ctx context.Context) error {
	return p.EditPost(postDTO, context.Background())
}

func (p postUseCase) GetPostsByUser(userId string, ctx context.Context) ([]dto.PostDTO, error) {
	return p.postRepository.GetPostsByUserId(userId, context.Background())
}

func (p postUseCase) GetPost(postId string, userId string, ctx context.Context) (dto.PostDTO, error) {
	return p.postRepository.GetPostsById(userId, postId)
}

func NewPostUseCase(postRepository repository.PostRepo) PostUseCase {
	return &postUseCase{postRepository: postRepository}
}
